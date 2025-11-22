package main

import (
	"context"
	"strings"

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
	for _, dev := range tsDevices {
		exclude := false
		for _, ex := range config.ExcludedDeviceNames {
			if strings.Contains(dev.Name, ex) {
				exclude = true
				break
			}
		}
		if exclude {
			continue
		}
		alias := dev.Hostname
		if cfg, ok := config.DeviceConfigs[dev.Hostname]; ok {
			alias = cfg.Alias
		}
		devices = append(devices, DeviceInfo{
			Href:     dev.Name,
			Hostname: dev.Hostname,
			Alias:    alias,
		})
	}
	return devices, nil
}
