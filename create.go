package client

import (
	"encoding/json"
	"io"
	"strings"

	apiSchemaPkg "github.com/giantswarm/api-schema"
)

func (this *Client) Create(userId string) (Token, error) {
	zeroVal := Token{}

	// Create payload and hash password.
	reader, err := newCreateTokenPayload(userId)
	if err != nil {
		return zeroVal, Mask(err)
	}

	res, err := this.post(this.endpointUrl("/token/"), "application/json", reader)
	if err != nil {
		return zeroVal, Mask(err)
	}

	var token Token
	if err := apiSchemaPkg.ParseData(&res.Body, &token); err != nil {
		return zeroVal, Mask(err)
	}

	return token, nil
}

func newCreateTokenPayload(userId string) (io.Reader, error) {
	payload := map[string]string{
		"user_id": userId,
	}

	byteSlice, err := json.Marshal(payload)
	if err != nil {
		return nil, Mask(err)
	}

	return strings.NewReader(string(byteSlice)), nil
}
