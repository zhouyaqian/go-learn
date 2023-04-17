package main

import (
	"go-learn/cmd"
	"log"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Excute err: %v", err)
	}
}
