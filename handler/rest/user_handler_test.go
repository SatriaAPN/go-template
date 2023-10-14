package handlerrest_test

import (
	"encoding/json"
	"fmt"
	dto "go-template/dto/general"
	dtousecase "go-template/dto/general/usecase"
	dtohttp "go-template/dto/http"
	handlerrest "go-template/handler/rest"
	"go-template/mocks"
	routerrest "go-template/server/rest/router"
	"go-template/share/general/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// buat yg post
func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

func serveUserHandler(req *http.Request, rc routerrest.RouterConfig) *httptest.ResponseRecorder {
	router := routerrest.NewRouter(rc)
	fmt.Println(router)
	rec := httptest.NewRecorder()
	router.GetRouter().ServeHTTP(rec, req)
	return rec
}

func TestUserHandler_CreateUser(t *testing.T) {
	assert := assert.New(t)
	var (
		mockUserUsecase = new(mocks.UserUsecase)
		uhc             = handlerrest.UserHandlerConfig{
			UserUsecase: mockUserUsecase,
		}
		uh = handlerrest.NewUserHandler(uhc)

		rc = routerrest.RouterConfig{
			UserHandler: uh,
		}
	)
	t.Run("should success when create user", func(t *testing.T) {
		var (
			r = dtohttp.CreateUserRequest{
				Name:     "satria",
				Email:    "satria@mail.com",
				Password: "password",
			}
			r1 = dtousecase.CreateUserRequest{
				Name:     r.Name,
				Email:    r.Email,
				Password: r.Password,
			}
			rs1 = dtousecase.CreateUserResponse{
				Name:  r1.Name,
				Email: r1.Email,
			}
		)

		mockUserUsecase.On("CreateUser", mock.Anything, r1).Return(rs1, nil)

		req, _ := http.NewRequest(http.MethodPost, "/users", MakeRequestBody(r))
		rec := serveUserHandler(req, rc)

		assert.Equal(http.StatusOK, rec.Code)
	})
}

func TestUserHandler_UserLogin(t *testing.T) {
	assert := assert.New(t)
	var (
		mockUserUsecase = new(mocks.UserUsecase)
		uhc             = handlerrest.UserHandlerConfig{
			UserUsecase: mockUserUsecase,
		}
		uh = handlerrest.NewUserHandler(uhc)

		rc = routerrest.RouterConfig{
			UserHandler: uh,
		}
	)
	t.Run("should success when login", func(t *testing.T) {
		var (
			u = dtousecase.LoginUserRequest{
				Email:    "satria@mail.com",
				Password: "password",
			}
			res = dtousecase.LoginUserResponse{Token: "bearer thisisjwttoken"}
		)

		mockUserUsecase.On("LoginUser", mock.Anything, u).Return(res, nil)
		req, _ := http.NewRequest(http.MethodPost, "/login", MakeRequestBody(u))
		rec := serveUserHandler(req, rc)

		assert.Equal(http.StatusOK, rec.Code)
	})
}

func TestUserHandler_GetProfile(t *testing.T) {
	assert := assert.New(t)
	var (
		mockUserUsecase = new(mocks.UserUsecase)
		uhc             = handlerrest.UserHandlerConfig{
			UserUsecase: mockUserUsecase,
		}
		uh = handlerrest.NewUserHandler(uhc)
		rc = routerrest.RouterConfig{
			UserHandler: uh,
		}

		at, _ = util.GetAuthTokenGenerator().Encode(dto.AuthData{ID: 1})
		at2   = "Bearer " + at
	)
	t.Run("should get user profile", func(t *testing.T) {

		mockUserUsecase.On("GetUserProfile", mock.Anything, dtousecase.GetUserProfileRequest{UserId: 1}).Return(dtousecase.ProfileUserResponse{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/profile", nil)
		req.Header.Set("Authorization", at2)
		rec := serveUserHandler(req, rc)

		assert.Equal(http.StatusOK, rec.Code)
	})
}

func TestUserHandler_ForgetPassword(t *testing.T) {
	assert := assert.New(t)
	var (
		mockUserUsecase = new(mocks.UserUsecase)
		uhc             = handlerrest.UserHandlerConfig{
			UserUsecase: mockUserUsecase,
		}
		uh = handlerrest.NewUserHandler(uhc)

		rc = routerrest.RouterConfig{
			UserHandler: uh,
		}
	)
	t.Run("should success when forget-password", func(t *testing.T) {
		var (
			u = dtohttp.ForgetPasswordRequest{
				Email: "satria@mail.com",
			}
		)

		mockUserUsecase.On("ForgetPassword", mock.Anything, dtousecase.ForgetPasswordRequest{
			Email: u.Email,
		}).Return(dtousecase.ForgetPasswordResponse{Token: "thisss", ExpiredAt: time.Now()}, nil)

		req, _ := http.NewRequest(http.MethodPost, "/forget-password", MakeRequestBody(u))
		rec := serveUserHandler(req, rc)

		assert.Equal(http.StatusOK, rec.Code)
	})
}
