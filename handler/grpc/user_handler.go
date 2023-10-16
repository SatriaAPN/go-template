package handlergrpc

import (
	"context"
	"fmt"
	dtousecase "go-template/dto/general/usecase"
	"go-template/pb"
	"go-template/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler interface {
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	pb.UnsafeAuthServer
}

type userHandler struct {
	pb.UnsafeAuthServer
	userUsecase usecase.UserUsecase
}

type UserHandlerConfig struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(config UserHandlerConfig) UserHandler {
	uh := &userHandler{}

	if config.UserUsecase != nil {
		uh.userUsecase = config.UserUsecase
	}

	return uh
}

func (uh *userHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res := &pb.LoginResponse{}

	fmt.Println(req)

	uRes, err := uh.userUsecase.LoginUser(ctx, dtousecase.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		e := status.Errorf(codes.InvalidArgument, err.Error())
		return res, e
	}
	res.Message = "success"
	res.Token = uRes.Token

	return res, nil
}
