package base

/*
# -*- coding: utf-8 -*-
# @Time : 2020/5/7 11:27
# @Author : Pitter
# @File : token.go
# @Software: GoLand
*/
import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 一些常量
// TokenExpired     error  = errors.New("Token is expired")
// TokenNotValidYet error  = errors.New("Token not active yet")
// TokenMalformed   error  = errors.New("That's not even a token")
var (
	TokenInvalid = errors.New("couldn't handle this token")
	SignKey      = "@*sa.stonewareWjTTs$#"
)

type JWT struct {
	SigningKey string
}

type CustomClaims struct {
	ID      int64  `json:"userId"` //用户ID
	Account string `json:"account"`
	PlatId  int64  `json:"platId"`
	*jwt.StandardClaims
}

// 创建token
func CreateTokenHs256(claims CustomClaims) (string, error) {
	//token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims = claims
	//signKey, _ := base64.RawStdEncoding.DecodeString(j.SigningKey)
	//res, err := token.SignedString(signKey)
	//fmt.Println("err:", err)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SignKey))
	//return token, err
	return token, err
}

func ParseTokenHs256(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

// RefreshToken 更新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return CreateTokenHs256(*claims)
	}

	return "", TokenInvalid
}

//func Login(w http.ResponseWriter, r *http.Request) {
//	//generateToken(w)
//}

//func generateToken(w http.ResponseWriter) {
//	j := &JWT{"man"}
//	claims := CustomClaims{
//		1, "Jaya", 123456, jwt.StandardClaims{
//			NotBefore: int64(time.Now().Unix() - 1000),
//			ExpiresAt: int64(time.Now().Unix() + 3600),
//			Issuer:    "man",
//		},
//	}
//
//	token, err := j.CreateToken(claims)
//	if err != nil {
//		io.WriteString(w, "it is wrong")
//	}
//
//	io.WriteString(w, token)
//}
