package jsonrpc

import "time"

type Config struct {
	Url        string
	Url2       string
	ZkEvm      string
	CardonaUrl string

	PollingInterval time.Duration
}
