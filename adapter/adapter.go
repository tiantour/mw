package adapter

import (
	"context"
	"time"
)

var (
	// CTX grpc context
	CTX = context.Background()

	// TTL grpc timeout
	TTL = time.Second * 5
)
