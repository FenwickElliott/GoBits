package main

import "net/http"

func main() {
	// assuming root of site is "./static"
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3456", nil)
}
