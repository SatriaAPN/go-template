package errorhttp

import "github.com/go-errors/errors"

var (
	ErrUserNotAuthorized      = "you are not authorized"
	ErrSortQueryNotAvailable  = errors.Errorf("available sort query are only date/amount/to")
	ErrOrderQueryNotAvailable = errors.Errorf("available order query are only asc/desc")
)
