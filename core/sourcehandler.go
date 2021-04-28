package core

import (
	"log"
	"os"
	"text/template"
)

func GenerateSourceHandler(variables []Variable, targetPath string, templatePath string) {
	temp, _ := template.ParseFiles(templatePath)
	for _, variable := range variables {
		variable := variable
		func() {
			file, err := os.OpenFile(targetPath+variable.Upper+"SourceHandler.java", os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal("open file err", err)
			}
			defer file.Close()

			_ = temp.Execute(file, variable)
		}()
	}
}
