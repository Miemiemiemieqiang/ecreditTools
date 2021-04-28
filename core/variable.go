package core

import (
	"github.com/zhnxin/csvreader"
	"strings"
	"time"
)

type Variable struct {
	Var          string
	Name         string
	Type         string
	Extra        string
	Lower, Upper string
	Datetime     string
	Author       string
}

//GetVariables get variable[] from the path of variable info which is a csv file.
func GetVariables(config *MainConfig) []Variable {
	var variables []Variable
	_ = csvreader.New().UnMarshalFile(config.CsvPath, &variables)
	fill(variables, config)
	return variables
}

func fill(variables []Variable, config *MainConfig) {
	for i, ele := range variables {
		variables[i].Upper = bigFirstName(ele.Var)
		variables[i].Lower = lowerCamel(variables[i].Upper)
		variables[i].Datetime = time.Now().Format("2006-01-02")
		variables[i].Author = config.Author
	}
}

func lowerCamel(str string) string {
	camel := str[0]
	if camel <= 'Z' && camel >= 'A' {
		camel += 32
	}
	return string(camel) + str[1:]
}

func bigFirstName(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				nz := vv[i]
				if nz <= 'z' && nz >= 'a' {
					vv[i] -= 32
				}
				//vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}
