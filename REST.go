package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://api.github.com/users/fenwickelliott/repos", nil)
	check(err)

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Println(string(bodyBytes))
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
