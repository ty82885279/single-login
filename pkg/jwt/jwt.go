package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

var MySecret = []byte("夏天夏天悄悄过去留下笑眯眯")

var (
	ErrorInvalidToken   = errors.New("AToken已过期")
	ErrorATokenNotExpir = errors.New("AToken尚未过期")
	ErrorRTokenHasExpir = errors.New("RToken已过期")
)

const TokenExpireDuration = time.Minute * 1

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return MySecret, nil
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int64 `json:"id"`
	jwt.StandardClaims
}

func GenToken(userid int64) (atoken, rtoken string, err error) {

	c := MyClaims{
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_atoken_expire")) * time.Minute).Unix(),
			Issuer:    "bluebell",
		},
	}
	//使用制定的签名方法创建签名对象
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)

	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("jwt_rtoken_expire")) * time.Minute).Unix(),
		Issuer:    "bluebell",
	}).SignedString(MySecret)
	fmt.Printf("time----:%d\n", viper.GetInt("jwt_rtoken_expire"))
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, keyFunc)
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return mc, ErrorInvalidToken
}

//  RefreshToken 刷新Token
func RefreshToken(userID int64) (newAToken, newRToken string) {

	newAToken, newRToken, _ = GenToken(userID)

	return
}
