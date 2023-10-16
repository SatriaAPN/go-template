package handlergrpc_test

import (
	"context"
	dtousecase "go-template/dto/general/usecase"
	handlergrpc "go-template/handler/grpc"
	"go-template/mocks"
	"go-template/pb"
	"testing"

	"github.com/stretchr/testify/assert"
)

type usecaseMocks struct {
	MockUserUsecase *mocks.UserUsecase
}

type handlers struct {
	UserHandler handlergrpc.UserHandler
}

func initHandlers() (*usecaseMocks, *handlers) {
	um := usecaseMocks{
		MockUserUsecase: &mocks.UserUsecase{},
	}

	uhc := handlergrpc.UserHandlerConfig{
		UserUsecase: um.MockUserUsecase,
	}

	h := handlers{
		UserHandler: handlergrpc.NewUserHandler(uhc),
	}

	return &um, &h
}

func TestUserHandler_Login(t *testing.T) {
	assert := assert.New(t)
	t.Run("should success when login with right data", func(t *testing.T) {
		um, h := initHandlers()
		ctx := context.Background()
		req := pb.LoginRequest{
			Email:    "satria@mail.com",
			Password: "password",
		}
		expctRes := pb.LoginResponse{
			Message: "success",
			Token:   "thisistoken",
		}
		uReq := dtousecase.LoginUserRequest{
			Email:    req.Email,
			Password: req.Password,
		}
		uRes := dtousecase.LoginUserResponse{
			Token: "thisistoken",
		}

		um.MockUserUsecase.On("LoginUser", ctx, uReq).Return(uRes, nil)

		res, err := h.UserHandler.Login(ctx, &req)

		assert.Nil(err)
		assert.Equal(expctRes.Message, res.Message)
		assert.Equal(expctRes.Token, res.Token)
	})
}
