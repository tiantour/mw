package oauth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tiantour/mw/header"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Oauth oauth
type Oauth struct{}

// NewOauth new oauth
func NewOauth() *Oauth {
	return &Oauth{}
}

// Verify oauth verify
func (o *Oauth) Verify(ctx context.Context, method string) (context.Context, error) {
	fmt.Println(1, method, ctx)
	if strings.Contains(method, "ServiceU") {
		return o.do(ctx, 0)
	} else if strings.Contains(method, "ServiceM") {
		return o.do(ctx, 2)
	}
	return ctx, nil
}

// do oauth do
func (o *Oauth) do(ctx context.Context, permission int32) (context.Context, error) {
	token := header.NewRequest().Authorization(ctx)
	fmt.Println(2, token)
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
