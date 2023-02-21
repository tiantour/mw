package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/tiantour/mw/protector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Target target
type Target struct{}

// NewTarget new target
func NewTarget() *Target {
	return &Target{}
}

// Dial target dial
func (t *Target) Dial(service string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	target := fmt.Sprintf("%s:///%s", protector.Scheme, service)
	return grpc.DialContext(
		ctx,
		target,
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
