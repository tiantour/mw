package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// Unary unary
type Unary struct{}

// NewUnary new unary
func NewUnary() *Unary {
	return &Unary{}
}

// Client unary client
func (u *Unary) Client(nxt Next) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, conn *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, conn, opts...)
	}
}

// Server unary server
func (u *Unary) Server(nxt Next) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println(ctx, info.FullMethod)
		next, err := nxt(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(next, req)
	}
}
