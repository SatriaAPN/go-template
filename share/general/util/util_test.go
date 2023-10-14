package util_test

import (
	dto "go-template/dto/general"
	dtohttp "go-template/dto/http"
	"go-template/share/general/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthTokenGenerator(t *testing.T) {
	assert := assert.New(t)
	t.Run("should success encode and decode token ", func(t *testing.T) {
		var (
			ud  = dto.AuthData{ID: 1}
			atg = util.GetAuthTokenGenerator()
		)

		token, err1 := atg.Encode(ud)

		ud2, err2 := atg.Decode(token)

		assert.NotNil(token)
		assert.Nil(err1)
		assert.EqualValues(ud, ud2)
		assert.Nil(err2)
	})

	t.Run("should success error when decoding unknown token ", func(t *testing.T) {
		var (
			atg   = util.GetAuthTokenGenerator()
			token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		)

		_, err := atg.Decode(token)

		assert.NotNil(err)
	})
}

func TestPasswordHasher(t *testing.T) {
	assert := assert.New(t)
	t.Run("should success hash and verify password", func(t *testing.T) {
		var (
			p  = "password"
			wp = "wrongPassword"

			ph = util.GetPasswordHasher()
		)

		hp, err1 := ph.GenerateHashFromPassword(p)

		match, err2 := ph.CompareHashAndPassword(p, hp)

		match2, err3 := ph.CompareHashAndPassword(wp, hp)

		assert.NotNil(hp)
		assert.Nil(err1)
		assert.Equal(true, match)
		assert.Nil(err2)
		assert.Equal(false, match2)
		assert.ErrorIs(err3, bcrypt.ErrMismatchedHashAndPassword)
	})
}

func TestDataValidator(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return true when email valid", func(t *testing.T) {
		var (
			dv    = util.NewDataValidator()
			email = "satria@mail.com"
		)

		b := dv.IsEmailValid(email)

		assert.Equal(true, b)
	})

	t.Run("should return false when email invalid", func(t *testing.T) {
		var (
			dv    = util.NewDataValidator()
			email = "satriamail.com"
		)

		b := dv.IsEmailValid(email)

		assert.Equal(false, b)
	})
}

// do I even need to test this? my question to me
func TestLogger(t *testing.T) {
	t.Run("test logging info and error", func(t *testing.T) {
		var (
			logger = util.GetLogger()
		)

		logger.Infof(dtohttp.NewHttpRequestLogging("/endpoint", "GET", "x-request-id", "request"))

		logger.Errorf(dtohttp.NewHttpRequestLogging("/endpoint", "GET", "x-request-id", "request"))
	})
}

func TestRandomTokenGenerator(t *testing.T) {
	assert := assert.New(t)
	t.Run("test logging info and error", func(t *testing.T) {
		var (
			rtg = util.GetRandomTokenGenerator()

			length = 6
		)

		token, err := rtg.Generate(length)

		assert.Nil(err)
		assert.Len(token, length)
	})
}
