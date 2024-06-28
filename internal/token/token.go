package token

import (
	"encoding/json"
	"fmt"
	"github.com/po2656233/superplace/const/code"
	extendCrypto "github.com/po2656233/superplace/extend/crypto"
	sgxTime "github.com/po2656233/superplace/extend/time"
	sgxLogger "github.com/po2656233/superplace/logger"
	"superman/internal/hints"
)

const (
	hashFormat      = "pid:%d,openid:%s,timestamp:%d"
	tokenExpiredDay = 3
)

type Token struct {
	PID       int32  `json:"pid"`
	OpenID    string `json:"open_id"`
	Timestamp int64  `json:"tt"`
	Hash      string `json:"hash"`
}

func New(pid int32, openId string, appKey string) *Token {
	token := &Token{
		PID:       pid,
		OpenID:    openId,
		Timestamp: sgxTime.Now().ToMillisecond(),
	}

	token.Hash = BuildHash(token, appKey)
	return token
}

func (t *Token) ToBase64() string {
	bytes, _ := json.Marshal(t)
	return extendCrypto.Base64Encode(string(bytes))
}

func DecodeToken(base64Token string) (*Token, bool) {
	if len(base64Token) < 1 {
		return nil, false
	}

	token := &Token{}
	bytes, err := extendCrypto.Base64DecodeBytes(base64Token)
	if err != nil {
		sgxLogger.Warnf("base64Token = %s, validate error = %v", base64Token, err)
		return nil, false
	}

	err = json.Unmarshal(bytes, token)
	if err != nil {
		sgxLogger.Warnf("base64Token = %s, unmarshal error = %v", base64Token, err)
		return nil, false
	}

	return token, true
}

func Validate(token *Token, appKey string) (int32, bool) {
	now := sgxTime.Now()
	now.AddDays(tokenExpiredDay)

	if token.Timestamp > now.ToMillisecond() {
		sgxLogger.Warnf("token is expired, token = %s", token)
		return hints.Login11, false
	}

	newHash := BuildHash(token, appKey)
	if newHash != token.Hash {
		sgxLogger.Warnf("hash validate fail. newHash = %s, token = %s", token)
		return hints.Login15, false
	}

	return code.OK, true
}

func BuildHash(t *Token, appKey string) string {
	value := fmt.Sprintf(hashFormat, t.PID, t.OpenID, t.Timestamp)
	return extendCrypto.MD5(value + appKey)
}
