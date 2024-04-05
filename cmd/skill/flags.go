package main

import (
	"flag"
	"os"
)

var FlagRunAddr string

func ParseFlags() {

	flag.StringVar(&FlagRunAddr, "a", "localhost:8080", "host and port to run server")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}

}
