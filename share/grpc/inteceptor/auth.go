package interceptor

import (
	"context"
	"errors"
	dto "go-template/dto/general"
	"go-template/share/general/util"
	utilgrpc "go-template/share/grpc/util"
	"strings"

	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !blacklistMethodAuth(info.FullMethod) {
		authorizationHeader, err := utilgrpc.GetAuthTokenFromGrpcContext(&ctx)

		if err != nil {
			return nil, err
		}

		if !strings.Contains(authorizationHeader, "Bearer") {
			return nil, errors.New("invalid token")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		ud, err := util.GetAuthTokenGenerator().Decode(tokenString)

		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, dto.AuthDataKey, ud)

		resp, err := handler(ctx, req)

		return resp, err
	} else {
		resp, err := handler(ctx, req)

		return resp, err
	}
}

// return true if the method doesn't allowed to use authorization
func blacklistMethodAuth(fullmethod string) bool {
	switch fullmethod {
	case "/auth.Auth/Login":
		return true
	}
	return false
}
