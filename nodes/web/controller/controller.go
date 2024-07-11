package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	superGin "github.com/po2656233/superplace/components/gin"
	"github.com/po2656233/superplace/const/code"
	exString "github.com/po2656233/superplace/extend/string"
	sgxLogger "github.com/po2656233/superplace/logger"
	"net/http"
	"strings"
	data2 "superman/internal/conf"
	. "superman/internal/constant"
	. "superman/internal/hints"
	pb "superman/internal/protocol/gofile"
	"superman/internal/redis"
	"superman/internal/rpc"
	"superman/nodes/web/sdk"
	"time"
)

var jwtKey []byte = []byte("SECRETKEY")

const (
	TOKEN_MAX_EXPIRE_HOUR      = 1  // token最长有效期
	TOKEN_MAX_REMAINING_MINUTE = 15 // token还有多久过期就返回新token
)

type customClaims struct {
	Username string `json:"username"`
	Password bool   `json:"password"`
	jwt.RegisteredClaims
}

func verifyToken(ctx *superGin.Context, tokenString string) (int, string) {
	token1, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return http.StatusUnauthorized, fmt.Sprintf("access token parse error: %v.", err)
	}
	if claims, ok := token1.Claims.(*customClaims); ok && token1.Valid {
		if time.Now().After(claims.ExpiresAt.Time) {
			return http.StatusRequestTimeout, "access token expired"
		}
		// 即将超过过期时间，则添加一个http header `new-token` 给前端更新
		if t := claims.ExpiresAt.Time.Add(-time.Minute); t.Before(time.Now()) {
			claims = &customClaims{
				Username: claims.Username,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(TOKEN_MAX_EXPIRE_HOUR * time.Hour)},
				},
			}
			token1 = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, _ = token1.SignedString(jwtKey)
			ctx.Header("new-token", tokenString)
		}
		ctx.Set("claims", claims)
	} else {
		return http.StatusForbidden, fmt.Sprintf("Claims parse error: %v", err)
	}
	return http.StatusOK, ""
}

// AuthRequired gin jwt 认证中间件
func AuthRequired() superGin.GinHandlerFunc {
	return func(ctx *superGin.Context) {
		uri := ctx.Request.RequestURI
		if uri == "/" || uri == "/register" || uri == "/login" {
			return
		}
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		retCode, msg := verifyToken(ctx, tokenString)
		if retCode != http.StatusOK {
			ctx.AbortWithStatusJSON(retCode, gin.H{"code": -1, "message": msg})
			return
		}
		ctx.Next()
	}
}

type Controller struct {
	superGin.BaseController
}

func (p *Controller) Init() {
	group := p.Group("/")
	group.GET("/", p.index)
	group.GET("/hello", p.hello)
	group.GET("/pid/list", p.pidList)
	group.GET("/server/list/:pid", p.serverList)
	group.POST("/register", p.register)
	group.POST("/login", p.login)
}

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

// register 开发模式帐号注册
// http://127.0.0.1:8089/register?account=test11&password=test11
func (p *Controller) register(c *superGin.Context) {
	accountName := c.PostString(Username, "")
	password := c.PostString(Password, "")
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
func (p *Controller) login(c *superGin.Context) {
	pid := c.PostInt32(ProcessId, 0)
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
	params[ProcessId] = exString.ToString(pid)
	// invoke login.
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
		// 从redis上获取token,如果获取到token,且有效时长超过5分钟的，则返回原token.
		userKey := GetTokenKey(params[Username])
		ret := redis.SingleRedis().DB.Get(c.Context, userKey)
		if ret.Err() == nil && ret.String() != "" { // 返回现有token
			strToken := ret.Val()
			//校验token时长
			retCode, _ := verifyToken(c, strToken)
			if retCode == http.StatusOK {
				RenderResult(c, http.StatusOK, strToken)
				return
			}
		}

		// 生成一个JWT Token
		tokenObj := jwt.New(jwt.SigningMethodHS256)
		claims := tokenObj.Claims.(jwt.MapClaims)
		claims["sub"] = "UserLogin"                                                     // subject
		claims["name"] = params[Username]                                               // name
		claims["iat"] = time.Now().Unix()                                               // issued at
		claims["exp"] = time.Now().Add(time.Minute * TOKEN_MAX_REMAINING_MINUTE).Unix() // expiration time
		// 签名Token
		signedToken, err := tokenObj.SignedString([]byte(jwtKey))
		if err != nil {
			fmt.Println("Error signing token:", err)
			return
		}

		// 返回Bearer Token
		redis.SingleRedis().DB.Set(c.Context, userKey, signedToken, time.Minute*TOKEN_MAX_REMAINING_MINUTE)
		RenderResult(c, http.StatusOK, signedToken)
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
