package interceptor

import "github.com/google/wire"

var Set = wire.NewSet(NewGrpcInterceptor)
