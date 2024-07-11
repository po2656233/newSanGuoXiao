package base

/*
# -*- coding: utf-8 -*-
# @Time : 2020/5/10 11:56
# @Author : Pitter
# @File : uuid.go
# @Software: GoLand
*/
import (
	"github.com/google/uuid"
	"sync/atomic"
	"time"
)

type UUID [16]byte

var timeBase = time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix()
var hardwareAddr = []byte("ChatServer") // 设备码
var clockSeq uint32

func GetUUID() UUID {
	return ToUUID(time.Now())
}

func ToUUID(aTime time.Time) UUID {
	var u UUID

	utcTime := aTime.In(time.UTC)
	t := int64(utcTime.Unix()-timeBase)*10000000 + int64(utcTime.Nanosecond()/100)
	u[0], u[1], u[2], u[3] = byte(t>>24), byte(t>>16), byte(t>>8), byte(t)
	u[4], u[5] = byte(t>>40), byte(t>>32)
	u[6], u[7] = byte(t>>56)&0x0F, byte(t>>48)

	clock := atomic.AddUint32(&clockSeq, 0)
	u[8] = byte(clock >> 8)
	u[9] = byte(clock)

	copy(u[10:], hardwareAddr)

	u[6] |= 0x10 // set version to 1 (time based uuid)
	u[8] &= 0x3F // clear variant
	u[8] |= 0x80 // set to IETF variant

	return u
}

func (u UUID) String() string {
	var offsets = [...]int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	const hexString = "0123456789abcdef"
	r := make([]byte, 36)
	for i, b := range u {
		r[offsets[i]] = hexString[b>>4]
		r[offsets[i]+1] = hexString[b&0xF]
	}
	r[8] = '-'
	r[13] = '-'
	r[18] = '-'
	r[23] = '-'
	return string(r)
}

//func init() {//测试用
//	println(GetGoogleUUID())
//}

func GetGoogleUUID() string {
	uid := uuid.New()
	println(`生成的UUID v4：`, uid.String())
	return uid.String()
	//
	//// 创建可以进行错误处理的 UUID v4
	//u2 := uuid.NewV5(u1, "server")
	//println(`生成的UUID v4：`)
	//println(u2.String())
	//
	//// 解析 字符串 到 UUID
	//u2, err2 := uuid.FromString(`6ba7b810-9dad-11d1-80b4-00c04fd430c8`)
	//if err2 != nil {
	//	println(`解析 字符串 到 UUID 时出错`)
	//	panic(err2)
	//}
	//println(`解析 字符串 到 UUID 成功！解析到的 UUID 如下：`)
	//println(u2.String())
	//return u2.String()

}
