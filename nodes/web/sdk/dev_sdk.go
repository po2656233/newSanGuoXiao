package sdk

import (
	sgxGin "github.com/po2656233/superplace/components/gin"
	"github.com/po2656233/superplace/const/code"
	sgxString "github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	sgxError "github.com/po2656233/superplace/logger/error"
	"sanguoxiao/internal/conf"
	"sanguoxiao/internal/hints"
	pb "sanguoxiao/internal/protocol/gofile"
	"sanguoxiao/internal/rpc"
)

type devSdk struct {
	app cfacade.IApplication
}

func (devSdk) SdkId() int32 {
	return DevMode
}

func (p devSdk) Login(_ *conf.SdkRow, params Params, callback Callback) {
	accountName, _ := params.GetString("account")
	password, _ := params.GetString("password")

	if accountName == "" || password == "" {
		err := sgxError.Errorf("account or password params is empty.")
		callback(hints.Login08, nil, err)
		return
	}

	resp, errCode := rpc.SendData(p.app, rpc.SourcePath, rpc.AccountActor, rpc.CenterType, &pb.LoginReq{
		Account:  accountName,
		Password: password,
	})
	if errCode != code.OK {
		callback(errCode, nil)
		return
	}
	loginRet := resp.(*pb.LoginResp)
	callback(code.OK, map[string]string{
		"open_id": sgxString.ToString(loginRet.MainInfo.UserInfo.UserID),
	})
}

func (devSdk) PayCallback(_ *conf.SdkRow, _ *sgxGin.Context) {
}
