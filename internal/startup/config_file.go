package startup

import (
	"net"
	"net/url"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/joshw34/navitui/internal/types"
	"github.com/zalando/go-keyring"
)

type config struct {
	User          string `toml:"user"`
	Server        string `toml:"server"`
	Port          int    `toml:"port"`
	PasswordStore string `toml:"password_store"`
	Password      string `toml:"-"`
	FullURL       string `toml:"-"`
}

type configValidationStatus int

const (
	InvalidUsername configValidationStatus = iota
	InvalidServer
	InvalidPort
	InvalidPasswordStore
	ValidationSuccess
)

type promptMessage string

const (
	ServerPrompt     promptMessage = "Server URL (without port number):"
	PortPrompt       promptMessage = "Server port:"
	PasswordPrompt   promptMessage = "Password:"
	UsernamePrompt   promptMessage = "Username:"
	ErrorContinue    promptMessage = "Press 'Enter' to continue or 'ctrl+c' to quit"
	ErrorSetupWizard promptMessage = "Press 'Enter' to run setup wizard (this will overwrite your current config file)\nPress 'ctrl+c' to quit"
	ErrorTryAgain    promptMessage = "Press 'Enter' to try again or 'ctrl+c' to quit"
)

func (t *config) parseConfig(filePath string, logger *types.NavituiLogger) error {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		logger.File("Failed to open config file: %s", err.Error())
		return err
	}
	defer f.Close()
	_, err = toml.NewDecoder(f).Decode(t)
	if err != nil {
		logger.File("Failed to parse config file: %s", err.Error())
		return err
	}
	logger.File("Successfully parsed config file")
	return nil
}

func (t *config) saveConfig(filePath string, logger *types.NavituiLogger) {
	f, err := os.Create(filePath)
	if err != nil {
		logger.File("Failed to open/create config file: %s", err.Error())
		_, _ = promptError("Unable to open/create config file, configuration will not be saved", ErrorContinue, logger)
		return
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	encoder := toml.NewEncoder(f)
	err = encoder.Encode(t)
	if err != nil {
		logger.File("Failed to save config file: %s", err.Error())
		_, _ = promptError("Unable to save config file, configuration will not be saved", ErrorContinue, logger)
		return
	}
	logger.File("Successfully saved config file")
}

func (t *config) validateConfig(logger *types.NavituiLogger) configValidationStatus {
	status := ValidationSuccess

	switch {
	case t.User == "":
		status = InvalidUsername
	case t.Server == "":
		status = InvalidServer
	case t.Port < 1 || t.Port > 65535:
		status = InvalidPort
	case t.PasswordStore != "keyring" && t.PasswordStore != "prompt":
		status = InvalidPasswordStore
	}

	if status != ValidationSuccess {
		logger.File("Failed to validate config file: %s", getValidationFailureMsg(status))
		return status
	}

	u, err := url.Parse(t.Server)
	if err != nil {
		logger.File("Failed to parse server URL: %s", getValidationFailureMsg(InvalidServer))
		return InvalidServer
	}
	u.Host = net.JoinHostPort(u.Hostname(), strconv.Itoa(t.Port))
	t.FullURL = u.String()

	logger.File("Successfully validated config file")
	return ValidationSuccess
}

func (t *config) setDefaultOptions() {
	t.PasswordStore = "keyring"
}

func (t *config) getPassword(logger *types.NavituiLogger) error {
	switch t.PasswordStore {
	case "keyring":
		p, err := keyring.Get("navitui", t.User)
		if err != nil {
			logger.File("Failed to get password from keyring: %s", err.Error())
			return err
		}
		t.Password = p
		logger.File("Successfully got password from keyring")
		return nil
	default:
		ClearScreen()
		p, err := promptText("", PasswordPrompt)
		if err != nil {
			logger.File("Failed to get password from prompt: %s", err.Error())
			return err
		}
		t.Password = p
		logger.File("Successfully got password from user prompt")
		return nil
	}
}

func (t *config) savePassword(logger *types.NavituiLogger) {
	err := keyring.Set("navitui", t.User, t.Password)
	if err != nil {
		t.PasswordStore = "prompt"
		logger.File("Failed to save password to keyring: %s", err.Error())
		_, _ = promptError("Unable to save password to keyring, password will be required on startup", ErrorContinue, logger) // Ignore prompt error -> continue
	} else {
		t.PasswordStore = "keyring"
		logger.File("Successfully saved password to keyring")
	}
}

func (t *config) getUserInput(logger *types.NavituiLogger) error {
	var err error

	t.User, err = promptText(welcomeMessage, UsernamePrompt)
	if err != nil {
		logger.File("Failed to get user input: %s", err.Error())
		return err
	}

	t.Password, err = promptText(welcomeMessage, PasswordPrompt)
	if err != nil {
		logger.File("Failed to get user input: %s", err.Error())
		return err
	}

	t.Server, err = promptText(welcomeMessage, ServerPrompt)
	if err != nil {
		logger.File("Failed to get user input: %s", err.Error())
		return err
	}

	t.Port, err = promptInt(welcomeMessage, PortPrompt)
	if err != nil {
		logger.File("Failed to get user input: %s", err.Error())
		return err
	}

	ClearScreen()

	logger.File("Successfully got user input")
	return nil
}

func (t *config) tryPing(getServer srvbuild, logger *types.NavituiLogger) (types.Server, types.PingResult) {
	srv, err := getServer("navidrome", Credentials{
		BaseUrl:  t.FullURL,
		User:     t.User,
		Password: t.Password,
	})
	if err != nil {
		logger.File("Failed to connect to server: %s", err.Error())
		return nil, types.PingServerError
	}
	res := srv.PingTest()
	switch res {
	case types.PingSuccess:
		logger.File("Successfully connected to server")
		return srv, res
	default:
		logger.File("Failed to connect to server: %s", getPingFailureMsg(res))
		return nil, res
	}
}
