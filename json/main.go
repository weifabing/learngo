package main

import (
	"encoding/json"
	"fmt"
	"strings"
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
	//fmt.Printf("-->%+v\n", "Json 解析 测试")
	//jsonDecode()

	//fmt.Printf("-->%+v\n", "Json 编码 测试")
	//jsonEncode()

	jsonObj := map[string]interface{}{
		"code":    200,
		"Message": "hello world",
		"Data": map[string]interface{}{
			"Name": "zhangsang",
			"Age":  30,
		},
	}

	s, _ := jsonToStruct(jsonObj, 0)
	fmt.Printf("Struct code:\n%+v\n", s)

}

// 根据json对象生成 Struct
// {"Data":{},"Message":"hello world","code":200}
func jsonToStruct(jsonobj map[string]interface{}, index int) (string, error) {
	var code = ""
	if index == 0 {
		code = fmt.Sprintf("type %s struct {", "MyStruct")
	} else {
		code = fmt.Sprintf("struct {")
	}

	//代码间隔
	var textIndent = strings.Repeat("    ", index+1)

	for key, value := range jsonobj {
		var isStruct = false
		dataType := "string"
		switch v2 := value.(type) {
		case string:
			dataType = "string"
		case int:
			dataType = "int"
		case map[string]interface{}:
			isStruct = true
			dataType, _ = jsonToStruct(v2, index+1)
		default:
			dataType = "string"
		}

		lowerKey := strings.ToLower(key)
		_key := strings.Title(lowerKey)

		if isStruct == false {
			code = fmt.Sprintf("%v\n%s%-8s %-10s `json:%s`", code, textIndent, _key, dataType, lowerKey)
		} else {
			code = fmt.Sprintf("%v\n\t%-8s %-10s ", code, _key, dataType)
		}
	}
	// 尾部代码间隔
	textIndent = strings.Repeat("    ", index)
	code = fmt.Sprintf("%v\n%s}\n", code, textIndent)
	return code, nil
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
