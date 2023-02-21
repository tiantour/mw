package oauth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	// Token token
	Token struct{}

	// User user
	User struct {
		Number     int32  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`         // 编号
		Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`              // 名称
		Avatar     string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`          // 头像
		Gender     int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`         // 性别
		Action     string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`          // 唯一值
		Password   string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`      // 密码
		Device     string `protobuf:"bytes,7,opt,name=device,proto3" json:"device,omitempty"`          // 地域
		Platform   int32  `protobuf:"varint,8,opt,name=platform,proto3" json:"platform,omitempty"`     // 平台
		Permission int32  `protobuf:"varint,9,opt,name=permission,proto3" json:"permission,omitempty"` // 权限
		Extend     string `protobuf:"bytes,10,opt,name=extend,proto3" json:"extend,omitempty"`         // 拓展
		Time       string `protobuf:"bytes,11,opt,name=time,proto3" json:"time,omitempty"`             // 时间
		Token      string `protobuf:"bytes,12,opt,name=token,proto3" json:"token,omitempty"`           // 令牌
	}
	// Claims claims
	Claims struct {
		*User
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
func (t *Token) Set(data *User) (*User, error) {
	body := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		data,
		jwt.RegisteredClaims{
			Issuer:    Issuer,                                                 // 1.1可选，发行者
			Subject:   data.Token,                                             // 1.2可选，主体
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 1.4可选，到期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                         // 1.5可选，生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                         // 1.6可选，发布时间
		},
	})
	secret := []byte(Secret)
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
func (t *Token) Get(sign string) (*User, error) {
	token, err := jwt.ParseWithClaims(sign, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
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
