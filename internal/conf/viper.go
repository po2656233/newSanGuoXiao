package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type ConfigData map[string]any

var data = make(map[string]ConfigData)

func InitViper(filePath string) {
	viperCfg := viper.New()

	strDir, fileName := filepath.Split(filePath)
	if fileName == "" {
		panic(fmt.Errorf("no include file:%v", filePath))
	}
	if strDir == "" {
		strDir = "."
	}
	viperCfg.AddConfigPath(strDir)

	suffix := filepath.Ext(fileName)
	confName := strings.TrimSuffix(fileName, suffix)
	viperCfg.SetConfigName(confName)
	if 1 < len(suffix) {
		viperCfg.SetConfigType(suffix[1:])
	}

	err := viperCfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var confData ConfigData
	err = viperCfg.Unmarshal(&confData)
	if err != nil {
		panic(err)
	}
	viperCfg.WatchConfig()
	viperCfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("OnConfigChange file changed:", e.Name)
		if err = viperCfg.Unmarshal(&confData); err != nil {
			fmt.Println("OnConfigChange err:", err)
		}
		data[e.Name] = confData
	})
	data[filePath] = confData
}

func GetViperConfig(filePath string) ConfigData {
	return data[filePath]
}
