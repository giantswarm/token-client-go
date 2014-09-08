package client

import (
	apiSchemaPkg "github.com/catalyst-zero/api-schema"
)

func (this *Client) Get(tokenId string) (Token, error) {
	zeroVal := Token{}

	res, err := this.get(this.endpointUrl("/token/" + tokenId))
	if err != nil {
		return zeroVal, Mask(err)
	}

	if ok, err := apiSchemaPkg.IsStatusResourceNotFound(&res.Body); err != nil {
		return zeroVal, Mask(err)
	} else if ok {
		return zeroVal, Mask(ErrTokenNotFound)
	}

	var token Token
	if err := apiSchemaPkg.ParseData(&res.Body, &token); err != nil {
		return zeroVal, Mask(err)
	}

	return token, nil
}
