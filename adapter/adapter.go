package adapter

import (
	"context"
	"time"
)

var (
	// CTX grpc context
	CTX = context.Background()

	// TTL timeout
	TTL = time.Second * 5
)
