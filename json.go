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
	var phoneSruct phonetic
	jsonBytes := []byte(`{"a":"alpha","b":"bravo","c":"charlie"}`)
	err := json.Unmarshal(jsonBytes, &phoneSruct)
	check(err)
	fmt.Printf("%+v\n", phoneSruct)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
