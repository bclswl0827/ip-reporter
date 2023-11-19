package main

import (
	"encoding/json"
	"time"
)

func reportNetworkAddrs(timeout int, deviceTag, reportURI string, ipAddress string) error {
	payload, err := json.Marshal(map[string]any{
		"ip_address": ipAddress,
		"device_tag": deviceTag,
	})
	if err != nil {
		return err
	}

	_, err = sendPostRequest(
		reportURI, string(payload), "application/json",
		time.Second*time.Duration(timeout), time.Second,
		3, false, nil,
	)
	if err != nil {
		return err
	}

	return nil
}
