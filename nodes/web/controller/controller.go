package controller

import (
	"github.com/golang-jwt/jwt/v5"
	superGin "github.com/po2656233/superplace/components/gin"
	"github.com/po2656233/superplace/const/code"
	exString "github.com/po2656233/superplace/extend/string"
	sgxLogger "github.com/po2656233/superplace/logger"
	"net/http"
	"strconv"
	data2 "superman/internal/conf"
	. "superman/internal/constant"
	. "superman/internal/hints"
	pb "superman/internal/protocol/gofile"
	"superman/internal/redis_cluster"
	"superman/internal/rpc"
	"superman/internal/token"
	"superman/nodes/web/sdk"
	"time"
)

// index h5客户端
func (p *Controller) index(c *superGin.Context) {
	c.HTML200("index.html")
}

// hello 输出json示例
// http://127.0.0.1:8089/hello
func (p *Controller) hello(c *superGin.Context) {
	// 输出json
	RenderResult(c, http.StatusOK, map[string]string{
		"data": "hello",
	})
}

// http://127.0.0.1:8089/class/list
func (p *Controller) classList(c *superGin.Context) {
	data, errCode := rpc.SendDataToDB(p.App, &pb.GetClassListReq{})
	//获取分类列表
	if errCode != code.OK {
		RenderResult(c, errCode, code.CodeTxt[errCode])
		return
	}
	RenderResult(c, http.StatusOK, data)
}

// http://127.0.0.1:8089/room/list
func (p *Controller) roomList(c *superGin.Context) {
	claims, ok := c.Get("claims")
	if !ok || claims == nil {
		sgxLogger.Warnf("roomCreate get claims fail. current params=%v", c.GetParams())
		RenderResult(c, Login15, StatusText[Login15])
		return
	}
	uid := int64(0)
	clms, ok := claims.(*jwt.RegisteredClaims)
	if !ok {
		sgxLogger.Warnf("roomCreate claims fail. current params=%v", c.GetParams())
		RenderResult(c, Service003, StatusText[Service003])
		return
	}
	uid, _ = strconv.ParseInt(clms.ID, 10, 64)
	data, errCode := rpc.SendDataToDB(p.App, &pb.GetRoomListReq{
		Uid: uid,
	})
	//获取房间列表
	if errCode != code.OK {
		RenderResult(c, errCode, code.CodeTxt[errCode])
		return
	}
	RenderResult(c, http.StatusOK, data)
}

// http://127.0.0.1:8089/table/list
func (p *Controller) tableList(c *superGin.Context) {
	rid := c.GetInt64("rid", 0)
	if rid < 1 {
		sgxLogger.Warnf("rid < 1 params=%v", c.GetParams())
		RenderResult(c, Service003, StatusText[Service003])
		return
	}
	data, errCode := rpc.SendDataToDB(p.App, &pb.GetTableListReq{
		Rid: rid,
	})
	//获取游戏列表
	if errCode != code.OK {
		RenderResult(c, errCode, code.CodeTxt[errCode])
		return
	}
	RenderResult(c, http.StatusOK, data)
}

// http://127.0.0.1:8089/game/list
func (p *Controller) gameList(c *superGin.Context) {
	kid := c.GetInt64("kid", -1)
	if kid == 0 {
		sgxLogger.Warnf("rid < 1 params=%v", c.GetParams())
		RenderResult(c, Service003, StatusText[Service003])
		return
	}
	data, errCode := rpc.SendDataToDB(p.App, &pb.GetGameListReq{
		Kid: kid,
	})
	//获取游戏列表
	if errCode != code.OK {
		RenderResult(c, errCode, code.CodeTxt[errCode])
		return
	}
	RenderResult(c, http.StatusOK, data)
}

//////////////////////////////////////////////////////////////////

// http://127.0.0.1:8089/room/create
func (p *Controller) roomCreate(c *superGin.Context) {
	claims, ok := c.Get("claims")
	if !ok || claims == nil {
		sgxLogger.Warnf("roomCreate get claims fail. current params=%v", c.GetParams())
		RenderResult(c, Login15, StatusText[Login15])
		return
	}
	uid := int64(0)
	if clms, ok := claims.(*jwt.RegisteredClaims); ok {
		uid, _ = strconv.ParseInt(clms.ID, 10, 64)
	}
	name, ok1 := c.GetPostForm("name")
	roomKey, ok2 := c.GetPostForm("roomkey")
	enterScore, ok3 := c.GetPostForm("enterscore")
	level, ok5 := c.GetPostForm("level")
	if !ok1 || !ok2 || !ok3 || !ok5 {
		sgxLogger.Warnf("roomCreate get hostid fail. current params=%v", c.GetParams())
		RenderResult(c, Service004, StatusText[Service004])
		return
	}
	enScore, _ := strconv.ParseInt(enterScore, 10, 64)
	//mCount, _ := strconv.ParseInt(maxCount, 10, 64)
	lev, _ := strconv.ParseInt(level, 10, 64)
	data, errCode := rpc.SendDataToDB(p.App, &pb.CreateRoomReq{
		HostId:     uid,
		Name:       name,
		RoomKey:    roomKey,
		EnterScore: enScore,
		Level:      pb.RoomLevel(lev),
	})

	// 创建房间结果
	if errCode != code.OK {
		RenderResult(c, errCode, code.CodeTxt[errCode])
		return
	}
	if resp, ok := data.(*pb.CreateRoomResp); ok {
		if resp.Info == nil || resp.Info.Id == 0 {
			RenderResult(c, http.StatusNotAcceptable, "房间创建失败")
			return
		}
	}
	RenderResult(c, http.StatusOK, data)
}

// http://127.0.0.1:8089/table/create
func (p *Controller) tableCreate(c *superGin.Context) {
	claims, ok := c.Get("claims")
	if !ok || claims == nil {
		sgxLogger.Warnf("roomCreate get claims fail. current params=%v", c.GetParams())
		RenderResult(c, Login15, StatusText[Login15])
		return
	}
	//uid := int64(0)
	//if clms, ok := claims.(*jwt.RegisteredClaims); ok {
	//	uid, _ = strconv.ParseInt(clms.ID, 10, 64)
	//}
	strRid, ok1 := c.GetPostForm("rid")
	strGid, ok2 := c.GetPostForm("gid")
	name, ok2 := c.GetPostForm("name")
	strPlayscore, ok3 := c.GetPostForm("playscore")
	strOpentime, ok4 := c.GetPostForm("opentime")
	strTaxation, ok5 := c.GetPostForm("taxation")
	strCommission, ok6 := c.GetPostForm("commission")
	strAmount, ok7 := c.GetPostForm("amount")
	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 || !ok6 || !ok7 {
		sgxLogger.Warnf("roomCreate get hostid fail. current params=%v", c.GetParams())
		RenderResult(c, Service004, StatusText[Service004])
		return
	}
	rid, _ := strconv.ParseInt(strRid, 10, 64)
	gid, _ := strconv.ParseInt(strGid, 10, 64)
	playscore, _ := strconv.ParseInt(strPlayscore, 10, 64)
	opentime, _ := strconv.ParseInt(strOpentime, 10, 64)
	taxation, _ := strconv.ParseInt(strTaxation, 10, 64)
	commission, _ := strconv.ParseInt(strCommission, 10, 64)
	amount, _ := strconv.ParseInt(strAmount, 10, 64)
	data, errCode := rpc.SendDataToDB(p.App, &pb.CreateTableReq{
		Rid:        rid,
		Gid:        gid,
		Playscore:  playscore,
		Name:       name,
		Opentime:   opentime,
		Taxation:   taxation,
		Commission: int32(commission),
		Amount:     int32(amount),
	})

	// 创建房间结果
	if errCode != code.OK {
		RenderResult(c, errCode, code.CodeTxt[errCode])
		return
	}
	if resp, ok := data.(*pb.CreateTableResp); ok {
		if resp.Table == nil || resp.Table.Id == 0 {
			RenderResult(c, http.StatusNotAcceptable, "牌桌创建失败")
			return
		}
	}
	RenderResult(c, http.StatusOK, data)
}

// http://127.0.0.1:8089/table/delete
func (p *Controller) tableDelete(c *superGin.Context) {
	claims, ok := c.Get("claims")
	if !ok || claims == nil {
		sgxLogger.Warnf("roomCreate get claims fail. current params=%v", c.GetParams())
		RenderResult(c, Login15, StatusText[Login15])
		return
	}
	uid := int64(0)
	if clms, ok := claims.(*jwt.RegisteredClaims); ok {
		uid, _ = strconv.ParseInt(clms.ID, 10, 64)
	}
	strTid, ok1 := c.GetPostForm("tid")
	if !ok1 {
		sgxLogger.Warnf("roomCreate get hostid fail. current params=%v", c.GetParams())
		RenderResult(c, Service004, StatusText[Service004])
		return
	}
	// 先检测游戏服相应游戏是否处于关闭状态，如果没有关闭则请求关闭
	tid, _ := strconv.ParseInt(strTid, 10, 64)
	data, errCode := rpc.SendDataToDB(p.App, &pb.DeleteTableReq{
		Tid:    tid,
		HostId: uid,
	})

	// 创建房间结果
	if errCode != code.OK {
		RenderResult(c, errCode, code.CodeTxt[errCode])
		return
	}
	if resp, ok := data.(*pb.DeleteTableResp); ok {
		if resp.Tid == 0 || resp.Rid == 0 {
			RenderResult(c, http.StatusNotAcceptable, "牌桌删除失败")
			return
		}
	}
	RenderResult(c, http.StatusOK, data)
}

// ///////////////////////////////////////////////////////////////////////////
// login 根据pid获取sdkConfig，与第三方进行帐号登陆效验
// http://127.0.0.1:8089/login?pid=2126001&account=test1&password=test1
func (p *Controller) login(c *superGin.Context) {
	pid := c.PostInt32(ProcessId, 0)
	if pid < 1 {
		sgxLogger.Warnf("pid < 1 params=%s", c.GetParams())
		RenderResult(c, Service002)
		return
	}

	config := data2.SdkConfig.Get(pid)
	if config == nil {
		sgxLogger.Warnf("if platformConfig == nil {. params=%s", c.GetParams())
		RenderResult(c, Login02, StatusText[Login02])
		return
	}
	sdkInvoke, err := sdk.GetInvoke(config.SdkId)
	if err != nil {
		sgxLogger.Warnf("[pid = %d] get invoke error. params=%s", pid, c.GetParams())
		RenderResult(c, Service001, StatusText[Service001])
		return
	}
	params := c.GetParams(true)
	params[ProcessId] = exString.ToString(pid)
	sdkInvoke.Login(config, params, func(statusCode int32, result sdk.Params, error ...error) {
		// 回调校验账号密码
		if code.IsFail(statusCode) {
			sgxLogger.Warnf("login validate fail. code = %d, params = %s", statusCode, c.GetParams())
			if len(error) > 0 {
				sgxLogger.Warnf("code = %d, error = %s", statusCode, error[0])
			}

			RenderResult(c, statusCode, StatusText[int(statusCode)])
			return
		}

		if result == nil {
			sgxLogger.Warnf("callback result map is nil. params= %s", c.GetParams())
			RenderResult(c, Login02, StatusText[Login02])
			return
		}
		strUid, _ := result.GetString("open_id")
		uid, err := strconv.ParseInt(strUid, 10, 64)
		if err != nil || uid == 0 {
			sgxLogger.Warnf("callback result map not found `open_id`. result = %s", result)
			RenderResult(c, Login02, StatusText[Login02])
			return
		}

		// 第一种 获取token
		// 从redis上获取token,如果获取到token,且有效时长超过5分钟的，则返回原token.
		userKey := GetTokenKey(params[Username])
		ret := redis_cluster.SingleRedis().DB.Get(c.Context, userKey)
		if ret.Err() == nil && ret.String() != "" { // 返回现有token
			strToken := ret.Val()
			//校验token时长
			retCode, _ := token.VerifyToken(c, strToken)
			if retCode == http.StatusOK {
				RenderResult(c, http.StatusOK, strToken)
				return
			}
		}

		timeout := time.Minute * token.TOKEN_MAX_REMAINING_MINUTE
		strToken, tkErr := token.CreateToken(uid, params[Username], timeout)
		if tkErr != nil {
			RenderResult(c, http.StatusUnauthorized)
			return
		}
		// 返回Bearer Token
		redis_cluster.SingleRedis().DB.Set(c.Context, userKey, strToken, timeout)
		RenderResult(c, http.StatusOK, strToken)
		// 另一种token
		//openid, found := result.GetString("open_id")
		//if found == false {
		//	sgxLogger.Warnf("callback result map not found `open_id`. result = %s", result)
		//	RenderResult(c, Login02, StatusText[Login02])
		//	return
		//}
		//
		//base64Token := token.New(pid, openid, config.Salt).ToBase64()
		//RenderResult(c, http.StatusOK, base64Token)
	})
}

// register 开发模式帐号注册
// http://127.0.0.1:8089/register?account=test11&password=test11
func (p *Controller) register(c *superGin.Context) {
	accountName := c.PostString(Username, "")
	password := c.PostString(Password, "")
	data, statusCode := rpc.SendDataToAcc(p.App, &pb.RegisterReq{
		Name:     accountName,
		Password: password,
		Address:  c.ClientIP(),
	})
	if statusCode == code.OK {
		resp, ok := data.(*pb.RegisterResp)
		if ok {
			plist := data2.SdkConfig.GetPidList()
			pid := int32(0)
			size := len(plist)
			if 0 < size {
				pid = plist[0]
				//index := rand.Int() % len(plist)
				//pid = plist[index]
			}
			config := data2.SdkConfig.Get(pid)
			resp.SdkId = config.SdkId
			resp.Pid = pid
			resp.Ip = c.ClientIP()
		}
		RenderResult(c, http.StatusOK, resp)
		return
	}
	RenderResult(c, statusCode, StatusText[int(statusCode)])
}

// severList 区服列表
// http://127.0.0.1:8089/server/list/2126001
func (p *Controller) serverList(c *superGin.Context) {
	pid := c.GetInt32("pid", 2126001)

	if pid < 1 {
		sgxLogger.Warnf("if pid < 1 {. params=%v", c.GetParams())
		RenderResult(c, Service002, StatusText[Service002])
		return
	}

	areaGroup, found := data2.AreaGroupConfig.Get(pid)
	if found == false {
		RenderResult(c, Service002, StatusText[Service002])
		return
	}

	dataList := &struct {
		Areas   []*data2.AreaRow       `json:"areas"`
		Servers []*data2.AreaServerRow `json:"servers"`
	}{}

	for _, areaId := range areaGroup.AreaIdList {
		areaRow, found := data2.AreaConfig.Get(areaId)
		if found == false {
			continue
		}
		dataList.Areas = append(dataList.Areas, areaRow)

		serverList := data2.AreaServerConfig.ListWithAreaId(areaRow.AreaId)
		if len(serverList) > 0 {
			dataList.Servers = append(dataList.Servers, serverList...)
		}
	}

	RenderResult(c, http.StatusOK, dataList)
}

// severList 区服列表
// http://127.0.0.1:8089/pid/list
func (p *Controller) pidList(c *superGin.Context) {
	pidList := &struct {
		Pids []int32 `json:"pids"`
	}{}
	pidList.Pids = data2.SdkConfig.GetPidList()
	RenderResult(c, http.StatusOK, pidList)

}
