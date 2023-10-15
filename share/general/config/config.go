package config

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func InitEnvReader() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AddConfigPath("path")
}

func getEnvValue(key string) string {
	a := viper.Get(key)

	return a.(string)
}

var (
	DbHost = func() string {
		return getEnvValue("DB_HOST")
	}

	DbUser = func() string {
		return getEnvValue("DB_USER")
	}

	DbPassword = func() string {
		return getEnvValue("DB_PASSWORD")
	}

	DbName = func() string {
		return getEnvValue("DB_NAME")
	}

	DbPort = func() string {
		return getEnvValue("DB_PORT")
	}

	DbSslMode = func() string {
		return getEnvValue("DB_SSLMODE")
	}

	DbTimezone = func() string {
		return getEnvValue("DB_TIMEZONE")
	}

	JwtSigningMethod = jwt.SigningMethodHS256

	JwtSignatureKey = func() []byte {
		sk := []byte(getEnvValue("JWT_SIGNATURE_KEY"))
		return []byte(sk)

	}

	ApplicationName = func() string {
		return getEnvValue("APPLICATION_NAME")
	}

	LoginExpirationDuration = 1 * time.Hour

	BcryptCost = func() int {
		costString := getEnvValue("BCRYPT_COST")
		cost, _ := strconv.Atoi(costString)

		return cost
	}

	HttpRequestTimeoutSeconds = func() int {
		costString := getEnvValue("HTTP_REQUEST_TIMEOUT_SECONDS")
		cost, _ := strconv.Atoi(costString)
		return cost
	}

	WalletNumberStart                = 4200000000000
	MinimumPasswordLength            = 6
	MaximumPasswordLength            = 20
	ForgetPasswordTokenLength        = 6
	MinimumTopUpAmount               = 50000
	MaximumTopUpAmount               = 10000000
	MinimumTransferAmount            = 1000
	MaximumTransferAmount            = 50000000
	MaximumTransferDescriptionLength = 35
	GachaBoard                       = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	GachaBoardMinimumChoose       = 1
	GachaBoardMaximumChoose       = 9
	ForgetPasswordExpiredDuration = 15 * time.Minute
	GachaRewardLevel1             = 0
	GachaRewardLevel2             = 10000
	GachaRewardLevel3             = 20000
	GachaRewardLevel4             = 50000
	GachaRewardLevel5             = 100000
	GachaRewardLevel6             = 150000
	GachaRewardLevel7             = 200000
	GachaRewardLevel8             = 250000
	GachaRewardLevel9             = 300000
)
