package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type ListOfTransactionRequest struct {
	UserId     int
	Sort       Sort
	Search     string
	Filter     DateFilter
	Pagination Pagination
}

type Sort struct {
	By    string `default:"date"`
	Order string `default:"descending"`
}

type DateFilter struct {
	From time.Time
	To   time.Time
}

type Pagination struct {
	Limit int `default:"10"`
	Page  int `default:"0"`
}

type ListOfTransactionResponseData struct {
	Type        string
	Amount      decimal.Decimal
	Description string
	Date        time.Time
	From        string
	To          string
}

type ListOfTransactionResponse struct {
	List       []ListOfTransactionResponseData
	TotalPages int
	TotalItems int
}

type ListOfTransactionCount struct {
	TotalPages int
	TotalItems int
}

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
