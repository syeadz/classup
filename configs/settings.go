package settings

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	DbPath string
}

func GetConfig() Config {
	tomlPath := os.Getenv("HOME") + "/.config/classup/config.toml"
	dat, err := os.ReadFile(tomlPath)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var cfg Config
	err = toml.Unmarshal(dat, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}
