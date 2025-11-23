package main

import (
	"log"
)

func main() {
	if err := StartServer("./config.json"); err != nil {
		log.Fatal(err)
	}
}
