package sdk

import (
	sgxGin "github.com/po2656233/superplace/components/gin"
	"github.com/po2656233/superplace/const/code"
	sgxHttp "github.com/po2656233/superplace/extend/http"
	cstring "github.com/po2656233/superplace/extend/string"
	sgxLogger "github.com/po2656233/superplace/logger"
	cerror "github.com/po2656233/superplace/logger/error"
	"superman/internal/conf"
	. "superman/internal/constant"
)

type (
	quickSdk struct {
	}
)

func (quickSdk) SdkId() int32 {
	return QuickSDK
}

func (quickSdk) Login(config *conf.SdkRow, params Params, callback Callback) {
	token, found := params.GetString("token")
	if found == false || cstring.IsBlank(token) {
		err := cerror.Error("token is nil")
		callback(Login11, nil, err)
		return
	}

	uid, found := params.GetString("uid")
	if found == false || cstring.IsBlank(uid) {
		err := cerror.Error("uid is nil")
		callback(Login02, nil, err)
		return
	}

	rspBody, _, err := sgxHttp.GET(config.LoginURL(), map[string]string{
		"token":        token,
		"uid":          uid,
		"product_code": config.GetString("productCode"),
	})

	if err != nil || rspBody == nil {
		callback(Login02, nil, err)
		return
	}

	var result = string(rspBody)
	if result != "1" {
		sgxLogger.Warnf("quick sdk login fail. rsp =%s", rspBody)
		callback(Login02, nil, err)
		return
	}

	callback(code.OK, map[string]string{
		OpenID: uid, //返回 quick的uid做为 open id
	})
}

func (s quickSdk) PayCallback(_ *conf.SdkRow, c *sgxGin.Context) {
	// TODO 这里实现quick sdk 支付回调的逻辑
	c.RenderHTML("FAIL")
}
