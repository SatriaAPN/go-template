package servergrpc

import (
	"go-template/database"
	handlergrpc "go-template/handler/grpc"
	"go-template/pb"
	"go-template/repository"
	"go-template/share/general/util"
	interceptor "go-template/share/grpc/inteceptor"
	"go-template/usecase"

	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	db := database.GetInstance()

	urc := repository.UserRepositoryConfig{
		Db: db,
	}
	ur := repository.NewUserRepository(urc)

	uuc := usecase.UserUsecaseConfig{
		UserRepository:       ur,
		PasswordHasher:       util.GetPasswordHasher(),
		AuthTokenGenerator:   util.GetAuthTokenGenerator(),
		RandomTokenGenerator: util.GetRandomTokenGenerator(),
	}
	uu := usecase.NewUserUsecase(uuc)

	uh := handlergrpc.NewUserHandler(uu)

	// setup the server
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryInterceptor),
	)

	// register the handler to server
	pb.RegisterAuthServer(server, uh)

	return server
}
