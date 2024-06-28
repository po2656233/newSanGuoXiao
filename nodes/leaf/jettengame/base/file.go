package base

import (
	"encoding/json"
	"fmt"
	exFile "github.com/po2656233/superplace/extend/file"
	"os"
)

func ToJsonFile(fileName string, data interface{}, prefix, indent string) {
	jsonData, err := json.MarshalIndent(data, prefix, indent)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}
	curPath := exFile.GetCurrentPath()
	fileName = curPath + "/../" + fileName
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("data 数据已成功写入文件", fileName)
}
