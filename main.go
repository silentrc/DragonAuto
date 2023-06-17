package main

import (
	log "dragonAuto/service/zap"
	"dragonAuto/services"
	"flag"
)

func init() {
	// config init
	flag.Parse()
}

func main() {
	log.Init()
	services.Run(
		"dragonAuto",
		services.CoreService,
	)
}
