package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "1474545380 <1474545380@qq.com>"
	e.To = []string{"liulongxin1819@outlook.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码是<b>123123</b>")
	//err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1474545380@qq.com", "mgepgzlrjrnmjeic", "smtp.qq.com"))
	//返回EOF时关闭SSL重试
	err := e.SendWithTLS("smtp.qq.com:587", smtp.PlainAuth("", "1474545380@qq.com", "mgepgzlrjrnmjeic", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
