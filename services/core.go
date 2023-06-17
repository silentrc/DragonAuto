package services

import (
	"dragonAuto/config"
	"dragonAuto/service/core"
)

var CoreService = new(coreService)

type coreService struct {
	conf config.ConfigValue
}

func (work *coreService) Init() {

}

func (work *coreService) Stop() {

}

func (work *coreService) Run() {
	core.NewCore().Init()
}
