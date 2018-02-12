package main

import (
	"encoding/json"
	"fmt"
)

type phonetic struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func main() {
	jsonBytes := []byte(`{"a":"alpha","b":"bravo","c":"charlie"}`)

	fmt.Print("Original json: ")
	fmt.Println(string(jsonBytes))

	var phoneStruct phonetic
	err := json.Unmarshal(jsonBytes, &phoneStruct)
	check(err)

	fmt.Print("Original json as struct: ")
	fmt.Printf("%+v\n", phoneStruct)

	phoneStruct.A = "apple"
	phoneStruct.B = "ballerina"
	phoneStruct.C = "catastrophe"

	fmt.Print("Edited struct: ")
	fmt.Printf("%+v\n", phoneStruct)

	editedJson, err := json.Marshal(phoneStruct)
	check(err)

	fmt.Print("Edited json: ")
	fmt.Println(string(editedJson))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
