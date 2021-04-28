package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type VariableMapping struct {
	Name          string                `json:"name"`
	ReturnType    string                `json:"returnType"`
	Description   string                `json:"description"`
	SourceHandler []map[string][]string `json:"sourceHandler"`
	VariableType  string                `json:"variableType"`
}

const DMLPrefix = "INSERT INTO vaas.variable_mapping(mapping) VALUES "

func GenerateDML(variables []Variable) {
	mappings := buildMappings(variables)
	indents := make([]string, 0, 10)

	// marshal slice
	for _, mapping := range mappings {
		indent, _ := json.MarshalIndent(mapping, "", "\t")
		indents = append(indents, "('"+string(indent)+"')")
	}

	join := strings.Join(indents, ", ")

	dmlSQL := DMLPrefix + join + ";"

	file := "mapping_dml_" + time.Now().Format("2006_01_02") + ".sql"
	if isExist(file) {
		fmt.Printf("File: %s has been existed.\n If you want to regenerate, please remove the pre file.\n", file)
		return
	}
	err := ioutil.WriteFile(file, []byte(dmlSQL), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func buildMappings(variables []Variable) []VariableMapping {
	var mappings []VariableMapping

	for _, variable := range variables {
		sourceHandlers := []map[string][]string{{variable.Lower: make([]string, 0)}}
		mappings = append(mappings, VariableMapping{
			Name:          variable.Var,
			ReturnType:    variable.Type,
			Description:   "变量:" + variable.Name,
			SourceHandler: sourceHandlers,
			VariableType:  "numerical",
		})
	}
	return mappings
}
