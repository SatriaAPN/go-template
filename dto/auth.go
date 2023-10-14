package dto

type AuthData struct {
	ID uint `json:"id"`
}

type AuthDataType string

var AuthDataKey AuthDataType = "authdDataKey"
