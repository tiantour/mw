package adapter

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// Target target
type Target struct{}

// NewTarget new target
func NewTarget() *Target {
	return &Target{}
}

// Dial target dial
func (t *Target) Dial(target string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(CTX, TTL)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("%s did not connect: %v", target, err)
	}
	return conn
}
