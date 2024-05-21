package data

import (
	dataConf "github.com/po2656233/superplace/components/data-config"
	mapStructure "github.com/po2656233/superplace/extend/mapstructure"
	"sanguoxiao/internal/types"
)

var (
	AreaConfig       = &areaConfig{}
	AreaGroupConfig  = &areaGroupConfig{}
	AreaServerConfig = &areaServerConfig{}
	SdkConfig        = &sdkConfig{}
	CodeConfig       = &codeConfig{}
	PlayerInitConfig = &playerInitConfig{}
)

func New() *dataConf.Component {
	dataConfig := dataConf.New()
	dataConfig.Register(
		AreaConfig,
		AreaGroupConfig,
		AreaServerConfig,
		SdkConfig,
		CodeConfig,
		PlayerInitConfig,
	)
	return dataConfig
}

func DecodeData(input interface{}, output interface{}) error {
	return mapStructure.HookDecode(
		input,
		output,
		"json",
		types.GetDecodeHooks(),
	)
}
