package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// Stream stream
type Stream struct {
	grpc.ServerStream
	ctx context.Context
}

// NewStream new stream
func NewStream() *Stream {
	return &Stream{}
}

// Client stream client
func (s *Stream) Client(nxt Next) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, conn *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(ctx, desc, conn, method, opts...)
	}
}

// Server stream server
func (s *Stream) Server(nxt Next) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		next, err := nxt(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, Stream{
			ServerStream: stream,
			ctx:          next,
		})
	}
}
