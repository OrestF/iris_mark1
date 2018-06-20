package main

import (
	"./app/server"
)

func main() {
	server := server.New()
	server.Start()
}
