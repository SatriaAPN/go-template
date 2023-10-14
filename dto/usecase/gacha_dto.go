package dtousecase

import "github.com/shopspring/decimal"

type GetGachaRequest struct {
	UserId int
}

type GetGachaResponse struct {
	AvailableAttempts int
	GachaBoard        [][]int
}

type ChooseGachaRequest struct {
	UserId        int
	ChoosenNumber int
}

type ChooseGachaResponse struct {
	ChoosenNumber int
	Reward        decimal.Decimal
}
