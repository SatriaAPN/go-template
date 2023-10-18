package util

import (
	"encoding/json"
	"fmt"
	dto "go-template/dto/general"
	"go-template/share/general/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthTokenGenerator interface {
	Encode(data dto.AuthData) (string, error)
	Decode(token string) (dto.AuthData, error)
}

type authTokenGenerator struct{}

var authTokenGeneratorInstance *authTokenGenerator

func GetAuthTokenGenerator() AuthTokenGenerator {
	if authTokenGeneratorInstance == nil {
		authTokenGeneratorInstance = newAuthTokenGenerator()
	}

	return authTokenGeneratorInstance
}

func newAuthTokenGenerator() *authTokenGenerator {
	return &authTokenGenerator{}
}

func (atg *authTokenGenerator) Encode(ad dto.AuthData) (string, error) {
	adByte, err := json.Marshal(ad)

	if err != nil {
		return "", err
	}

	adString := string(adByte)

	claims := jwt.MapClaims{
		"iat": time.Now().UnixMilli(),
		"iss": config.ApplicationName(),
		"exp": jwt.NewNumericDate(time.Now().Add(time.Duration(config.LoginExpirationDuration))),
		"aud": adString,
	}

	token := jwt.NewWithClaims(
		config.JwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(config.JwtSignatureKey())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (atg *authTokenGenerator) Decode(tokenString string) (dto.AuthData, error) {
	ad := dto.AuthData{}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != config.JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}

		return config.JwtSignatureKey(), nil
	})

	if err != nil {
		return ad, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ad, err
	}

	audience, err := claims.GetAudience()

	if err != nil {
		return ad, err
	}

	adByte := []byte(audience[0])

	err = json.Unmarshal(adByte, &ad)

	if err != nil {
		return ad, err
	}

	return ad, nil
}
