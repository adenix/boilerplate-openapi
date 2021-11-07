package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/adenix/openapi-boilerplate/server"
)

func main() {
	var port = flag.Int("port", 8080, "port where to serve traffic")

	if err := server.Run(fmt.Sprintf("0.0.0.0:%d", *port)); err != nil {
		log.Fatal("Failed to start service", err)
	}
}
