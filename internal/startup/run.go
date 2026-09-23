package startup

import (
	"net"
	"net/url"
	"strconv"

	"github.com/adrg/xdg"
	"github.com/joshw34/navitui/internal/types"
)

type ConfigOptions struct {
	Cred Credentials
}

func RunStartup(getServer func(string, Credentials) (types.Server, error)) (types.Server, ConfigOptions, error) {
	configFile, err := xdg.ConfigFile("navitui/navitui.toml")
	if err != nil {
		return nil, ConfigOptions{}, err
	}

	var tomlConfig tomlConfigFile
	err = tomlConfig.parseConfig(configFile)
	if err != nil {
		return nil, ConfigOptions{}, err
	}

	if len(tomlConfig.User) != 0 {
		u, err := url.Parse(tomlConfig.Server)
		if err != nil {
			return nil, ConfigOptions{}, err
		}
		u.Host = net.JoinHostPort(u.Hostname(), strconv.Itoa(tomlConfig.Port))
		var cred Credentials
		cred.User = tomlConfig.User
		cred.Password = tomlConfig.Password
		cred.BaseUrl = u.String()
		srv, err := getServer("navidrome", cred)
		if err != nil {
			return nil, ConfigOptions{}, err
		}
		return srv, ConfigOptions{Cred: cred}, nil
	}
	return nil, ConfigOptions{}, nil
}
