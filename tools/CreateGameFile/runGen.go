package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

// GameST 游戏结构
type GameST struct {
	Name string
	ID   int
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("获取当前执行文件绝对路径 出错:%v", err)
		return ""
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
func CreateGame(gameName, fileName, data string) {
	tmpl, err := template.New(gameName).Parse(data)
	if err != nil {
		fmt.Printf("err: %v!", err)
		return
	}
	if fileName == "" {
		fileName = gameName
	}
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Printf("err1: %v!", err1)
		return
	}

	gst := GameST{
		Name: gameName,
	}
	tmpl.Execute(f, gst)
}

func main() {
	//fmt.Println("getTmpDir（当前系统临时目录） = ", getTmpDir())
	//fmt.Println("getCurrentAbPathByExecutable（仅支持go build） = ", getCurrentAbPathByExecutable())
	//fmt.Println("getCurrentAbPathByCaller（仅支持go run） = ", getCurrentAbPathByCaller())
	//fmt.Println("getCurrentAbPath（最终方案-全兼容） = ", getCurrentAbPath())
	fileName := filepath.Join(getCurrentAbPath(), "game.tmpl")
	f, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("ReadFile 文件路径:%s err:%v", fileName, err)
		return
	}
	// 创建所需的游戏
	gameName := "Sanguoxiao"
	lowName := strings.ToLower(gameName)
	if err = os.Mkdir(filepath.Join(getCurrentAbPath(), lowName), os.ModePerm); err != nil {
		fmt.Printf("Mkdir 文件路径:%s err:%v", lowName, err)
		return
	}

	gameFile := filepath.Join(getCurrentAbPath(), lowName, lowName+".go")
	CreateGame(gameName, gameFile, string(f))

}
