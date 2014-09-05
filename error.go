package client

import (
	"github.com/juju/errgo"
)

var (
	ErrUnexpectedResponse = errgo.New("Unexpected response from token service")

	Mask = errgo.MaskFunc()
)
