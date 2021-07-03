package oauth

import (
	"context"
	"errors"
	"strings"

	"github.com/tiantour/mw/header"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Oauth oauth
type Oauth struct{}

// NewOauth new oauth
func NewOauth() *Oauth {
	return &Oauth{}
}

// Verify oauth verify
func (o *Oauth) Verify(ctx context.Context, method string) (context.Context, error) {
	if strings.Contains(method, "ServiceU") {
		return o.do(ctx, 0)
	} else if strings.Contains(method, "ServiceM") {
		return o.do(ctx, 1)
	}
	return ctx, nil
}

// do oauth do
func (o *Oauth) do(ctx context.Context, permission int32) (context.Context, error) {
	token := header.NewRequest().Authorization(ctx)
	if !strings.HasPrefix(token, "Bearer ") {
		err := errors.New("令牌缺失")
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	user, err := NewToken().Get(token[7:])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	if user.Permission < permission {
		err := errors.New("权限不足")
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	if user.Permission != 0 {
		user.Number = 110119120
	}
	ctx = context.WithValue(ctx, interface{}("user"), user)
	return ctx, nil
}
