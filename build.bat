
go build -tags=master -o superman.exe -gcflags "all=-N -l" main.go

@echo off
setlocal enabledelayedexpansion

:: 定义软件的路径和启动参数
set "softwarePath=superman.exe"
set "master=master --path=./config/profile-gc.json --node=gc-master"
set "center=center --path=./config/profile-gc.json --node=gc-center"
set "gate=gate --path=./config/profile-gc.json --node=gc-gate-1"
set "web=web --path=./config/profile-gc.json --node=gc-web-1"
set "game=game --path=./config/profile-gc.json --node=10001"
set "leaf=leaf --path=./config/profile-gc.json --node=20001"

:: 启动第一个软件
start "" %softwarePath% %master%

:: 根据不同的启动参数启动第二个软件
start "" %softwarePath% %center%

:: 根据不同的启动参数启动第三个软件
start "" %softwarePath% %gate%

:: 根据不同的启动参数启动第四个软件
start "" %softwarePath% %web%

:: 根据不同的启动参数启动第五个软件
start "" %softwarePath% %game%

:: 根据不同的启动参数启动第六个软件
start "" %softwarePath% %leaf%