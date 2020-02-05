package oauth

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/tiantour/mw/header"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type (
	// Oauth oauth
	Oauth struct{}

	// User user
	User struct {
		Number     int32  `json:"number,omitempty"`     // 编号
		Name       string `json:"name,omitempty"`       // 名称
		Avatar     string `json:"avatar,omitempty"`     // 头像
		Gender     int32  `json:"gender,omitempty"`     // 性别
		Permission int32  `json:"permission,omitempty"` // 权限
		Token      string `json:"token,omitempty"`      // 令牌
	}
)

// NewOauth new oauth
func NewOauth() *Oauth {
	return &Oauth{}
}

// Verify oauth verify
func (o *Oauth) Verify(ctx context.Context, method string) (context.Context, error) {
	reg := regexp.MustCompile("Service[A-Z]")
	switch reg.FindString(method) {
	case "ServiceF":
		return ctx, nil
	case "ServiceU":
		return o.do(ctx, 0)
	case "ServiceM":
		return o.do(ctx, 2)
	default:
		ctx, nil
	}
}

// do oauth do
func (o *Oauth) do(ctx context.Context, permission int32) (context.Context, error) {
	token := header.NewRequest().Authorization(ctx)
	if !strings.HasPrefix(token, "Bearer ") {
		err := errors.New("令牌缺失")
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	user, err := NewToken().Get(token[7:])
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	if user.Permission < permission {
		err := errors.New("权限不足")
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	ctx = context.WithValue(ctx, interface{}("user"), user)
	return ctx, nil
}
