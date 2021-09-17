package microbackendfilm

import (
	"log"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Database database `toml:"database"`
	Redis    redis    `toml:"redis"`
	Servers  map[string]server
	Services map[string]service
}

type database struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

type redis struct {
	Host string
	Pass string
	Db   string
}

type server struct {
	Port string
}

type service struct {
	Prot string
	Port string
}

func Config() TomlConfig {
	var config TomlConfig
	if _, err := toml.DecodeFile("../../../config.toml", &config); err != nil {
		log.Fatal(err)
	}

	return config
}
