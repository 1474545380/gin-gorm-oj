package help

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var myKey = []byte("liulongxin")

//md5生成密码
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

//生成token
func GenerateToken(identity, name string) (string, error) {
	UserClaim := UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("错误的token,err: %v", err)
	}
	return userClaim, nil
}

//发送验证码给单个用户
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "1474545380 <1474545380@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码是<b>" + code + "</b>")
	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1474545380@qq.com", "mgepgzlrjrnmjeic", "smtp.qq.com"))
	//返回EOF时关闭SSL重试
	//return e.SendWithTLS("smtp.qq.com:587", smtp.PlainAuth("", "1474545380@qq.com", "mgepgzlrjrnmjeic", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}

//生成UUID,user identity
func GetUUID() string {
	return uuid.NewV4().String()
}

//生成验证码
func GetRandom() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
