package errorapp

import (
	"fmt"
	"go-template/share/config"

	"github.com/go-errors/errors"
)

func ErrorHandling(err error) error {
	switch e := err.(type) {
	case *errors.Error:
		return e
	default:
		return errors.New(e.Error())
	}
}

var ErrEmailAlreadyExist = errors.Errorf("email already exist")
var ErrEmailIsNotValid = errors.Errorf("email is not valid")
var ErrWrongPassword = errors.Errorf("wrong password")
var ErrEmailNotFound = errors.Errorf("email not found")
var ErrMinimumPasswordLength = errors.Errorf("minimum password length is " + fmt.Sprint(config.MinimumPasswordLength))
var ErrMaximumPasswordLength = errors.Errorf("maximum password length is " + fmt.Sprint(config.MaximumPasswordLength))
var ErrForgetPasswordTokenLength = errors.Errorf("token length is " + fmt.Sprint(config.ForgetPasswordTokenLength))
var ErrResetCodeNotFound = errors.Errorf("reset code not found")
