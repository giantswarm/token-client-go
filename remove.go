package client

import (
	"io"
	"net/http"

	"github.com/juju/errgo"
)

func (this *Client) Remove(token string) (*http.Response, error) {
	res, err := this.delete(this.endpointUrl("/token/" + token))
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return res, nil
}
