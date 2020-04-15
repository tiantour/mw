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
	omd := metadata.Pairs(kv...)
	if imd, ok := metadata.FromIncomingContext(ctx); ok {
		omd = metadata.Join(omd, imd)
	}
	if omd.Len() != 0 {
		return metadata.NewOutgoingContext(ctx, omd)
	}
	return ctx
}
