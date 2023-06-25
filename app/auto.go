package app

import (
	"dragonAuto/config"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type ResultDragonAutoCollect struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    float64 `json:"data"`
}

const (
	DragonEgg     = "龙蛋"
	DragonSoul    = "龙魂"
	DragonEssence = "龙精"
)

func DragonAutoCollect() {
	//检测token
	ok := checkToken()
	if !ok {
		zap.L().Sugar().Errorf("token 验证失败,请检查")
		return
	}
	res, err := sendRequestDragonList(config.Instance.DragonAuto.IncubatorId1, DragonEgg)
	if err == nil {
		time.Sleep(time.Second * 5)
		sendRequestDragonCollect(config.Instance.DragonAuto.IncubatorId1, DragonEgg, res.Data.List[0].DragonEgg, res.Data.List[0].DragonGold)
	}
	time.Sleep(time.Second * 5)
	res, err = sendRequestDragonList(config.Instance.DragonAuto.IncubatorId2, DragonSoul)
	if err == nil {
		time.Sleep(time.Second * 5)
		sendRequestDragonCollect(config.Instance.DragonAuto.IncubatorId2, DragonSoul, res.Data.List[0].DragonSoul, 0)
	}
	time.Sleep(time.Second * 5)
	res, err = sendRequestDragonList(config.Instance.DragonAuto.IncubatorId3, DragonEssence)
	if err == nil {
		time.Sleep(time.Second * 5)
		sendRequestDragonCollect(config.Instance.DragonAuto.IncubatorId3, DragonEssence, res.Data.List[0].DragonEssence, 0)
	}
}

func sendRequestDragonCollect(id, name string, num float64, goldNum float64) {
	res, err := NewAppCommon().RequestDragonCollect(id)
	if err != nil {
		zap.L().Sugar().Errorf("发送收集请求失败 err :%v", err)
		return
	}
	result := ResultDragonAutoCollect{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		zap.L().Sugar().Errorf("返回解析失败 data:%v err :%v", string(res), err)
		return
	}
	if result.Code != 0 {
		if result.Message == "暂无龙蛋可收取" || result.Message == "暂无资源可收取" {
			zap.L().Sugar().Infof("收集失败，暂无资源可收取 result:%+v", string(res))
			return
		}
		zap.L().Sugar().Errorf("收集失败 err:%v, res:%+v", err, result)
		return
	}
	zap.L().Sugar().Infof("收集成功 本次收集%v个%v,result:%+v", num, name, string(res))
	if goldNum != 0 {
		zap.L().Sugar().Infof("收集成功 本次收集%v个金蛋,result:%+v", num, string(res))
	}
}

type ResultDragonList struct {
	Code int `json:"code"`
	Data struct {
		List []DragonListResult `json:"list"`
	} `json:"data"`
	Message string `json:"message"`
}

type DragonListResult struct {
	Durability    float64 `json:"durability"`         //当前耐久
	DragonEgg     float64 `json:"eggDepositNum"`      //龙蛋
	DragonEssence float64 `json:"eggDepositEnergy"`   //龙精
	DragonSoul    float64 `json:"eggDepositJindou"`   //龙魂
	DragonGold    float64 `json:"eggDepositMaterial"` //金蛋
	IncubatorType float64 `json:"incubatorType"`      //孵化器类型
}

func sendRequestDragonList(id, name string) (result ResultDragonList, err error) {
	zap.L().Sugar().Infof("sendRequestDragonList %v", name)
	res, err := NewAppCommon().RequestDragonList(id)
	if err != nil {
		zap.L().Sugar().Errorf("发送收集%v请求失败 err :%v", name, err)
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		zap.L().Sugar().Errorf("返回解析失败%v, data:%v err :%v", name, string(res), err)
		return
	}
	if result.Code == 10400 && result.Message == "登录已过期请重新登录" {
		zap.L().Sugar().Errorf("返回解析失败%v, data:%v err :%v", name, string(res), err)
		return
	}

	if result.Code != 0 {
		zap.L().Sugar().Errorf("获取%v列表失败 :%+v", name, result)
		err = errors.New(fmt.Sprintf("获取%v列表失败 :%+v", name, result))
		return
	}
	return
}

// 列表获取当前龙
// 循环获取当前耐久
//
//	if 低于耐久预设值 {
//	    获取当前龙魂数量
//	    if 低于每次添加龙魂预设值 {
//			stop or skip
//			if stop {
//				添加全部龙魂
//			}
//	    }
//	}
//
// 开启孵化
// https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/startIncubation post {id: "447496257439551488"}
func sendRequestDragonStartIncubation() {

}

// 关闭孵化
// https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/stopIncubation post {id: "447496257439551488"}
func sendRequestDragonStopIncubation() {

}

// 添加龙魂
func sendRequestDragonAddSoul() {

}

//耐久
//https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/replenishDurabilityPre
//列表获取 post {id: "447496257439551488"}
//result replenishDurabilityMaxValue : 1200 //最大补充耐久
//result jindou : 24 当前龙魂数量

//补充龙魂
//post  https://ld.douxiangapp.com/ld/api/v1/dragon/dragonIncubator/replenishDurability
//{dragonIncubatorId: "447496257439551488", jindou: 4, replenishType: 1}
//jindou 龙魂数量
//replenishType 孵化器类型？？？ 参数固定 1

//精力
//nftId 24240 精力药水200
//nftId 24331 精力药水600
//nftId 24284 精力药水1200
//nftId 24144 精力药水3000

// 获取精力药水list
func sendRequestDragonPotionList() {

}

// 补充精力
func sendRequestDragonAddEssence() {

}

//补充精力 https://ld.douxiangapp.com/ld/api/v1/member/sysMemberNft/petProp
// post {attrId: 1, memberNftId: "1667518253500600322", propMemberNftIds: ["1670566077914099713"]}
//attrId :1
//获取龙属性信息 https://ld.douxiangapp.com/ld/api/v1/dragon/dragonCardSlot/getOneById?id=1665818363241279489
//memberNftId : "1668932059820904449" = sysMemberNftId
//propMemberNftIds : ["1670756768378966018"]

// 检测token
func checkToken() bool {
	if config.Instance.DragonAuto.Mode == 1 {
		if config.Instance.DragonAuto.ReqToken == "" {
			zap.L().Sugar().Errorf("token为空，如需使用账号密码模式，请把mode改为1")
		} else {
			config.Instance.DragonAuto.Token = config.Instance.DragonAuto.ReqToken
		}
		return false
	}
	if config.Instance.DragonAuto.Mode == 2 {
		if config.Instance.DragonAuto.Account == "" || config.Instance.DragonAuto.Pwd == "" {
			zap.L().Sugar().Errorf("账号或密码为空，如需使用token模式，请把mode改为2")
			return false
		}
		sendRequestDragonLogin()
	}
	if config.Instance.DragonAuto.Token != "" {
		return true
	}
	return false
}

type ResultDragonLogin struct {
	Code int `json:"code"`
	Data struct {
		AccessToken string `json:"access_token"`
	} `json:"data"`
	Message string `json:"message"`
}

// 用户登陆
func sendRequestDragonLogin() (result ResultDragonLogin, err error) {
	res, err := NewAppCommon().RequestPasswordLogin()
	if err != nil {
		zap.L().Sugar().Errorf("用户登陆失败 err :%v", err)
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		zap.L().Sugar().Errorf("返回解析失败, data:%v err :%v", string(res), err)
		return
	}
	if result.Code != 0 {
		zap.L().Sugar().Errorf("用户登陆失败 :%+v", result)
		return
	}
	if result.Data.AccessToken == "" {
		zap.L().Sugar().Errorf("用户登陆失败 :%+v", result)
		return
	}
	config.Instance.DragonAuto.Token = "Bearer " + result.Data.AccessToken
	return
}
