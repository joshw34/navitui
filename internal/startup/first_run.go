package startup

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/joshw34/navitui/internal/types"
)

const welcomeMessage = "Welcome to Navitui!"

func firstRun(getServer srvbuild, configFile string, logger *types.NavituiLogger) (types.Server, ConfigOptions, error) {
	for {
		var conf config
		conf.setDefaultOptions()
		err := conf.getUserInput(logger)
		if err != nil {
			return nil, ConfigOptions{}, err
		}

		v := conf.validateConfig(logger)
		if v != ValidationSuccess {
			_, _ = promptError(getValidationFailureMsg(v), ErrorContinue, logger) // Ignore error -> try again
			continue
		}

		srv, res := conf.tryPing(getServer, logger)
		if res == types.PingSuccess {
			conf.savePassword(logger)
			conf.saveConfig(configFile, logger)
			return srv, ConfigOptions{}, nil
		}

		_, _ = promptError(getPingFailureMsg(res), ErrorContinue, logger) // Ignore error -> retry
		continue
	}
}

func promptText(header string, msg promptMessage) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		ClearScreen()
		if len(header) != 0 {
			fmt.Printf("%s\n\n", header)
		}
		fmt.Println(msg)
		password, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		p := strings.TrimSpace(password)
		if len(p) != 0 {
			return p, nil
		}
	}
}

func promptError(errorInfo string, msg promptMessage, logger *types.NavituiLogger) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		ClearScreen()
		fmt.Printf("%s\n\n", errorInfo)
		fmt.Println(msg)
		password, err := reader.ReadString('\n')
		if err != nil {
			logger.File("Error prompt failure: %v", err)
			return "", err
		}
		p := strings.TrimSpace(password)
		return p, nil
	}
}

func promptInt(header string, msg promptMessage) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	var i int
	for {
		ClearScreen()
		if len(header) != 0 {
			fmt.Println(header)
		}
		fmt.Println(msg)
		n, err := fmt.Fscan(reader, &i)
		if err != nil || n != 1 {
			continue
		}
		return i, nil
	}
}

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func getValidationFailureMsg(status configValidationStatus) string {
	switch status {
	case InvalidUsername:
		return "Username must not be empty"
	case InvalidServer:
		return "Server address must not be empty"
	case InvalidPort:
		return "Port must be between 1 and 65535"
	case InvalidPasswordStore:
		return "Password store must be \"keyring\" or \"prompt\""
	default:
		return "Invalid Configuration"
	}
}

func getPingFailureMsg(result types.PingResult) string {
	switch result {
	case types.PingAuthFailure:
		return "Authentication Failure"
	case types.PingServerError:
		return "Unable to reach server"
	case types.PingURLBuildFailure:
		return "Unable to build server url"
	default:
		return "Ping Failure"
	}
}
