package controller

import (
	sgxGin "github.com/po2656233/superplace/components/gin"
	"github.com/po2656233/superplace/const/code"
	sgxString "github.com/po2656233/superplace/extend/string"
	sgxLogger "github.com/po2656233/superplace/logger"
	data2 "sanguoxiao/internal/data"
	"sanguoxiao/internal/hints"
	pb "sanguoxiao/internal/protocol/gofile"
	"sanguoxiao/internal/rpc"
	"sanguoxiao/internal/token"
	"sanguoxiao/nodes/web/sdk"
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
// http://127.0.0.1/hello
func (p *Controller) hello(c *sgxGin.Context) {
	// 输出json
	hints.RenderResult(c, code.OK, map[string]string{
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
	hints.RenderResult(c, statusCode, data)
}

// login 根据pid获取sdkConfig，与第三方进行帐号登陆效验
// http://127.0.0.1:8089/login?pid=2126001&account=test1&password=test1
func (p *Controller) login(c *sgxGin.Context) {
	pid := c.GetInt32("pid", 0, true)

	if pid < 1 {
		sgxLogger.Warnf("if pid < 1 {. params=%s", c.GetParams())
		hints.RenderResult(c, hints.Service002)
		return
	}

	config := data2.SdkConfig.Get(pid)
	if config == nil {
		sgxLogger.Warnf("if platformConfig == nil {. params=%s", c.GetParams())
		hints.RenderResult(c, hints.Login02)
		return
	}

	sdkInvoke, err := sdk.GetInvoke(config.SdkId)
	if err != nil {
		sgxLogger.Warnf("[pid = %d] get invoke error. params=%s", pid, c.GetParams())
		hints.RenderResult(c, hints.Service001)
		return
	}

	params := c.GetParams(true)
	params["pid"] = sgxString.ToString(pid)

	// invoke login
	sdkInvoke.Login(config, params, func(statusCode int32, result sdk.Params, error ...error) {
		if code.IsFail(statusCode) {
			sgxLogger.Warnf("login validate fail. code = %d, params = %s", statusCode, c.GetParams())
			if len(error) > 0 {
				sgxLogger.Warnf("code = %d, error = %s", statusCode, error[0])
			}

			hints.RenderResult(c, statusCode)
			return
		}

		if result == nil {
			sgxLogger.Warnf("callback result map is nil. params= %s", c.GetParams())
			hints.RenderResult(c, hints.Login02)
			return
		}

		openId, found := result.GetString("open_id")
		if found == false {
			sgxLogger.Warnf("callback result map not found `open_id`. result = %s", result)
			hints.RenderResult(c, hints.Login02)
			return
		}

		base64Token := token.New(pid, openId, config.Salt).ToBase64()
		hints.RenderResult(c, code.OK, base64Token)
	})
}

// severList 区服列表
// http://127.0.0.1:8089/server/list/2126001
func (p *Controller) serverList(c *sgxGin.Context) {
	pid := c.GetInt32("pid", 2126001)

	if pid < 1 {
		sgxLogger.Warnf("if pid < 1 {. params=%v", c.GetParams())
		hints.RenderResult(c, hints.Service002)
		return
	}

	areaGroup, found := data2.AreaGroupConfig.Get(pid)
	if found == false {
		hints.RenderResult(c, hints.Service002)
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

	hints.RenderResult(c, code.OK, dataList)
}

// severList 区服列表
// http://127.0.0.1:8089/pid/list
func (p *Controller) pidList(c *sgxGin.Context) {
	pidList := &struct {
		Pids []int32 `json:"pids"`
	}{}
	pidList.Pids = data2.SdkConfig.GetPidList()
	hints.RenderResult(c, code.OK, pidList)

}
