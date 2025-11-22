package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func StartServer() error {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		config, err := NewConfig("./config.json")
		if err != nil {
			log.Fatal(err)
		}
		devices, err := FetchDevices(config)
		if err != nil {
			log.Fatal(err)
		}
		marshaled, err := json.Marshal(devices)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(marshaled))
	})
	return http.ListenAndServe("0.0.0.0:5000", nil)
}
