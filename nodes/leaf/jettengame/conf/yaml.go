package conf

import (
	"github.com/po2656233/goleaf/log"
	"gopkg.in/yaml.v2"
	"os"
)

type Yaml struct {
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
		Duration struct {
			FreeTime  int32 `yaml:"freeTime"`
			StartTime int32 `yaml:"startTime"`
			BetTime   int32 `yaml:"betTime"`
			OpenTime  int32 `yaml:"openTime"`
			OverTime  int32 `yaml:"overTime"`
		} ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Baccarat"`
	//百人牛牛
	Brcowcow struct {
		Duration struct {
			FreeTime  int32 `yaml:"freeTime"`
			StartTime int32 `yaml:"startTime"`
			BetTime   int32 `yaml:"betTime"`
			OpenTime  int32 `yaml:"openTime"`
			OverTime  int32 `yaml:"overTime"`
		} ` yaml:"duration"`
		Inventory   int64 `yaml:"inventory"`
		BankerScore int64 `yaml:"bankerScore"`
		MaxHost     uint8 `yaml:"maxHost"`
	} ` yaml:"Brcowcow"`
	//百人骰宝
	Brtoubao struct {
		Duration struct {
			FreeTime  int32 `yaml:"freeTime"`
			StartTime int32 `yaml:"startTime"`
			BetTime   int32 `yaml:"betTime"`
			OpenTime  int32 `yaml:"openTime"`
			OverTime  int32 `yaml:"overTime"`
		} ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Brtoubao"`
	//百人推筒子
	Brtuitongzi struct {
		Duration struct {
			FreeTime  int32 `yaml:"freeTime"`
			StartTime int32 `yaml:"startTime"`
			BetTime   int32 `yaml:"betTime"`
			OpenTime  int32 `yaml:"openTime"`
			OverTime  int32 `yaml:"overTime"`
		} ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"Brtuitongzi"`
	//推筒子上庄喝水
	BrtuitongziSZHS struct {
		Duration struct {
			FreeTime  int32 `yaml:"freeTime"`
			StartTime int32 `yaml:"startTime"`
			BetTime   int32 `yaml:"betTime"`
			OpenTime  int32 `yaml:"openTime"`
			OverTime  int32 `yaml:"overTime"`
		} ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"BrtuitongziSZHS"`
	//龙虎斗
	TigerXdragon struct {
		Duration struct {
			FreeTime  int32 `yaml:"freeTime"`
			StartTime int32 `yaml:"startTime"`
			BetTime   int32 `yaml:"betTime"`
			OpenTime  int32 `yaml:"openTime"`
			OverTime  int32 `yaml:"overTime"`
		} ` yaml:"duration"`
		Inventory int64 `yaml:"inventory"`
	} ` yaml:"TigerXdragon"`

	//---------对战类游戏基本配置--------------------
	//斗地主
	Landlord struct {
		Duration struct {
			FreeTime   int32 `yaml:"free"`
			ReadyTime  int32 `yaml:"ready"`
			DealTime   int32 `yaml:"deal"`
			CallTime   int32 `yaml:"call"`
			DoubleTime int32 `yaml:"double"`
			PlayTime   int32 `yaml:"play"`
			OpenTime   int32 `yaml:"open"`
			OverTime   int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"Landlord"`
	//跑得快
	PaoDeKuai struct {
		Duration struct {
			FreeTime   int32 `yaml:"free"`
			ReadyTime  int32 `yaml:"ready"`
			DealTime   int32 `yaml:"deal"`
			CallTime   int32 `yaml:"call"`
			DoubleTime int32 `yaml:"double"`
			PlayTime   int32 `yaml:"play"`
			OpenTime   int32 `yaml:"open"`
			OverTime   int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"PaoDeKuai"`
	//扎金花
	Zhajinhua struct {
		Duration struct {
			FreeTime    int32 `yaml:"free"`
			NoReadyTime int32 `yaml:"ready"`
			DealTime    int32 `yaml:"deal"`
			CompTime    int32 `yaml:"compare"`
			PlayTime    int32 `yaml:"play"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson  int32 `yaml:"minPerson"`
		MaxPerson  int32 `yaml:"maxPerson"`
		LimitScore int64 `yaml:"limitScore"`
	} ` yaml:"Zhajinhua"`
	//通比牛牛
	Tbcowcow struct {
		Duration struct {
			FreeTime int32 `yaml:"free"`
			CallTime int32 `yaml:"call"`
			DealTime int32 `yaml:"deal"`
			CompTime int32 `yaml:"compare"`
			PlayTime int32 `yaml:"play"`
			OverTime int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"Tbcowcow"`
	//抢庄牛牛
	Qzcowcow struct {
		Duration struct {
			FreeTime int32 `yaml:"free"`
			CallTime int32 `yaml:"call"`
			DealTime int32 `yaml:"deal"`
			CompTime int32 `yaml:"compare"`
			PlayTime int32 `yaml:"play"`
			OverTime int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"Qzcowcow"`
	//三公
	Sangong struct {
		Duration struct {
			FreeTime int32 `yaml:"free"`
			CallTime int32 `yaml:"call"`
			DealTime int32 `yaml:"deal"`
			CompTime int32 `yaml:"compare"`
			PlayTime int32 `yaml:"play"`
			OverTime int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"Sangong"`
	//极速扎金花
	ZhajinhuaJiSu struct {
		Duration struct {
			FreeTime int32 `yaml:"free"`
			CallTime int32 `yaml:"call"`
			DealTime int32 `yaml:"deal"`
			CompTime int32 `yaml:"compare"`
			PlayTime int32 `yaml:"play"`
			OverTime int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"ZhajinhuaJiSu"`
	//推筒子
	Tuitongzi struct {
		Duration struct {
			FreeTime int32 `yaml:"free"`
			CallTime int32 `yaml:"call"`
			DealTime int32 `yaml:"deal"`
			CompTime int32 `yaml:"compare"`
			PlayTime int32 `yaml:"play"`
			OverTime int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"Tuitongzi"`

	//所有麻将配置
	//国标麻将
	Mahjong struct {
		Duration struct {
			DirectTime  int32 `yaml:"direct"`
			DecideTime  int32 `yaml:"decide"`
			RollTime    int32 `yaml:"roll"`
			DealTime    int32 `yaml:"deal"`
			PlayTime    int32 `yaml:"play"`
			OperateTime int32 `yaml:"operate"`
			OpenTime    int32 `yaml:"open"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"Mahjong"`
	//二人麻将
	MahjongER struct {
		Duration struct {
			DirectTime  int32 `yaml:"direct"`
			DecideTime  int32 `yaml:"decide"`
			RollTime    int32 `yaml:"roll"`
			DealTime    int32 `yaml:"deal"`
			PlayTime    int32 `yaml:"play"`
			OperateTime int32 `yaml:"operate"`
			OpenTime    int32 `yaml:"open"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongER"`
	//长沙麻将
	MahjongCS struct {
		Duration struct {
			DirectTime  int32 `yaml:"direct"`
			DecideTime  int32 `yaml:"decide"`
			RollTime    int32 `yaml:"roll"`
			DealTime    int32 `yaml:"deal"`
			PlayTime    int32 `yaml:"play"`
			OperateTime int32 `yaml:"operate"`
			OpenTime    int32 `yaml:"open"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongCS"`
	//广东麻将
	MahjongGD struct {
		Duration struct {
			DirectTime  int32 `yaml:"direct"`
			DecideTime  int32 `yaml:"decide"`
			RollTime    int32 `yaml:"roll"`
			DealTime    int32 `yaml:"deal"`
			PlayTime    int32 `yaml:"play"`
			OperateTime int32 `yaml:"operate"`
			OpenTime    int32 `yaml:"open"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongGD"`
	//四川麻将
	MahjongSC struct {
		Duration struct {
			DirectTime  int32 `yaml:"direct"`
			DecideTime  int32 `yaml:"decide"`
			RollTime    int32 `yaml:"roll"`
			DealTime    int32 `yaml:"deal"`
			PlayTime    int32 `yaml:"play"`
			OperateTime int32 `yaml:"operate"`
			OpenTime    int32 `yaml:"open"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongSC"`
	//血流成河
	MahjongXLCH struct {
		Duration struct {
			DecideTime  int32 `yaml:"decide"`
			RollTime    int32 `yaml:"roll"`
			DealTime    int32 `yaml:"deal"`
			ChangeTime  int32 `yaml:"change"`
			DingQueTime int32 `yaml:"dingque"`
			PlayTime    int32 `yaml:"play"`
			OperateTime int32 `yaml:"operate"`
			OpenTime    int32 `yaml:"open"`
			TingTime    int32 `yaml:"ting"`
			HuaZhuTime  int32 `yaml:"huazhu"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongXLCH"`
	//血战到底
	MahjongXZDD struct {
		Duration struct {
			DecideTime  int32 `yaml:"decide"`
			RollTime    int32 `yaml:"roll"`
			DealTime    int32 `yaml:"deal"`
			ChangeTime  int32 `yaml:"change"`
			DingQueTime int32 `yaml:"dingque"`
			PlayTime    int32 `yaml:"play"`
			OperateTime int32 `yaml:"operate"`
			OpenTime    int32 `yaml:"open"`
			TingTime    int32 `yaml:"ting"`
			HuaZhuTime  int32 `yaml:"huazhu"`
			OverTime    int32 `yaml:"over"`
		} ` yaml:"duration"`
		MinPerson int32 `yaml:"minPerson"`
		MaxPerson int32 `yaml:"maxPerson"`
	} ` yaml:"MahjongXZDD"`
	// 三国消
	SanGuoXiao struct {
		Duration struct {
			StartTime int32 `yaml:"start"`
			PlayTime  int32 `yaml:"play"`
			OpenTime  int32 `yaml:"open"`
			OverTime  int32 `yaml:"over"`
		} ` yaml:"duration"`
		Col int32 `yaml:"col"`
		Row int32 `yaml:"row"`
	} ` yaml:"SanGuoXiao"`
}

var YamlObj = new(Yaml)

func InitYml() {
	//读取配置信息
	yamlFile, err := os.ReadFile(GameYamlPath)
	if err != nil {
		log.Fatal("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, YamlObj)
	if err != nil {
		log.Error("Unmarshal: %v", err)

	}

	log.Debug("conf chips info:%v ok", YamlObj.Chips)
}
func updateYamlConf(data string) (err error) {
	temp := YamlObj
	err = yaml.Unmarshal([]byte(data), YamlObj)
	if err != nil {
		log.Error("updateYamlConf data:[%v] err:%v", data, err)
		return
	}
	YamlObj = temp
	return
}
