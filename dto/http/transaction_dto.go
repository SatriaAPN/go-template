package dtohttp

import (
	"time"

	"github.com/shopspring/decimal"
)

type ListOfTransactionRequest struct {
	UserId     int
	Sort       Sort
	Search     search
	Filter     DateFilter
	Pagination Pagination
}

type Sort struct {
	By    string `default:"date"`
	Order string `default:"descending"`
}

type search string

type DateFilter struct {
	From time.Time
	To   time.Time
}

type Pagination struct {
	Limit int `default:"10"`
	Page  int `default:"0"`
}

type ListOfTransactionResponseData struct {
	Type        string          `json:"list"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description"`
	Date        time.Time       `json:"date"`
	From        string          `json:"from"`
	To          string          `json:"to"`
}

type ListOfTransactionResponse struct {
	List       []ListOfTransactionResponseData `json:"list"`
	TotalPages int                             `json:"total_pages" default:"1"`
	TotalItems int                             `json:"total_items" default:"0"`
}

type CreateTopUpRequest struct {
	UserId         int
	Amount         decimal.Decimal `json:"amount" binding:"required"`
	SourceOfFundId int             `json:"source_of_fund_id" binding:"required"`
}

type CreateTopUpResponse struct {
	Date time.Time
}

type CreateTransferRequest struct {
	ReceiverWalletNumber string          `json:"receiver_wallet_number" binding:"required"`
	Amount               decimal.Decimal `json:"amount" binding:"required"`
	Description          string          `json:"description" binding:"required"`
}

type CreateTransferResponse struct {
	Date time.Time `json:"date"`
}
