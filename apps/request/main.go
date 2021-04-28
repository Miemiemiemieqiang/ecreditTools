package main

import (
	"ecreditTools/core"
	"ecreditTools/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type variables struct {
	Name     string `json:"name"`
	Variable string `json:"variable"`
}

type params struct {
	Val map[string]interface{}
}

type results struct {
}

type Request struct {
	OrderId       string
	EmergencyInfo string
	Contact        string
}

func DoRequest(config *core.Config, builder model.Builder, writer model.Writer) {
	vars := getVars(config)
	v, _ := json.MarshalIndent(vars, "", "  ")
	fmt.Println(string(v))

	//ch := make(chan *model.MachineReview)
	re := make(chan *model.Response, 100)
	wait := sync.WaitGroup{}
	wait.Wait()
	close(re)
	//
	//var r []Request
	//_ = csvreader.New().UnMarshalFile("data/request/contact.csv", &r)
	//
	//m := make([]map[string]interface{}, 0)
	//marshal, _ := json.Marshal(r)
	//_ = json.Unmarshal(marshal, &m)
	//
	//go buildRequest(ch, builder, m)
	//
	//go send(ch, vars, re)
	//
	//save(re , writer)
}

func save(re chan *model.Response, writer model.Writer) {
	for response := range re {
		fmt.Println(response)
		writer.Write(response)
	}
}

type ABC struct {
	Params map[string]interface{} `json:"params"`
	Variables []variables `json:"variables"`
}

func send(ch chan *model.MachineReview, vars []variables, re chan *model.Response, group *sync.WaitGroup) {
	group.Add(1)
	defer group.Done()
	for review := range ch {
		abc := &ABC{}
		marshal, _ := json.Marshal(review)
		abc.Params = make(map[string]interface{})
		abc.Params["machineReviewInput"] = string(marshal)
		abc.Params["productName"] = "pycharm"
		abc.Params["eventId"] = "1_0_17862826732_1548053505007"
		abc.Variables = vars

		url := "http://localhost:20787/vodka/variable/compute"
		method := "POST"
		bytes, _ := json.Marshal(abc)
		payload := strings.NewReader(string(bytes))
		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("partner-code", "Kreditpedia")
		req.Header.Add("partner-key", "786XjbmHO7cdYQr724MG")
		req.Header.Add("app-name", "caikaiqiang")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
		response := &model.Response{}
		_ = json.Unmarshal(body, response)
		response.OrderId = review.OrderID
		fmt.Println(123)
		//response.OrderId =
		re <- response
	}
}

func buildRequest(ch chan *model.MachineReview, builder model.Builder, params []map[string]interface{}) {
	defer close(ch)
	for _, param := range params {
		machineReview := builder.Build(param)
		ch <- machineReview
	}
}

func main() {
	builder := model.NewContactBuilder()
	writer := model.NewExcelWriter()

	DoRequest(core.GetDefaultConfig(), builder, writer)

	//// 用于拼装请求参数 中的 variables
	//vars := getVars(nil)
	//v, _ := json.Marshal(vars)
	//fmt.Println(string(v))
	//// read data     (by multi ways)
	//ch := make(chan *model.MachineReview)
	////re := make(chan *results, 20)
	//
	//var send func()
	//send = func() {
	//
	//}
	//
	//go send()
	//for i := range ch {
	//	marshal, _ := json.Marshal(i)
	//	fmt.Println(string(marshal))
	//}

	// 5 concurrent
	//for i := 0; i < 5; i++ {
	//	go sendRequest(ch, re)
	//}

	// from csv, mysql, --json file--

	// build request

	// send request  (chose env)

	// save response (by multi ways)
	// to csv , mysql
	// csv need to flatten

}

func sendRequest(ch chan *params, re chan *results) {
	for p := range ch {
		fmt.Println(p)
		re <- &results{}
	}
}

func getVars(config *core.Config) []variables {
	if config == nil {
		config = core.GetDefaultConfig()
	}

	// main config
	mainConfig := config.MainConfig
	rows := core.GetVariables(mainConfig)

	// get vars
	var vars []variables
	for _, row := range rows {
		vars = append(vars, variables{Name: row.Var, Variable: row.Var})
	}
	return vars
}
