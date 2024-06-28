# 构建约束

以master举例
//go:build master
go build -tags=master -o superman -gcflags "all=-N -l" main.go 
./superman master --path=./config/profile-gc.json --node=gc-master

