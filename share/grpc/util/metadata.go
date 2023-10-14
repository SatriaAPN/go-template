package utilgrpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

func GetAuthTokenFromGrpcContext(ctx *context.Context) (string, error) {
	result := ""

	md, found := metadata.FromIncomingContext(*ctx)

	if !found {
		return "", errors.New("grpc context is empty")
	}

	authToken, found := md["authorization"]

	if !found {
		return "", errors.New("auth token didn't found")
	}

	result = authToken[0]

	return result, nil
}
