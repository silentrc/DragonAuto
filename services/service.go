package services

import (
	_interface "dragonAuto/services/interface"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var (
	services   []_interface.Service // 服务列表
	serverName = ""                 // 服务名称
)

func Run(name string, servs ..._interface.Service) {
	serverName = name
	// append to the services list
	for i := 0; i < len(servs); i++ {
		servs[i].Init()
		services = append(services, servs[i])
	}

	for _, server := range services {
		go server.Run()
	}

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-c
	zap.L().Debug(fmt.Sprintf("%s closing down (signal: %v)", serverName, sig))
	Stop()

}

func Stop() {
	for _, server := range services {
		server.Stop()
	}
}
