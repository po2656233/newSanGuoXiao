package conf

import (
	"github.com/po2656233/superplace/logger"
	superError "github.com/po2656233/superplace/logger/error"
)

type (
	AreaRow struct {
		AreaId   int32  `json:"areaId"`   // 游戏区id
		AreaName string `json:"areaName"` // 游戏区名称
		Gate     string `json:"gate"`     // 游戏区对应的网关地址
		GateTcp  string `json:"gate_tcp"` // 游戏区对应的网关地址
		GateKcp  string `json:"gate_kcp"` // 游戏区对应的网关地址
	}

	// 游戏区
	areaConfig struct {
		maps map[int32]*AreaRow
	}
)

// Name 根据名称读取 ./config/data/areaConfig.json文件
func (p *areaConfig) Name() string {
	return "areaConfig"
}

func (p *areaConfig) Init() {
	p.maps = make(map[int32]*AreaRow)
}

func (p *areaConfig) OnLoad(maps interface{}, _ bool) (int, error) {
	list, ok := maps.([]interface{})
	if !ok {
		return 0, superError.Error("maps convert to []interface{} error.")
	}

	loadMaps := make(map[int32]*AreaRow)
	for index, data := range list {
		loadConfig := &AreaRow{}
		err := DecodeData(data, loadConfig)
		if err != nil {
			logger.Warnf("decode error. [row = %d, %v], err = %s", index+1, loadConfig, err)
			continue
		}

		loadMaps[loadConfig.AreaId] = loadConfig
	}

	p.maps = loadMaps

	return len(list), nil
}

func (p *areaConfig) OnAfterLoad(_ bool) {
}

func (p *areaConfig) Get(pk int32) (*AreaRow, bool) {
	i, found := p.maps[pk]
	return i, found
}

func (p *areaConfig) Contain(pk int32) bool {
	_, found := p.Get(pk)
	return found
}
