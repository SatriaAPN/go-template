package config

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = func() []byte {
	sk := []byte(os.Getenv("JWT_SIGNATURE_KEY"))
	return []byte(sk)
}
var ApplicationName = func() string {
	return os.Getenv("APPLICATION_NAME")
}
var LoginExpirationDuration = 1 * time.Hour
var BcryptCost = func() int {
	costString := os.Getenv("BCRYPT_COST")
	cost, _ := strconv.Atoi(costString)

	return cost
}
var HttpRequestTimeoutSeconds = func() int {
	costString := os.Getenv("HTTP_REQUEST_TIMEOUT_SECONDS")
	cost, _ := strconv.Atoi(costString)
	return cost
}
var WalletNumberStart = 4200000000000
var MinimumPasswordLength = 6
var MaximumPasswordLength = 20
var ForgetPasswordTokenLength = 6
var MinimumTopUpAmount = 50000
var MaximumTopUpAmount = 10000000
var MinimumTransferAmount = 1000
var MaximumTransferAmount = 50000000
var MaximumTransferDescriptionLength = 35
var GachaBoard = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}
var GachaBoardMinimumChoose = 1
var GachaBoardMaximumChoose = 9
var ForgetPasswordExpiredMinutes = 15 * time.Minute
var GachaRewardLevel1 = 0
var GachaRewardLevel2 = 10000
var GachaRewardLevel3 = 20000
var GachaRewardLevel4 = 50000
var GachaRewardLevel5 = 100000
var GachaRewardLevel6 = 150000
var GachaRewardLevel7 = 200000
var GachaRewardLevel8 = 250000
var GachaRewardLevel9 = 300000
