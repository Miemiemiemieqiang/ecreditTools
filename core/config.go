package core

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	MainConfig *MainConfig
}

type JSONConfig struct {
	Path string
}

func (J *JSONConfig) GetConfig() *Config {
	buf, err := ioutil.ReadFile(J.Path)
	if err != nil {
		panic(err)
	}
	var config Config
	var mainConfig MainConfig
	if json.Unmarshal(buf, &mainConfig) != nil {
		panic(err)
	} else {
		config.MainConfig = &mainConfig
	}
	return &config
}

type MainConfig struct {
	TargetPath   string
	CsvPath      string
	Author       string
	TemplatePath string
}

func GetDefaultConfig() *Config {
	config := JSONConfig{Path: "config/config2.json"}
	return config.GetConfig()
}

func GetConfig(path string) *Config {
	config := JSONConfig{Path: path}
	return config.GetConfig()
}
