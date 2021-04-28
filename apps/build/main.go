package main

import (
	"ecreditTools/core"
	"fmt"
)

func main() {
	//config := core.GetConfig("config/config2.json")
	//
	//// main config
	//mainConfig := config.MainConfig
	//variables := core.GetVariables(mainConfig)
	//
	////generate SourceHandler
	//core.GenerateSourceHandler(variables, mainConfig.TargetPath, mainConfig.TemplatePath)
	//
	//// generate DML
	////core.GenerateDML(variables)
	//
	//fmt.Println("Done")

	context := core.NewContext("config/config2.json")
	context.GenerateSourceHandler()
	context.GenerateDML()
	fmt.Println("Done")


}
