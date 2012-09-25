package main

import "os"

func main() {
	var user = os.Getenv("USER")
	if user == "" {
		panic("no value for $USER")
	}
}
