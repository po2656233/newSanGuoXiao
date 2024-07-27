package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	superGin "github.com/po2656233/superplace/components/gin"
	"net/http"
	"strconv"
	"time"
)

var jwtKey []byte = []byte("SECRETKEY")

const (
	TOKEN_MAX_EXPIRE_HOUR      = 1  // token最长有效期
	TOKEN_MAX_REMAINING_MINUTE = 72 //15 // token还有多久过期就返回新token
)

func VerifyToken(ctx *superGin.Context, tokenString string) (int, string) {
	token1, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return http.StatusUnauthorized, fmt.Sprintf("access token parse error: %v.", err)
	}
	if claims, ok := token1.Claims.(*jwt.RegisteredClaims); ok && token1.Valid {
		if time.Now().After(claims.ExpiresAt.Time) {
			return http.StatusRequestTimeout, "access token expired"
		}
		// 即将超过过期时间，则添加一个http header `new-token` 给前端更新
		if t := claims.ExpiresAt.Time.Add(-time.Minute); t.Before(time.Now()) {
			claims.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(TOKEN_MAX_EXPIRE_HOUR * time.Hour)}
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

func GetIDForToken(tokenString string) (int64, error) {
	token1, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return -1, fmt.Errorf("access token parse error: %v. ", err)
	}
	claims, ok := token1.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return -1, fmt.Errorf("token is broken")
	}
	return strconv.ParseInt(claims.ID, 10, 64)
}

func CreateToken(uid int64, username string, timeout time.Duration) (string, error) {
	// 生成一个JWT Token
	tokenObj := jwt.New(jwt.SigningMethodHS256)
	claims := tokenObj.Claims.(jwt.MapClaims)
	claims["sub"] = "UserLogin"                    // subject
	claims["iss"] = username                       // name
	claims["jti"] = fmt.Sprintf("%d", uid)         // ID
	claims["iat"] = time.Now().Unix()              // issued at
	claims["exp"] = time.Now().Add(timeout).Unix() // expiration time
	// 签名Token
	signedToken, err := tokenObj.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	return signedToken, nil
}
