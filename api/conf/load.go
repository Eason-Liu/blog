package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var config *Config //为了该变量不被外部修改，未暴露

func Conf() *Config {
	if config == nil {
		panic("load config first")
	}
	return config
}

func LoadConfigFromEnv() error {
	config = NewDefaultConfig()

	if err := env.Parse(config); err != nil {
		return err
	}
	return nil
}

func LoadConfigFromToml(filepath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filepath, config)
	if err != nil {
		return err
	}
	return nil
}
