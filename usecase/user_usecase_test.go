package usecase_test

import (
	"context"
	dto "go-template/dto/general"
	dtousecase "go-template/dto/general/usecase"
	"go-template/entity"
	"go-template/mocks"
	appconstant "go-template/share/general/constant"
	errorapp "go-template/share/general/error"
	"go-template/usecase"
	"testing"
	"time"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	assert := assert.New(t)
	t.Run("should success creating a user with right data", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.CreateUserRequest{Name: "satria", Email: "satria@mail.com", Password: "password"}
			res  = dtousecase.CreateUserResponse{Name: "satria", Email: "satria@mail.com"}
			req1 = req.Password
			res1 = "hashedPassword"
			req2 = entity.User{Name: req.Name, Email: req.Email, Password: res1}
			res2 = entity.User{Name: req.Name, Email: req.Email, Password: res1, ID: 1}
		)

		phmock.On("GenerateHashFromPassword", req1).Return(res1, nil)

		urmock.On("Create", ctx, req2).Return(res2, nil)

		cur, err := us.CreateUser(ctx, req)

		assert.Equal(res, cur)
		assert.Nil(err)
	})

	t.Run("should return error creating a user with invalid email", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us  = usecase.NewUserUsecase(usc)
			ctx = context.Background()
			req = dtousecase.CreateUserRequest{Name: "satria", Email: "satriamail.com", Password: "password"}
		)

		_, err := us.CreateUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrEmailIsNotValid)
	})

	t.Run("should return error creating a user with below minimum password length", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us  = usecase.NewUserUsecase(usc)
			ctx = context.Background()
			req = dtousecase.CreateUserRequest{Name: "satria", Email: "satria@mail.com", Password: "pass"}
		)

		_, err := us.CreateUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrMinimumPasswordLength)
	})

	t.Run("should return error creating a user with below maximum password length", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us  = usecase.NewUserUsecase(usc)
			ctx = context.Background()
			req = dtousecase.CreateUserRequest{Name: "satria", Email: "satria@mail.com", Password: "pasadwadawdawdwadawdawdawdawdawdawds"}
		)

		_, err := us.CreateUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrMaximumPasswordLength)
	})

	t.Run("should return error when password hasher return error", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.CreateUserRequest{Name: "satria", Email: "satria@mail.com", Password: "password"}
			req1 = req.Password
			res1 = "hashedPassword"
			err1 = errors.Errorf("error")
		)

		phmock.On("GenerateHashFromPassword", req1).Return(res1, err1)

		_, err := us.CreateUser(ctx, req)

		assert.ErrorIs(err, err1)
	})

	t.Run("should return error when email already register", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.CreateUserRequest{Name: "satria", Email: "satria@mail.com", Password: "password"}
			req1 = req.Password
			res1 = "hashedPassword"
			err1 = errors.New(errorapp.ErrEmailAlreadyExist)
			req2 = entity.User{Name: req.Name, Email: req.Email, Password: res1}
			res2 = entity.User{Name: req.Name, Email: req.Email, Password: res1, ID: 1}
		)

		phmock.On("GenerateHashFromPassword", req1).Return(res1, nil)

		urmock.On("Create", ctx, req2).Return(res2, err1)

		_, err := us.CreateUser(ctx, req)

		assert.ErrorIs(err, err1)
	})

}

func TestUserUsecase_LoginUser(t *testing.T) {
	assert := assert.New(t)
	t.Run("should success login with right data", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "password"}
			res  = dtousecase.LoginUserResponse{Token: "thisistoken"}
			req1 = req.Email
			res1 = entity.User{ID: 1, Email: req.Email, Password: "hashedPassword"}
			req2 = dto.AuthData{ID: res1.ID}
		)

		urmock.On("FindByEmail", ctx, req1).Return(res1, nil)
		phmock.On("CompareHashAndPassword", req.Password, res1.Password).Return(true, nil)
		atmock.On("Encode", req2).Return(res.Token, nil)

		lur, err := us.LoginUser(ctx, req)

		assert.Equal(res, lur)
		assert.Nil(err)
	})

	t.Run("should return error when email invalid", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us  = usecase.NewUserUsecase(usc)
			ctx = context.Background()
			req = dtousecase.LoginUserRequest{Email: "satriamail.com", Password: "password"}
		)

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrEmailIsNotValid)
	})

	t.Run("should return error when password below minimum length", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us  = usecase.NewUserUsecase(usc)
			ctx = context.Background()
			req = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "asda"}
		)

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrMinimumPasswordLength)
	})

	t.Run("should return error when password above maximum length", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us  = usecase.NewUserUsecase(usc)
			ctx = context.Background()
			req = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "asdawdawdsdawdsadawdasdawdasdawdasdawdawda"}
		)

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrMaximumPasswordLength)
	})

	t.Run("should success error when email not found", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "password"}
			req1 = req.Email
			res1 = entity.User{ID: 1, Email: req.Email, Password: "hashedPassword"}
		)

		urmock.On("FindByEmail", ctx, req1).Return(res1, errors.New(errorapp.ErrEmailNotFound))

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrEmailNotFound)
	})

	t.Run("should success error when repository didn't return data", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "password"}
			req1 = req.Email
			res1 = entity.User{ID: 0, Email: req.Email, Password: "hashedPassword"}
		)

		urmock.On("FindByEmail", ctx, req1).Return(res1, nil)

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrEmailNotFound)
	})

	t.Run("should return error when password doesn't match", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "password"}
			req1 = req.Email
			res1 = entity.User{ID: 1, Email: req.Email, Password: "hashedPassword"}
		)

		urmock.On("FindByEmail", ctx, req1).Return(res1, nil)
		phmock.On("CompareHashAndPassword", req.Password, res1.Password).Return(false, nil)

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrWrongPassword)
	})

	t.Run("should return error when password validator return error", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "password"}
			req1 = req.Email
			res1 = entity.User{ID: 1, Email: req.Email, Password: "hashedPassword"}
		)

		urmock.On("FindByEmail", ctx, req1).Return(res1, nil)
		phmock.On("CompareHashAndPassword", req.Password, res1.Password).Return(true, errors.New(errorapp.ErrWrongPassword))

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrWrongPassword)
	})

	t.Run("should return error when token encoder return error", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:     &urmock,
				PasswordHasher:     &phmock,
				AuthTokenGenerator: &atmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.LoginUserRequest{Email: "satria@mail.com", Password: "password"}
			res  = dtousecase.LoginUserResponse{Token: "thisistoken"}
			req1 = req.Email
			res1 = entity.User{ID: 1, Email: req.Email, Password: "hashedPassword"}
			req2 = dto.AuthData{ID: res1.ID}
		)

		urmock.On("FindByEmail", ctx, req1).Return(res1, nil)
		phmock.On("CompareHashAndPassword", req.Password, res1.Password).Return(true, nil)
		atmock.On("Encode", req2).Return(res.Token, errors.New(errorapp.ErrWrongPassword))

		_, err := us.LoginUser(ctx, req)

		assert.ErrorIs(err, errorapp.ErrWrongPassword)
	})
}

func TestUserUsecase_ForgetPassword(t *testing.T) {
	assert := assert.New(t)
	t.Run("should success call ForgetPassword with right data", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			rtmock = mocks.RandomTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:       &urmock,
				PasswordHasher:       &phmock,
				AuthTokenGenerator:   &atmock,
				RandomTokenGenerator: &rtmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.ForgetPasswordRequest{Email: "satria@mail.com"}
			res  = dtousecase.ForgetPasswordResponse{Token: "token1"}
			res1 = entity.User{ID: 1, Email: req.Email, Name: "satria"}
			res2 = entity.UserResetPassword{
				UserId:    int(res1.ID),
				Token:     res.Token,
				ExpiredAt: time.Now().Add(appconstant.ForgetPasswordExpiredDuration),
			}
		)

		urmock.On("FindByEmail", ctx, req.Email).Return(res1, nil)
		urmock.On("DeletePreviousResetPassword", ctx, int(res1.ID)).Return(nil)
		rtmock.On("Generate", 6).Return(res.Token, nil)
		urmock.On("CreateUserForgetPassword", ctx, mock.Anything).Return(res2, nil)

		lur, err := us.ForgetPassword(ctx, req)

		assert.Equal(res.Token, lur.Token)
		assert.Nil(err)
	})
}

func TestUserUsecase_ResetPassword(t *testing.T) {
	assert := assert.New(t)
	t.Run("should success call CheckPassword with right data", func(t *testing.T) {
		var (
			urmock = mocks.UserRepository{}
			phmock = mocks.PasswordHasher{}
			atmock = mocks.AuthTokenGenerator{}
			rtmock = mocks.RandomTokenGenerator{}
			usc    = usecase.UserUsecaseConfig{
				UserRepository:       &urmock,
				PasswordHasher:       &phmock,
				AuthTokenGenerator:   &atmock,
				RandomTokenGenerator: &rtmock,
			}
			us   = usecase.NewUserUsecase(usc)
			ctx  = context.Background()
			req  = dtousecase.ResetPasswordRequest{Email: "satria@mail.com", NewPassword: "newPassword", Token: "token1"}
			res1 = entity.UserResetPassword{ID: 1, Token: req.Token, UserId: 1}
			res2 = entity.User{ID: 1, Email: req.Email}
			res3 = "hashedPassword"
			req4 = entity.User{ID: res1.ID, Email: res2.Email, Password: res3}
			res4 = entity.User{ID: res1.ID, Email: res2.Email, Password: res3}
			req5 = entity.UserResetPassword{ID: 1, Token: req.Token, UserId: 1}
			res5 = entity.UserResetPassword{ID: 1, Token: req.Token, UserId: 1}
		)

		urmock.On("GetResetPasswordTokenByToken", ctx, req.Token).Return(res1, nil)
		urmock.On("FindByEmail", ctx, req.Email).Return(res2, nil)
		phmock.On("GenerateHashFromPassword", req.NewPassword).Return(res3, nil)
		urmock.On("UpdateUser", ctx, req4).Return(res4, nil)
		urmock.On("DeleteUsedResetPassword", ctx, req5).Return(res5, nil)

		err := us.ResetPassword(ctx, req)

		assert.Nil(err)
	})
}
