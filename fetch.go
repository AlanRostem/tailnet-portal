package main

import (
	"context"
	"strings"

	"tailscale.com/client/tailscale/v2"
)

func FetchDevices(config *Config) ([]tailscale.Device, error) {
	client := &tailscale.Client{
		Tailnet: config.Tailnet,
		APIKey:  config.APIKey,
	}
	devices, err := client.Devices().List(context.Background())
	if err != nil {
		return nil, err
	}
	for i, dev := range devices {
		for _, ex := range config.ExcludedDeviceNames {
			if strings.Contains(dev.Name, ex) {
				devices = append(devices[:i], devices[i+1:]...)
			}
		}
	}
	return devices, nil
}
