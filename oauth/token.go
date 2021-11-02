package oauth

import (
	"errors"
	"log"
	"time"

	"gitee.com/tiantour/account/pb/user"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tiantour/conf"
)

type (
	// Token token
	Token struct{}

	// Claims claims
	Claims struct {
		*user.User
		jwt.RegisteredClaims
	}
)

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Set set user to jwt token
// date 2016-12-17
// author andy.jiang
func (t *Token) Set(data *user.User) (*user.User, error) {
	now := time.Now()
	body := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		data,
		jwt.RegisteredClaims{
			Issuer:    conf.NewToken().Data.Issuer,                     // 1.1可选，发行者
			Subject:   data.Token,                                      // 1.2可选，主体
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)), // 1.4可选，到期时间
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	})
	secret := []byte(conf.NewToken().Data.Secret)
	token, err := body.SignedString(secret)
	if err != nil {
		return nil, err
	}
	data.Token = token
	return data, nil
}

// Get get user for jwt token
// date 2016-12-17
// author andy.jiang
func (t *Token) Get(sign string) (*user.User, error) {
	token, err := jwt.ParseWithClaims(sign, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		secret := conf.NewToken().Data.Secret
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("令牌错误")
	}
	data, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("令牌过期")
	}
	return data.User, nil
}
