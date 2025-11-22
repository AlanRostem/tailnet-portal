package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := NewConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	devices, err := FetchDevices(config)
	if err != nil {
		log.Fatal(err)
	}
	for _, dev := range devices {
		fmt.Println(dev.Name)
	}
}
