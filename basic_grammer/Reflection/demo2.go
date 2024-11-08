package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//解析json

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonStr := `{"id":1, "name":"yss", "age":18}`
	var user User
	// func Unmarshal(data []byte, v any) error
	//先解析，看是否报错，要是报错，给的就不是json字符串
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 使用反射获取User结构体的字段信息
	val := reflect.ValueOf(user)
	typ := reflect.TypeOf(user)

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)
		fmt.Printf("Field Name: %s, Field Value: %v, Field Type: %s\n", fieldType.Name, fieldVal, fieldType.Type)
	}
}
