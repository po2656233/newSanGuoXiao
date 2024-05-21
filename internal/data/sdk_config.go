package data

import (
	cstring "github.com/po2656233/superplace/extend/string"
	sgxLogger "github.com/po2656233/superplace/logger"
	sgxError "github.com/po2656233/superplace/logger/error"
	"sort"
)

type (
	SdkRow struct {
		SdkId        int32             `json:"sdkId"`        // sdk id
		CallbackName string            `json:"callbackName"` // 支付回调名称路由使用
		Salt         string            `json:"-"`            // !禁止JSON输出
		Params       map[string]string `json:"-"`            // !禁止JSON输出
		PIDList      []int32           `json:"pidList"`      // sdk包id列表(一个sdk可以输出多个安装包)
		Desc         string            `json:"desc"`         // 描述
	}

	// sdk平台参数
	sdkConfig struct {
		maps map[int32]*SdkRow // key:pid, value:PlatformRow
	}
)

func (p *sdkConfig) Name() string {
	return "sdkConfig"
}

func (p *sdkConfig) Init() {
}

func (p *sdkConfig) OnLoad(maps interface{}, _ bool) (int, error) {
	list, ok := maps.([]interface{})
	if !ok {
		return 0, sgxError.Error("maps convert to []interface{} error.")
	}

	loadMaps := make(map[int32]*SdkRow)

	for index, data := range list {
		loadConfig := &SdkRow{}
		err := DecodeData(data, loadConfig)
		if err != nil {
			sgxLogger.Warnf("decode error. [row = %d, %v], err = %s", index+1, loadConfig, err)
			continue
		}

		for _, pid := range loadConfig.PIDList {
			loadMaps[pid] = loadConfig
		}
	}

	p.maps = loadMaps

	return len(list), nil
}

func (p *sdkConfig) OnAfterLoad(_ bool) {
}

func (p *sdkConfig) Get(pid int32) *SdkRow {
	platformRow, found := p.maps[pid]
	if found {
		return platformRow
	}

	return nil
}

func (p *sdkConfig) GetWithName(callName string) *SdkRow {
	if callName == "" {
		return nil
	}

	for _, row := range p.maps {
		if row.CallbackName == callName {
			return row
		}
	}

	return nil
}
func (p *sdkConfig) GetPidList() []int32 {
	pids := make([]int32, 0)
	for pid, _ := range p.maps {
		pids = append(pids, pid)
	}
	sort.Slice(pids, func(i, j int) bool {
		return pids[i] < pids[j]
	})
	return pids
}

func (p *SdkRow) AppId() string {
	return p.Params["appId"]
}

func (p *SdkRow) AppKey() string {
	return p.Params["appKey"]
}

func (p *SdkRow) LoginURL() string {
	return p.Params["loginUrl"]
}

func (p *SdkRow) GetString(key string) string {
	v, found := p.Params[key]
	if !found {
		return ""
	}
	return v
}

func (p *SdkRow) GetInt(key string) int {
	v, found := p.Params[key]
	if !found {
		return 0
	}

	intValue, _ := cstring.ToInt(v, 0)

	return intValue
}
