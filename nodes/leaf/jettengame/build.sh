#记得下载GCC进行编译 麻将库放置gamedata/goclib/algorithm
export   GOOS=linux
export   GOARCH=amd64
#export  CGO_ENABLED=0
#go build  -o qpserver main.go
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o go-qpserver-linux main.go

# set   GOOS=linux
# set   GOARCH=amd64
# set  CGO_ENABLED=0
# go build  -o qpserver main.go

# 找到goleaf的ws_conn.go文件  修改ReadMsg
func (wsConn *WSConn) ReadMsg() ([]byte, error) {
	wsConn.conn.SetReadDeadline(time.Now().Add(35 * time.Second))// 加上
	_, b, err := wsConn.conn.ReadMessage()
	wsConn.conn.SetReadDeadline(time.Time{})// 加上
	return b, err
}
# newWSConn函数内
      conn.SetWriteDeadline(time.Now().Add(60 * time.Second))// 加上
			err := conn.WriteMessage(websocket.BinaryMessage, b)
			if err != nil {
				break
			}
			conn.SetWriteDeadline(time.Time{})// 加上
