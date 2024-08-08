package category

import (
	exViper "github.com/po2656233/superplace/extend/viper"
	log "github.com/po2656233/superplace/logger"
)

var YamlObj = new(ConfSetting)

func InitConfig() {
	var vp = exViper.NewViper("config/gameconf/game.yaml")
	err := vp.Unmarshal(&YamlObj)
	if err != nil {
		panic(err)
	}
	vp.ToToml()
	vp.ToJson()
	log.Infof("conf chipsxxxx info:%v ok", YamlObj.Chips)
}

type ConfSetting struct {
	Chips struct {
		General []int32 `toml:"general"`
		High    []int32 `toml:"high"`
		Middle  []int32 `toml:"middle"`
		Other   []int32 `toml:"other"`
	} `toml:"chips"`
	Duration struct {
		Free  int32 `toml:"free"`
		Open  int32 `toml:"open"`
		Over  int32 `toml:"over"`
		Play  int32 `toml:"play"`
		Start int32 `toml:"start"`
	} `toml:"duration"`
	Chinesechess struct {
		Col      int32 `toml:"col"`
		Row      int32 `toml:"row"`
		Duration struct {
			Confirm int32 `toml:"confirm"`
			Free    int32 `toml:"free"`
			Open    int32 `toml:"open"`
			Over    int32 `toml:"over"`
			Play    int32 `toml:"play"`
			Ready   int32 `toml:"ready"`
			Settime int32 `toml:"settime"`
			Start   int32 `toml:"start"`
		} `toml:"duration"`
	} `toml:"chinesechess"`
	Landlord struct {
		Maxperson int32 `toml:"maxperson"`
		Minperson int32 `toml:"minperson"`
		Duration  struct {
			Call   int32 `toml:"call"`
			Deal   int32 `toml:"deal"`
			Double int32 `toml:"double"`
			Free   int32 `toml:"free"`
			Open   int32 `toml:"open"`
			Over   int32 `toml:"over"`
			Play   int32 `toml:"play"`
			Ready  int32 `toml:"ready"`
			Start  int32 `toml:"start"`
		} `toml:"duration"`
	} `toml:"landlord"`
	Mahjong struct {
		Maxperson int32 `toml:"maxperson"`
		Minperson int32 `toml:"minperson"`
		Duration  struct {
			Deal    int32 `toml:"deal"`
			Decide  int32 `toml:"decide"`
			Direct  int32 `toml:"direct"`
			Free    int32 `toml:"free"`
			Open    int32 `toml:"open"`
			Operate int32 `toml:"operate"`
			Over    int32 `toml:"over"`
			Play    int32 `toml:"play"`
			Roll    int32 `toml:"roll"`
			Start   int32 `toml:"start"`
		} `toml:"duration"`
	} `toml:"mahjong"`
	Mahjonger struct {
		Maxperson int32 `toml:"maxperson"`
		Minperson int32 `toml:"minperson"`
		Duration  struct {
			Deal    int32 `toml:"deal"`
			Decide  int32 `toml:"decide"`
			Direct  int32 `toml:"direct"`
			Free    int32 `toml:"free"`
			Open    int32 `toml:"open"`
			Operate int32 `toml:"operate"`
			Over    int32 `toml:"over"`
			Play    int32 `toml:"play"`
			Roll    int32 `toml:"roll"`
			Start   int32 `toml:"start"`
		} `toml:"duration"`
	} `toml:"mahjonger"`
	Sanguoxiao struct {
		Col   int32 `toml:"col"`
		Free  int32 `toml:"free"`
		Open  int32 `toml:"open"`
		Over  int32 `toml:"over"`
		Play  int32 `toml:"play"`
		Row   int32 `toml:"row"`
		Start int32 `toml:"start"`
	} `toml:"sanguoxiao"`
	// 百人游戏
	Baccarat struct {
		Duration struct {
			Free  int32 `toml:"free"`
			Open  int32 `toml:"open"`
			Over  int32 `toml:"over"`
			Play  int32 `toml:"play"`
			Start int32 `toml:"start"`
		} `toml:"duration"`
		Inventory int64 `toml:"inventory"`
	} `toml:"baccarat"`
	BrCowcow struct {
		Duration struct {
			Free  int32 `toml:"free"`
			Open  int32 `toml:"open"`
			Over  int32 `toml:"over"`
			Play  int32 `toml:"play"`
			Start int32 `toml:"start"`
		} `toml:"duration"`
		Inventory   int64 `toml:"inventory"`
		BankerScore int64 `toml:"bankerscore"`
		MaxHost     uint8 `toml:"maxhost"`
	} `toml:"brcowcow"`
	//龙虎斗
	TigerXdragon struct {
		Duration struct {
			Free  int32 `toml:"free"`
			Open  int32 `toml:"open"`
			Over  int32 `toml:"over"`
			Play  int32 `toml:"play"`
			Start int32 `toml:"start"`
		} `toml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"TigerXdragon"`
	//百人骰宝
	Brtoubao struct {
		Duration struct {
			Free  int32 `yaml:"free"`
			Start int32 `yaml:"start"`
			Play  int32 `yaml:"play"`
			Open  int32 `yaml:"open"`
			Over  int32 `yaml:"over"`
		} ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Brtoubao"`
	//百人推筒子
	Brtuitongzi struct {
		Duration struct {
			Free  int32 `yaml:"free"`
			Start int32 `yaml:"start"`
			Play  int32 `yaml:"play"`
			Open  int32 `yaml:"open"`
			Over  int32 `yaml:"over"`
		} ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Brtuitongzi"`
}
