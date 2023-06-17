package core

import (
	"dragonAuto/app"
	"dragonAuto/config"
	"dragonAuto/utils/utils"
	"time"
)

type core struct {
}

func NewCore() *core {
	return &core{}
}

// 开始核心逻辑
func (w *core) Init() {
	w.Run()
}

func (w *core) Run() {
	if config.Instance.DragonAuto.Enable {
		utils.Go(startAutoCollect)
	}
}

func startAutoCollect() {
	for {
		utils.Go(app.DragonAutoCollect)
		time.Sleep(time.Duration(config.Instance.DragonAuto.CollectTime) * time.Minute)
	}

}
