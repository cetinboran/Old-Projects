package main

import (
	"log"

	"github.com/cetinboran/ssher/core"
	inputhandler "github.com/cetinboran/ssher/input"
)

func main() {
	handler := inputhandler.NewInputHandler()

	userOptions, err := handler.Start()
	if err != nil {
		log.Fatal(err)
	}

	server := inputhandler.Get(userOptions, "-s").Input()
	port := inputhandler.Get(userOptions, "--port").Input()
	bruteForce := core.NewBruteForce(server, port, userOptions)

	if err := bruteForce.Start(); err != nil {
		log.Fatal(err)
	}
}
