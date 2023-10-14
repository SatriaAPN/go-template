package dtohttp

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateUserRequest struct {
	Name     string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponse struct {
	Name  string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type ProfileUserResponse struct {
	Name          string          `json:"username"`
	Email         string          `json:"email"`
	WalletNumber  string          `json:"wallet_number"`
	WalletBalance decimal.Decimal `json:"wallet_balance"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type ForgetPasswordResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
	Token       string `json:"unique_reset_password_code"`
}
