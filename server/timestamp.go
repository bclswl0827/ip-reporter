package main

import (
	"time"
)

func getTimestamp(utc bool) int64 {
	if utc {
		return int64(time.Now().UTC().UnixMilli())
	}

	return int64(time.Now().UnixMilli())
}
