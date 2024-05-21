package main

import (
	"fmt"
	"testing"
)

func TestSendAndVerify(t *testing.T) {
	mailUser := "po2656233@qq.com"
	mailPass := "doaqmadsooivbced" // 授权码
	mailTo := "2437854119@qq.com"
	codeLen := 4
	codeTTL := 5 * 60
	from := "login" // 业务场景标记

	code := GenRandomCode(codeLen)
	fmt.Println("当前获取到的验证码", code)
	subjectText := "登录验证码"                              // 邮件主题
	bodyText := fmt.Sprintf("您的验证码是:%s， 有效期为5分钟", code) // 邮件正文

	options := &MailOptions{
		MailUser: mailUser,
		MailPass: mailPass,
		MailTo:   mailTo,
		Subject:  subjectText,
		Body:     bodyText,
	}

	err := SendMailCode(options, code, from, "", codeTTL)
	if err != nil {
		t.Error("SendMailCode error", err)
	}

	err = ValidateMailCode(mailTo, code, from)
	if err != nil {
		t.Error("ValidateMailCode error", err)
	}
}
