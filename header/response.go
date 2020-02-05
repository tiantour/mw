package header

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// Response response
type Response struct{}

// NewResponse new response
func NewResponse() *Response {
	return &Response{}
}

// Set set response metadata
func (r *Response) Set(ctx context.Context, kv ...string) context.Context {
	md := metadata.Pairs(kv...)
	if outMD, ok := metadata.FromIncomingContext(ctx); ok {
		md = metadata.Join(md, outMD)
	}
	return metadata.NewOutgoingContext(ctx, md)
}
