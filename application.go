package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

//const url = "http://47.242.16.235:20088/vodka/variable/list"

func main() {
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Get(DevUrl)
	if err != nil {
		fmt.Println(err)
	}
	buffer := make([]byte, 1024)
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer)
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	name := "cai"
	fmt.Println(reflect.TypeOf(name))

	//jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	//var v interface{}
	//json.Unmarshal(result.Bytes(), &v)
	//data := v.(map[string]interface{})
	//
	//for k, v := range data {
	//	switch v := v.(type) {
	//	case string:
	//		fmt.Println(k, v, "(string)")
	//	case float64:
	//		fmt.Println(k, v, "(float64)")
	//	case []interface{}:
	//		fmt.Println(k, "(array):")
	//		for i, u := range v {
	//			fmt.Println("    ", i, u)
	//		}
	//	case bool:
	//		fmt.Println(k, v, "(bool)")
	//	default:
	//		fmt.Println(k, v, "(unknown)")
	//	}
	//}
}
