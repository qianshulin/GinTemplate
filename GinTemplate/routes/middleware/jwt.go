package middleware

import (
	"GinTemplate/models"
	"github.com/kataras/jwt"

	"time"
)

var (
	secret = []byte("")
	expire = int64(0)
)

//InitJWT 初始化JWT
func InitJWT(accessSecret string, accessExpire int64) {
	secret = []byte(accessSecret)
	expire = accessExpire
}

//GenerateToken 生成Token
func GenerateToken(claims interface{}) (*models.JwtToken, error) {
	token, err := jwt.Sign(jwt.HS256, secret, claims, jwt.MaxAge(time.Duration(expire)*time.Minute))
	return &models.JwtToken{
		Token:  string(token),
		Expire: time.Now().Unix() + expire*60,
	}, err
}

//VerifyToken 校验token
func VerifyToken(token string) (*jwt.VerifiedToken, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, secret, []byte(token))
	return verifiedToken, err
}
