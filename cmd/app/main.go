package main

import (
	"log"

	"github.com/mrbelka12000/mock_server/cmd/client"
	"github.com/mrbelka12000/mock_server/cmd/server"
)

func main() {
	go func() {
		client.Run()
	}()

	if err := server.Run(); err != nil {
		log.Println(err)
		return
	}
}
