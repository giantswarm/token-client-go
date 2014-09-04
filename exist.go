package client

import (
	apiSchemaPkg "github.com/catalyst-zero/api-schema"
)

func (this *Client) UserHasToken(userId string) (bool, error) {
	zeroVal := false

	res, err := this.get(this.endpointUrl("/token/" + userId + "/exist"))
	if err != nil {
		return zeroVal, Mask(err)
	}

	// Check token service response.
	var tokenExist bool
	if err := apiSchemaPkg.ParseData(&res.Body, &tokenExist); err != nil {
		return zeroVal, Mask(err)
	}

	return tokenExist, nil
}
