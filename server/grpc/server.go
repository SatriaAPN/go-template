package servergrpc

import (
	"go-template/database"
	handlergrpc "go-template/handler/grpc"
	interceptor "go-template/inteceptor"
	"go-template/pb"
	"go-template/repository"
	"go-template/share/general/config"
	"go-template/share/general/util"
	"go-template/usecase"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func NewServer() *grpc.Server {
	config.InitEnvReader()

	db := database.GetInstance()

	// setup the server
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.SetRequestId,
			interceptor.ErrorHandling,
			interceptor.Logger,
			interceptor.Auth,
		),
	)

	// register the handler to server
	pb.RegisterAuthServer(server, initUserHandler(db))

	return server
}

func initUserHandler(db *gorm.DB) handlergrpc.UserHandler {

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

	uhc := handlergrpc.UserHandlerConfig{
		UserUsecase: uu,
	}
	uh := handlergrpc.NewUserHandler(uhc)

	return uh
}
