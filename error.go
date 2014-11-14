package client

import (
	"github.com/juju/errgo"
)

var (
	ErrTokenNotFound      = errgo.New("Token not found")
	ErrUnexpectedResponse = errgo.New("Unexpected response from token service")

	Mask = errgo.MaskFunc(IsErrTokenNotFound)
)

func IsErrTokenNotFound(err error) bool {
	return errgo.Cause(err) == ErrTokenNotFound
}
