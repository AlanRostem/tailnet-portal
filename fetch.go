package main

import (
	"context"

	"tailscale.com/client/tailscale/v2"
)

type DeviceInfo struct {
	Hostname string
	Href     string
	Alias    string
}

func FetchDevices(config *Config) ([]DeviceInfo, error) {
	client := &tailscale.Client{
		Tailnet: config.Tailnet,
		APIKey:  config.APIKey,
	}
	tsDevices, err := client.Devices().List(context.Background())
	if err != nil {
		return nil, err
	}
	devices := make([]DeviceInfo, 0)

	for _, devConfig := range config.Devices {
		var dev *tailscale.Device = nil
		var hostname = devConfig.Hostname
		for _, tsDev := range tsDevices {
			if tsDev.Hostname == hostname {
				dev = &tsDev
				break
			}
		}
		if dev == nil {
			continue
		}
		alias := dev.Hostname
		if devConfig.Alias != "" {
			alias = devConfig.Alias
		}
		devices = append(devices, DeviceInfo{
			Href:     dev.Name,
			Hostname: dev.Hostname,
			Alias:    alias,
		})
	}
	return devices, nil
}
