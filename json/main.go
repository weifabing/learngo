package main

import (
	"encoding/json"
	"fmt"
)

type Base struct {
	Msg1 string `json:"msg1"`
	Msg2 string `json:"msg2"`
}

type Result struct {
	Base
	Code    int         `json:code`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	fmt.Printf("-->%+v\n", "Json 解析 测试")
	jsonDecode()

	fmt.Printf("-->%+v\n", "Json 编码 测试")
	jsonEncode()
}

func jsonEncode() {
	o := Result{
		Base: Base{
			Msg1: "this msg_1",
			Msg2: "this msg_2",
		},
		Code:    200,
		Message: "hello world",
		Data:    nil,
	}
	bytes, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Debug: %s\n", bytes)

	// 编码测试2
	obj1 := map[string]interface{}{
		"code":    200,
		"Message": "hello world",
		"Data":    []int{1, 2, 3, 4},
	}

	marshal, err := json.Marshal(obj1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Debug: %s\n", marshal)

}

// json decode
func jsonDecode() {
	s := `{ "code": 1,  "message": "hi",  "data": {"list":[1,2,3]}}`

	// 解析到结构体
	obj := Result{}
	err := json.Unmarshal([]byte(s), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Object: %+v\n", obj)

	// 解析成为通用格式
	obj1 := map[string]interface{}{}
	err = json.Unmarshal([]byte(s), &obj1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Object1: %+v\n", obj1)
}
