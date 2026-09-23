package startup

import (
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfigFile struct {
	User     string `toml:"user"`
	Server   string `toml:"server"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
}

func (t *tomlConfigFile) parseConfig(filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = toml.NewDecoder(f).Decode(t)
	return err
}

func (t *tomlConfigFile) saveConfig(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := toml.NewEncoder(f)
	return encoder.Encode(t)
}
