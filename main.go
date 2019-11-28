package main

import (
	"quotation/server"
)

func main() {
	stopped := make(chan bool)
	go server.StartServer()
	stopped<-true
}