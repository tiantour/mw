package oauth

import (
	"errors"
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
	body := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		data,
		jwt.RegisteredClaims{
			Issuer:    conf.NewToken().Data.Issuer,                            // 1.1可选，发行者
			Subject:   data.Token,                                             // 1.2可选，主体
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 1.4可选，到期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                         // 1.5可选，生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                         // 1.6可选，发布时间
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
		return nil, errors.New("令牌错误")
	}
	data, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("令牌过期")
	}
	return data.User, nil
}
