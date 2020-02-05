package oauth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/tiantour/conf"
	"github.com/tiantour/tempo"
)

type (
	// Token token
	Token struct{}

	// Claims claims
	Claims struct {
		*User
		jwt.StandardClaims
	}
)

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Set set user to jwt token
// date 2016-12-17
// author andy.jiang
func (t *Token) Set(data *User) (*User, error) {
	body := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		data,
		jwt.StandardClaims{
			Issuer:    conf.NewToken().Data.Issuer,       // 1.1可选，发行者
			Subject:   data.Token,                        // 1.2可选，主体
			ExpiresAt: tempo.NewNow().Unix() + 7*24*3600, // 1.4可选，到期时间
		},
	})
	secret := []byte(conf.NewToken().Data.Secret)
	token, err := body.SignedString(secret)
	if err != nil {
		return data, err
	}
	data.Token = token
	return data, nil
}

// Get get user for jwt token
// date 2016-12-17
// author andy.jiang
func (t *Token) Get(sign string) (*User, error) {
	token, err := jwt.ParseWithClaims(sign, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			secret := []byte(conf.NewToken().Data.Secret)
			return secret, nil
		},
	)
	if err != nil {
		return nil, errors.New("令牌错误")
	}
	data, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("令牌过期")
	}
	return data.User, nil
}
