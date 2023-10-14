package dtohttp

import "github.com/shopspring/decimal"

type GetGachaRequest struct {
	UserId int
}

type GetGachaResponse struct {
	AvailableAttempts int     `json:"available_attempts"`
	GachaBoard        [][]int `json:"gacha_board"`
}

type ChooseGachaRequest struct {
	ChoosenNumber int `json:"choosen_number"`
}

type ChooseGachaResponse struct {
	ChoosenNumber int             `json:"choosen_number"`
	Reward        decimal.Decimal `json:"reward"`
}
