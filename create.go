package client

import (
	"io"
  "strings"
  "encoding/json"
	"net/http"

	"github.com/juju/errgo"
)

func (this *Client) Create(userId string) (*http.Response, error) {
	// Create payload and hash password.
	reader, err := newCreateTokenPayload(userId)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	res, err := this.post(this.endpointUrl("/token/"), "application/json", reader)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return res, nil
}

func newCreateTokenPayload(userId string) (io.Reader, error) {
	payload := map[string]string{
		"user_id": userId,
	}

	byteSlice, err := json.Marshal(payload)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return strings.NewReader(string(byteSlice)), nil
}
