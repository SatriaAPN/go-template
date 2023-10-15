package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	SetRequestId(&ctx)

	if !blacklistMethodAuth(info.FullMethod) {
		err := auth(&ctx, &req, info)

		if err != nil {
			return nil, err
		}
	}

	resp, err := handler(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
