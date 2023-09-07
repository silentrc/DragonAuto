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

var Incubators = []IncubatorInfo{}

type IncubatorInfo struct {
	IncubatorId string `json:"id"`
	Name        string `json:"name"`
}

func DragonAutoCollect() {
	//检测token
	if config.Instance.DragonAuto.Token == "" {
		ok := checkToken()
		if !ok {
			zap.L().Sugar().Errorf("token 验证失败,请检查")
			return
		}
	}

	var resultIncubator ResultDragonAutoIncubatorId
	var err error
	if len(Incubators) <= 0 {
		for {
			//获取孵化器ID
			resultIncubator, err = sendRequestDragonGetIncubatorId()
			if err != nil {
				//获取孵化器ID失败
				zap.L().Sugar().Errorf("获取孵化器ID失败，睡眠10s后重新获取 err :%v", err)
				time.Sleep(10 * time.Second)
			}
			break
		}
		if config.Instance.DragonAuto.Stars {
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.EggIncubator,
				Name:        "龙蛋/金蛋",
			})
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.SoulIncubator,
				Name:        "龙魂",
			})
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.EssenceIncubator,
				Name:        "龙精",
			})
		}
		if config.Instance.DragonAuto.Chaos {
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.BallIncubator,
				Name:        "白龙",
			})
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.CoreIncubator,
				Name:        "黑龙",
			})
		}
		if config.Instance.DragonAuto.Saint {
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.ScaleIncubator,
				Name:        "龙鳞",
			})
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.BloodIncubator,
				Name:        "龙血",
			})
			Incubators = append(Incubators, IncubatorInfo{
				IncubatorId: resultIncubator.Data.GlassIncubator,
				Name:        "龙晶",
			})
		}
	}

	for _, v := range Incubators {
		res, err := sendRequestDragonList(v.IncubatorId, v.Name)
		if err == nil {
			time.Sleep(time.Second * 5)
			sendRequestDragonCollect(v.IncubatorId, v.Name, res.Data.List[0].DragonEgg, res.Data.List[0].DragonGold)
		}

	}

}

func sendRequestDragonCollect(id, name string, num float64, goldNum float64) {
	isSelect := false
	var option string
Select:
	if isSelect {
		option = "collect"
	} else {
		option = "startIncubation"
	}
	res, err := NewAppCommon().RequestDragonCollectOrStart(id, option)
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
		if result.Message == "此孵化园已处于孵化中" {
			isSelect = true
			goto Select
		}
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
	if isSelect {
		return
	}
	isSelect = true
	goto Select
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
	restartNum := 0
	zap.L().Sugar().Infof("sendRequestDragonList %v", name)
Restart:
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

	if result.Code == 10040 && result.Message == "登录已过期请重新登录" {
		config.Instance.DragonAuto.Token = ""
		if restartNum > 3 {
			zap.L().Sugar().Errorf("重新登录三次失败停止本次采集,name%v, data:%v err :%v", name, string(res), err)
			err = errors.New(fmt.Sprintf("重新登录三次失败停止本次采集,name%v, data:%v err :%v", name, string(res), err))
			return
		}
		time.Sleep(time.Second * 5)
		checkToken()
		restartNum++
		zap.L().Sugar().Errorf("登录已过期,开始重新登录%v, data:%v err :%v", name, string(res), err)
		goto Restart
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
func sendRequestDragonStartIncubation(id, name string) (result ResultDragonList, err error) {
	return
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
	if config.Instance.DragonAuto.Account == "" || config.Instance.DragonAuto.Pwd == "" {
		zap.L().Sugar().Errorf("账号或密码为空，如需使用token模式，请把mode改为2")
		return false
	}
	sendRequestDragonLogin()
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

type ResultDragonAutoIncubatorId struct {
	Code int `json:"code"`
	Data struct {
		EggIncubator     string `json:"incubatorId1"`
		SoulIncubator    string `json:"incubatorId2"`
		EssenceIncubator string `json:"incubatorId3"`
		BallIncubator    string `json:"incubatorId5"`
		CoreIncubator    string `json:"incubatorId6"`
		ScaleIncubator   string `json:"incubatorId7"`
		BloodIncubator   string `json:"incubatorId8"`
		GlassIncubator   string `json:"incubatorId9"`
	} `json:"data"`
	Message string `json:"message"`
}

func sendRequestDragonGetIncubatorId() (result ResultDragonAutoIncubatorId, err error) {
	res, err := NewAppCommon().RequestDragonGetIncubatorId()
	if err != nil {
		zap.L().Sugar().Errorf("发送孵化器请求失败 err :%v", err)
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		zap.L().Sugar().Errorf("返回解析失败 data:%v err :%v", string(res), err)
		return
	}
	if result.Code != 0 {
		zap.L().Sugar().Errorf("获取孵化器ID失败 err:%v, res:%+v", err, result)
		return
	}
	return
}
