package interceptor

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func SetRequestIdInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	uuid := uuid.NewString()
	ctx = context.WithValue(ctx, "X-Request-Id", uuid)

	resp, err := handler(ctx, req)

	return resp, err
}
