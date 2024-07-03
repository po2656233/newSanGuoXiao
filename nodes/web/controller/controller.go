package controller

import (
	sgxGin "github.com/po2656233/superplace/components/gin"
	"github.com/po2656233/superplace/const/code"
	sgxString "github.com/po2656233/superplace/extend/string"
	sgxLogger "github.com/po2656233/superplace/logger"
	"net/http"
	data2 "superman/internal/conf"
	. "superman/internal/hints"
	pb "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	"superman/internal/token"
	"superman/nodes/web/sdk"
)

type Controller struct {
	sgxGin.BaseController
}

func (p *Controller) Init() {
	group := p.Group("/")
	group.GET("/", p.index)
	group.GET("/hello", p.hello)
	group.GET("/register", p.register)
	group.GET("/login", p.login)
	group.GET("/pid/list", p.pidList)
	group.GET("/server/list/:pid", p.serverList)
}

// index h5客户端
func (p *Controller) index(c *sgxGin.Context) {
	c.HTML200("index.html")
}

// hello 输出json示例
// http://127.0.0.1:8089/hello
func (p *Controller) hello(c *sgxGin.Context) {
	// 输出json
	RenderResult(c, http.StatusOK, map[string]string{
		"data": "hello",
	})
}

// register 开发模式帐号注册
// http://127.0.0.1:8089/register?account=test11&password=test11
func (p *Controller) register(c *sgxGin.Context) {
	accountName := c.GetString("account", "", true)
	password := c.GetString("password", "", true)
	data, statusCode := rpc.SendData(p.App, rpc.SourcePath, rpc.AccountActor, rpc.CenterType, &pb.RegisterReq{
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

// login 根据pid获取sdkConfig，与第三方进行帐号登陆效验
// http://127.0.0.1:8089/login?pid=2126001&account=test1&password=test1
func (p *Controller) login(c *sgxGin.Context) {
	pid := c.GetInt32("pid", 0, true)

	if pid < 1 {
		sgxLogger.Warnf("if pid < 1 {. params=%s", c.GetParams())
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
	params["pid"] = sgxString.ToString(pid)

	// invoke login.
	sdkInvoke.Login(config, params, func(statusCode int32, result sdk.Params, error ...error) {
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

		openId, found := result.GetString("open_id")
		if found == false {
			sgxLogger.Warnf("callback result map not found `open_id`. result = %s", result)
			RenderResult(c, Login02, StatusText[Login02])
			return
		}

		base64Token := token.New(pid, openId, config.Salt).ToBase64()
		RenderResult(c, http.StatusOK, base64Token)
	})
}

// severList 区服列表
// http://127.0.0.1:8089/server/list/2126001
func (p *Controller) serverList(c *sgxGin.Context) {
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
func (p *Controller) pidList(c *sgxGin.Context) {
	pidList := &struct {
		Pids []int32 `json:"pids"`
	}{}
	pidList.Pids = data2.SdkConfig.GetPidList()
	RenderResult(c, http.StatusOK, pidList)

}
