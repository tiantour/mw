package adapter

import (
	"github.com/tiantour/mw/protector"
	"google.golang.org/grpc/resolver"
)

func init() {
	b := protector.NewResolver()
	resolver.Register(b)
}
