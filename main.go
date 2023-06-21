package main

import (
	"ACME/secscan/cmd/secscancli"
	"fmt"
)

func main() {
	err := secscancli.Execute()

	if err != nil {
		fmt.Println("Try secscan help")
	}
}
