package main

import "github.com/mrbelka12000/mock_server/cmd/server"

func main() {
	go func() {
		server.Run()
	}()

	<-make(chan bool)
}
