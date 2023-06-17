package app

import (
	"dragonAuto/config"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

type ResultDragonAutoCollect struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    int    `json:"data"`
}

func DragonAutoCollect() {
	sendRequest(config.Instance.DragonAuto.IncubatorId1, "龙蛋")
	time.Sleep(time.Second * 5)
	sendRequest(config.Instance.DragonAuto.IncubatorId2, "龙魂")
	time.Sleep(time.Second * 5)
	sendRequest(config.Instance.DragonAuto.IncubatorId3, "龙精")
}

func sendRequest(id, name string) {
	res, err := NewAppCommon().RequestDragonCollect(id)
	if err != nil {
		zap.L().Sugar().Errorf("发送收集请求失败 err :%v", err)
		return
	}
	result := ResultDragonAutoCollect{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		zap.L().Sugar().Errorf("返回解析失败 err :%v", err)
		return
	}
	if result.Code != 0 {
		if result.Message == "暂无龙蛋可收取" {
			zap.L().Sugar().Infof("收集失败，暂无龙蛋可收取%v", result)
			return
		}
		zap.L().Sugar().Errorf("收集失败 err :%+v", result)
		return
	}
	zap.L().Sugar().Infof("收集成功 本次收集%v个%v", result.Data, name)
}
