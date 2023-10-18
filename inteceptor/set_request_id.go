package interceptor

import (
	"context"
	dto "go-template/dto/general"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func SetRequestId(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	uuid := uuid.NewString()
	ctx = context.WithValue(ctx, dto.RequestIdKey, uuid)

	resp, err := handler(ctx, req)

	return resp, err
}
