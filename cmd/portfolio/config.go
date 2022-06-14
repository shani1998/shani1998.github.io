package main

import (
	"github.com/spf13/pflag"
)

const (
	defaultServerAddress = "127.0.0.1"
	defaultServerPort    = "8080"

	defaultLogFormat = "json"
	defaultLogLevel  = "info"
)

var (
	_ = pflag.String("server.address", defaultServerAddress, "Address to bind the server listener to")
	_ = pflag.String("server.port", defaultServerPort, "Port to bind the server listener to")

	_ = pflag.String("log.format", defaultLogFormat, "set the logging format: json or text")
	_ = pflag.String("log.level", defaultLogLevel, "set the logging level: debug, info, warning, error, fatal, panic")
)
