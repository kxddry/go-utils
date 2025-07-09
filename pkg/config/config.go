package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

// Config resembles any struct you can possibly use for your files.
// Note that you must use structure tags in order for the utility to work properly.
type Config interface{}

func MustParseConfig() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	return MustLoadByPath(path)
}

func MustLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exist " + path)
	}
	var res Config
	if err := cleanenv.ReadConfig(path, &res); err != nil {
		panic(err)
	}
	return &res
}

// fetchConfigPath parses the config path from flags or env and returns it.
// It prioritizes flags over env.
// Default return value: empty string
func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res != "" {
		return res
	}
	env := os.Getenv("CONFIG_PATH")
	return env
}
