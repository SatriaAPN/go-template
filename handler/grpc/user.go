package handlergrpc

import (
	"context"
	"fmt"
	dtousecase "go-template/dto/usecase"
	"go-template/pb"
	"go-template/usecase"
)

type UserHandler interface {
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	pb.UnsafeAuthServer
}

type userHandler struct {
	pb.UnsafeAuthServer
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	uh := &userHandler{
		userUsecase: uu,
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
		return res, err
	}
	res.Message = "success"
	res.Token = uRes.Token

	return res, nil
}
