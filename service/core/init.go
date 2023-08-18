package core

import (
	"dragonAuto/app"
	"dragonAuto/config"
	"dragonAuto/utils/utils"
	"go.uber.org/zap"
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
		if !config.Instance.DragonAuto.Stars && !config.Instance.DragonAuto.Chaos && !config.Instance.DragonAuto.Saint {
			zap.L().Sugar().Errorf("至少开启一种岛屿自动收取")
			return
		}
		utils.Go(startAutoCollect)
	}
}

func startAutoCollect() {
	for {
		utils.Go(app.DragonAutoCollect)
		time.Sleep(time.Duration(config.Instance.DragonAuto.CollectTime) * time.Minute)
	}

}
