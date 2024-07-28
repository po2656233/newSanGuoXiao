package category

import (
	exViper "github.com/po2656233/superplace/extend/viper"
	log "github.com/po2656233/superplace/logger"
)

var YamlObj = new(ConfSetting)
var vp = exViper.NewViper("config/gameconf/game.yaml")

func InitConfig() {
	err := vp.Unmarshal(&YamlObj)
	if err != nil {
		panic(err)
	}
	//vp.ToToml()
	//vp.ToJson()
	log.Infof("conf chipsxxxx info:%v ok", YamlObj.Chips)
}

type Duration struct {
	FreeTime  int32 `yaml:"free"`
	StartTime int32 `yaml:"start"`
	PlayTime  int32 `yaml:"play"`
	OpenTime  int32 `yaml:"open"`
	OverTime  int32 `yaml:"over"`
}
type FightDuration struct {
	Duration
	ReadyTime  int32 `yaml:"ready"`
	DealTime   int32 `yaml:"deal"`
	CallTime   int32 `yaml:"call"`
	DoubleTime int32 `yaml:"double"`
}
type BattleDuration struct {
	Duration
	ReadyTime int32 `yaml:"ready"`
	DealTime  int32 `yaml:"deal"`
	CompTime  int32 `yaml:"compare"`
}

type MahjongDuration struct {
	Duration
	DirectTime  int32 `yaml:"direct"`
	DecideTime  int32 `yaml:"decide"`
	RollTime    int32 `yaml:"roll"`
	DealTime    int32 `yaml:"deal"`
	OperateTime int32 `yaml:"operate"`
}

type MjXLCHDuration struct {
	MahjongDuration
	ChangeTime  int32 `yaml:"change"`
	DingQueTime int32 `yaml:"dingque"`
	TingTime    int32 `yaml:"ting"`
	HuaZhuTime  int32 `yaml:"huazhu"`
}

type ChineseChessDuration struct {
	Duration
	SetTime     int32 `yaml:"settime"`
	ConfirmTime int32 `yaml:"confirm"`
}

type ConfSetting struct {
	//---------筹码-------------
	Chips struct {
		General []int32 `yaml:"general,flow"`
		Middle  []int32 `yaml:"middle,flow"`
		High    []int32 `yaml:"high,flow"`
		Other   []int32 `yaml:"other,flow"`
	} ` yaml:"Chips"`

	//---------百人类游戏基本配置-------------
	//百家乐
	Baccarat struct {
		Duration  ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Baccarat"`
	//百人牛牛
	Brcowcow struct {
		Duration    ` yaml:"duration"`
		Inventory   int64 `yaml:"inventory"`
		BankerScore int64 `yaml:"bankerScore"`
		MaxHost     uint8 `yaml:"maxHost"`
	} ` yaml:"Brcowcow"`
	//百人骰宝
	Brtoubao struct {
		Duration  ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Brtoubao"`
	//百人推筒子
	Brtuitongzi struct {
		Duration  ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Brtuitongzi"`
	//推筒子上庄喝水
	BrtuitongziSZHS struct {
		Duration  ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"BrtuitongziSZHS"`
	//龙虎斗
	TigerXdragon struct {
		Duration  ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"TigerXdragon"`

	//---------对战类游戏基本配置--------------------

	//斗地主
	Landlord struct {
		FightDuration ` yaml:"duration"`
		MinPerson     int32 `yaml:"minPerson"`
		MaxPerson     int32 `yaml:"maxPerson"`
	} ` yaml:"Landlord"`
	//跑得快
	PaoDeKuai struct {
		FightDuration ` yaml:"duration"`
		MinPerson     int32 `yaml:"minPerson"`
		MaxPerson     int32 `yaml:"maxPerson"`
	} ` yaml:"PaoDeKuai"`
	//扎金花
	Zhajinhua struct {
		BattleDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
		LimitScore     int64 `yaml:"limitScore"`
	} ` yaml:"Zhajinhua"`
	//通比牛牛
	Tbcowcow struct {
		BattleDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
	} ` yaml:"Tbcowcow"`
	//抢庄牛牛
	Qzcowcow struct {
		BattleDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
	} ` yaml:"Qzcowcow"`
	//三公
	Sangong struct {
		BattleDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
	} ` yaml:"Sangong"`
	//极速扎金花
	ZhajinhuaJiSu struct {
		BattleDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
	} ` yaml:"ZhajinhuaJiSu"`
	//推筒子
	Tuitongzi struct {
		BattleDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
	} ` yaml:"Tuitongzi"`

	//所有麻将配置
	//国标麻将
	Mahjong struct {
		MahjongDuration ` yaml:"duration"`
		MinPerson       int32 `yaml:"minPerson"`
		MaxPerson       int32 `yaml:"maxPerson"`
	} ` yaml:"Mahjong"`
	//二人麻将
	MahjongER struct {
		MahjongDuration ` yaml:"duration"`
		MinPerson       int32 `yaml:"minPerson"`
		MaxPerson       int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongER"`
	//长沙麻将
	MahjongCS struct {
		MahjongDuration ` yaml:"duration"`
		MinPerson       int32 `yaml:"minPerson"`
		MaxPerson       int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongCS"`
	//广东麻将
	MahjongGD struct {
		MahjongDuration ` yaml:"duration"`
		MinPerson       int32 `yaml:"minPerson"`
		MaxPerson       int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongGD"`
	//四川麻将
	MahjongSC struct {
		MahjongDuration ` yaml:"duration"`
		MinPerson       int32 `yaml:"minPerson"`
		MaxPerson       int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongSC"`
	//血流成河
	MahjongXLCH struct {
		MjXLCHDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongXLCH"`
	//血战到底
	MahjongXZDD struct {
		MjXLCHDuration ` yaml:"duration"`
		MinPerson      int32 `yaml:"minPerson"`
		MaxPerson      int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongXZDD"`
	// 三国消
	SanGuoXiao struct {
		Duration ` yaml:"duration"`
		Col      int32 `yaml:"col"`
		Row      int32 `yaml:"row"`
	} ` yaml:"SanGuoXiao"`
	// 象棋
	ChineseChess struct {
		ChineseChessDuration `yaml:"duration"`
		Col                  int32 `yaml:"col"`
		Row                  int32 `yaml:"row"`
	} ` yaml:"ChineseChess"`
}
