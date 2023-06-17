package app

import (
	"dragonAuto/config"
	"dragonAuto/utils/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type appRequest struct {
}

func NewAppCommon() *appRequest {
	return &appRequest{}
}

func (a *appRequest) RequestDragonCollect(id string) (result []byte, err error) {
	resp, err := utils.NewUtils().NewHttpUtils().Client().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		SetHeader("Connection", "keep-alive").
		SetHeader("content-Type", "application/json").
		SetHeader("content-Type", "application/json").
		SetHeader("Host", "ld.douxiangapp.com").
		SetHeader("platform", "h5").
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41").
		SetHeader("Authorization", config.Instance.DragonAuto.Token).
		SetBody(map[string]string{"id": id}).
		Post("https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/collect")
	if err != nil {
		zap.L().Sugar().Errorf("resp get DragonNowPrice err :%v", err)
		return
	}
	if resp.StatusCode() != 200 {
		zap.L().Sugar().Errorf("request  DragonCollect error:%v", err)
		return
	}
	if resp.Body() == nil {
		err = errors.New(fmt.Sprintf("request  DragonCollect error:%v", err))
		return
	}
	return resp.Body(), nil
}

func (a *appRequest) RequestDragonList(id string) (result []byte, err error) {
	resp, err := utils.NewUtils().NewHttpUtils().Client().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		SetHeader("Connection", "keep-alive").
		SetHeader("content-Type", "application/json").
		SetHeader("content-Type", "application/json").
		SetHeader("Host", "ld.douxiangapp.com").
		SetHeader("platform", "h5").
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41").
		SetHeader("Authorization", config.Instance.DragonAuto.Token).
		Get(fmt.Sprintf("https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/list?id=%v", id))
	if err != nil {
		zap.L().Sugar().Errorf("resp get DragonNowPrice err :%v", err)
		return
	}
	if resp.StatusCode() != 200 {
		zap.L().Sugar().Errorf("request  DragonCollect error:%v", err)
		return
	}
	if resp.Body() == nil {
		err = errors.New(fmt.Sprintf("request  DragonCollect error:%v", err))
		return
	}
	return resp.Body(), nil
}
