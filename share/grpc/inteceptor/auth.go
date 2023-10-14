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

func auth(ctx *context.Context, req *interface{}, info *grpc.UnaryServerInfo) error {
	authorizationHeader, err := utilgrpc.GetAuthTokenFromGrpcContext(ctx)

	if err != nil {
		return err
	}

	if !strings.Contains(authorizationHeader, "Bearer") {
		return errors.New("invalid token")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	ud, err := util.GetAuthTokenGenerator().Decode(tokenString)

	if err != nil {
		return err
	}

	*ctx = context.WithValue(*ctx, dto.AuthDataKey, ud)

	return nil
}

// return true if the method doesn't allowed to use authorization
func blacklistMethodAuth(fullmethod string) bool {
	switch fullmethod {
	case "/Auth/Login":
		return true
	}
	return false
}
