package main

import "github.com/unawaretub86/leal-challenge/cmd/server"

func main() {
	initiator := server.NewInitiator()
	initiator.InitAll()
}
