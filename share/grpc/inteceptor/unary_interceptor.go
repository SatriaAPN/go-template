package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	SetRequestId(&ctx)

	UnaryRequestLogging(ctx, req, info)

	if !blacklistMethodAuth(info.FullMethod) {
		err := auth(&ctx, &req, info)

		if err != nil {
			return nil, err
		}
	}

	resp, err := handler(ctx, req)

	UnaryResponseLogging(ctx, req, info, err)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
