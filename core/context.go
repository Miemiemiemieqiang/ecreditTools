package core

import (
	"encoding/json"
	"io/ioutil"
)

type Context struct {
	config *Config
}

func NewContextDefault() *Context {
	return NewContext("config/config.json")
}

func NewContext(configPath string) *Context {
	buf, err := ioutil.ReadFile(configPath)
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

	context := Context{config: &config}
	return &context
}

func (c Context) GetConfig() *Config {
	return c.config
}

func (c Context) GenerateSourceHandler() {
	mainConfig := c.config.MainConfig
	variables := GetVariables(mainConfig)
	GenerateSourceHandler(variables, mainConfig.TargetPath, mainConfig.TemplatePath)
}

func (c Context) GenerateDML() {
	mainConfig := c.config.MainConfig
	variables := GetVariables(mainConfig)
	GenerateDML(variables)
}