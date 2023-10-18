package interceptor

import (
	"context"
	dto "go-template/dto/general"
	dtogrpc "go-template/dto/grpc"
	"go-template/share/general/util"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	UnaryRequestLogging(ctx, req, info)
	t := time.Now()

	resp, err := handler(ctx, req)

	UnaryResponseLogging(ctx, req, info, err, t)

	return resp, err
}

func UnaryRequestLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) {
	ld := dtogrpc.NewRequestGrpcLogger(info.FullMethod, ctx.Value(dto.RequestIdKey).(string), "request-grpc")

	util.GetLogger().Infof(ld)
}

func UnaryResponseLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, err error, requestTime time.Time) {
	s := status.Code(err)

	timePassed := time.Since(requestTime)
	ld := dtogrpc.NewResponseGrpcLogger(info.FullMethod, ctx.Value(dto.RequestIdKey).(string), "response-grpc", int(s), timePassed)

	util.GetLogger().Infof(ld)
}
