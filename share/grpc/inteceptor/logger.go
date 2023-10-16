package interceptor

import (
	"context"
	dtogrpc "go-template/dto/grpc"
	"go-template/share/general/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	UnaryRequestLogging(ctx, req, info)

	resp, err := handler(ctx, req)

	UnaryResponseLogging(ctx, req, info, err)

	return resp, err
}

func UnaryRequestLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) {
	ld := dtogrpc.NewRequestGrpcLogger(info.FullMethod, ctx.Value("X-Request-Id").(string), "request-grpc")

	util.GetLogger().Infof(ld)
}

func UnaryResponseLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, err error) {
	s := status.Code(err)

	ld := dtogrpc.NewResponseGrpcLogger(info.FullMethod, ctx.Value("X-Request-Id").(string), "response-grpc", int(s))

	util.GetLogger().Infof(ld)
}
