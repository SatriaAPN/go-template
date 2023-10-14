package dtorepository

import (
	"time"

	"github.com/shopspring/decimal"
)

type GetUserAvailaibleAttemptsRequest struct {
	UserId int
}

type GetUserAvailableAttemptsResponse struct {
	AvailableAttempts int
}

type FindGachaRewardRequest struct {
	ChoosenNumber int
}

type FindGachaRewardResponse struct {
	Reward decimal.Decimal
}

type SaveGachaRewardRequest struct {
	UserId int
	Reward decimal.Decimal
}

type SaveGachaRewardResponse struct {
	UserId int
	Reward decimal.Decimal
	Date   time.Time
}

type SaveGachaZonkRequest struct {
	UserId int
}

type SaveGachaZonkResponse struct {
	UserId int
	Date   time.Time
}
