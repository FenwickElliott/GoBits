package main

import (
	"os/exec"
)

func main() {
	send("Title", "Subtitle", "Body")
}

func send(title, subtitle, message string) {
	args := `-e display notification "` + message + `" with title "` + title + `" subtitle "` + subtitle + `"`
	err := exec.Command("osascript", args).Run()
	if err != nil {
		panic(err)
	}
}
