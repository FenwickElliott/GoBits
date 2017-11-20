package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "bar\n")
	})
	fmt.Println("Listening...")
	http.ListenAndServe(":7410", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo\n")
	print("foo...")
}

func print(thing string) {
	fmt.Println(thing)
}
