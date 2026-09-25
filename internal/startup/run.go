package startup

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/joshw34/navitui/internal/types"
)

type srvbuild func(string, Credentials, *types.NavituiLogger) (types.Server, error)

type ConfigOptions struct {
	Cred Credentials
}

type Credentials struct {
	BaseUrl  string
	User     string
	Password string
}

func RunStartup(getServer srvbuild, logger *types.NavituiLogger) (types.Server, ConfigOptions, error) {
	configFile, err := xdg.ConfigFile("navitui/navitui.toml")
	if err != nil { // Could not generate config file path -> return error
		return nil, ConfigOptions{}, fmt.Errorf("could not generate config file path: %v", err)
	}

	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) { // Config file doesn't exist -> run setup
			logger.File("Config file %v not found, starting setup wizard", configFile)
			return firstRun(getServer, configFile, logger)
		}
		return nil, ConfigOptions{}, fmt.Errorf("could not access config file %v: %v", configFile, err) // Config file exists, not accessible -> return error
	}

	var parsed config

	if err = parsed.parseConfig(configFile, logger); err != nil {
		_, _ = promptError("Config file '"+configFile+"' could not be parsed: "+err.Error(), ErrorSetupWizard, logger) // Choose exit for manual fix or run setup wizard
		return firstRun(getServer, configFile, logger)                                                                 // Ignore prompt error -> go to setup wizard
	}

	if result := parsed.validateConfig(logger); result == ValidationSuccess {
		for range 3 {
			ClearScreen()

			err = parsed.getPassword(logger)
			if err != nil {
				_, _ = promptError("Failed to retrieve password", ErrorTryAgain, logger) //Ignore prompt error -> try again
				continue
			}

			srv, res := parsed.tryPing(getServer, logger)
			if srv != nil && res == types.PingSuccess {
				return srv, ConfigOptions{}, nil // Successful ping --> return server
			}

			if res == types.PingAuthFailure && parsed.PasswordStore == "keyring" {
				break
			}
			_, _ = promptError(getPingFailureMsg(res), ErrorTryAgain, logger) // Ignore prompt error -> try again
		}
	}
	_, _ = promptError("Unable to connect with current configuration", ErrorSetupWizard, logger) // Ignore prompt error -> go to setup wizard
	return firstRun(getServer, configFile, logger)
}
