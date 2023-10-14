package server

import (
	"go-template/database"
	"go-template/handler"
	"go-template/repository"
	"go-template/server/router"
	utilhttp "go-template/share/http/util"
	"go-template/share/util"
	"go-template/usecase"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func NewServer() *http.Server {

	r := router.NewRouter(initHandler()).GetRouter()

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &s
}

func InitServer() {
	godotenv.Load()

	srv := NewServer()

	utilhttp.GracefulShutdown(srv)
}

func initHandler() router.RouterConfig {
	db := database.GetInstance()

	nrc := router.RouterConfig{
		UserHandler: initUserHandler(db),
	}

	return nrc
}

func initUserHandler(db *gorm.DB) handler.UserHandler {
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

	uhc := handler.UserHandlerConfig{
		UserUsecase: uu,
	}
	uh := handler.NewUserHandler(uhc)

	return uh
}
