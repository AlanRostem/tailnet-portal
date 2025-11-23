package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func StartServer(configPath string) error {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	config, err := NewConfig(configPath)
	if err != nil {
		return err
	}
	http.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", config.Port), nil)
}
