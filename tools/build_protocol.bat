@echo off
chcp 65001 >nul
@echo "-----------fix package name(本地化)------------------"
@echo "-----------校正 proto文件------------------"
py  .\rectify.py
timeout 1

echo build proto to "go and js" complete!

@echo off

rem node onekey.js
@echo "-----------实例化 各个REQ消息体 的处理函数------------------"
py .\realize_go.py
rem timeout /t 2
echo 开始 3s 倒计时...
choice /c  abcdQ /n /t 3 /d a /m "如需暂停,请按Q键终止,否则退出"
if %errorlevel%==5 goto stopTerminal
goto :eof
:stopTerminal
pause


