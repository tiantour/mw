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

// Dail target dail
func (t *Target) Dail(target string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(CTX, TTL)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
