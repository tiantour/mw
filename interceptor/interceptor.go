package interceptor

import (
	"context"
)

// Next next
type Next func(ctx context.Context, method string) (context.Context, error)
