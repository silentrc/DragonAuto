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

func (a *appRequest) RequestDragonGetIncubatorId() (result []byte, err error) {
	//https://ld.douxiangapp.com/ld/api/v1/member/sysMember/getMemberDetail
	resp, err := utils.NewHttpUtils().Client().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		SetHeader("Connection", "keep-alive").
		SetHeader("content-Type", "application/json").
		SetHeader("content-Type", "application/json").
		SetHeader("Host", "ld.douxiangapp.com").
		SetHeader("platform", config.Instance.DragonAuto.Platform).
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41").
		SetHeader("Authorization", config.Instance.DragonAuto.Token).
		Get("https://ld.douxiangapp.com/ld/api/v1/member/sysMember/getMemberDetail")
	if err != nil {
		zap.L().Sugar().Errorf("resp get RequestDragonGetIncubatorId err :%v", err)
		return
	}
	if resp.StatusCode() != 200 {
		zap.L().Sugar().Errorf("request  RequestDragonGetIncubatorId error:%v", err)
		return
	}
	if resp.Body() == nil {
		err = errors.New(fmt.Sprintf("request  RequestDragonGetIncubatorId error:%v", err))
		return
	}
	return resp.Body(), nil
}

func (a *appRequest) RequestDragonCollectOrStart(id, option string) (result []byte, err error) {
	resp, err := utils.NewHttpUtils().Client().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		SetHeader("Connection", "keep-alive").
		SetHeader("content-Type", "application/json").
		SetHeader("content-Type", "application/json").
		SetHeader("Host", "ld.douxiangapp.com").
		SetHeader("platform", config.Instance.DragonAuto.Platform).
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41").
		SetHeader("Authorization", config.Instance.DragonAuto.Token).
		SetBody(map[string]string{"id": id}).
		Post(fmt.Sprintf("https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/%v", option))
	//Post("https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/collect")
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
	resp, err := utils.NewHttpUtils().Client().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		SetHeader("Connection", "keep-alive").
		SetHeader("content-Type", "application/json").
		SetHeader("content-Type", "application/json").
		SetHeader("Host", "ld.douxiangapp.com").
		SetHeader("platform", config.Instance.DragonAuto.Platform).
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

func (a *appRequest) RequestPasswordLogin() (result []byte, err error) {
	resp, err := utils.NewHttpUtils().Client().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		SetHeader("Connection", "keep-alive").
		SetHeader("content-Type", "application/json").
		SetHeader("content-Type", "application/json").
		SetHeader("Host", "ld.douxiangapp.com").
		SetHeader("platform", config.Instance.DragonAuto.Platform).
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41").
		SetHeader("Authorization", config.Instance.DragonAuto.Token).
		SetBody(map[string]interface{}{
			"account": config.Instance.DragonAuto.Account,
			"pwd":     config.Instance.DragonAuto.Pwd,
		}).
		Post(fmt.Sprintf("https://ld.douxiangapp.com/ld/api/v1/token/passwordLogin"))
	if err != nil {
		zap.L().Sugar().Errorf("resp get DragonNowPrice err :%v", err)
		return
	}
	if resp.StatusCode() != 200 {
		zap.L().Sugar().Errorf("request  PasswordLogin error:%v", err)
		return
	}
	if resp.Body() == nil {
		err = errors.New(fmt.Sprintf("request  PasswordLogin error:%v", err))
		return
	}
	return resp.Body(), nil
}
