package main

import (
	"encoding/json"
	"os"
)

type DeviceConfig struct {
	Hostname string
	Alias    string
}

type Config struct {
	Tailnet string
	APIKey  string
	Port    string
	Devices []DeviceConfig
}

func NewConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, stats.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
