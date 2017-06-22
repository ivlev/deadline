package main

import (
	"os"
)

func main() {
	file, err := os.Create("init.dl")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("test")
}