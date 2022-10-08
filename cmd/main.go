package main

import (
	"log"
	"video-conferencing/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
