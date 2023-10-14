package dtorepository

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateTopUpRequest struct {
	UserId         int
	Amount         decimal.Decimal
	SourceOfFundId int
}

type CreateTopUpResponse struct {
	Date time.Time
}

type CreateTransferRequest struct {
	UserId               int
	ReceiverWalletNumber string
	Amount               decimal.Decimal
	Description          string
}

type CreateTransferResponse struct {
	Date time.Time
}

type UserWalletDataResponse struct {
	WalletNumber  string
	WalletBalance decimal.Decimal
}

type GetWalletBalanceRequest struct {
	UserId int
}

type GetWalletBalanceResponse struct {
	UserId        int
	WalletBalance decimal.Decimal
}
