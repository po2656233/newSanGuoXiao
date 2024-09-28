# 构建约束

以master举例
//go:build master

go build -tags=master -o superman -gcflags "all=-N -l" main.go 

./superman master --path=./config/profile-gc.json --node=gc-master



tools目录存放的是 
辅助项目的工具
// CreateGameFile目录下实现的是，根据游戏模板，自动生成实例化文件。
// genSqlmodel实现的是 将数据库的结构数据生成到./internal/sql_model下。

build_protocol.bat脚本
是调用realize.py和rectify.py将协议文件生成到./internal/protocol/gofile目录下
 并注册协议到handle_msg.go文件中





