package client

import (
	"net/http"
)

func (this *Client) Remove(token string) (*http.Response, error) {
	res, err := this.delete(this.endpointUrl("/token/" + token))
	if err != nil {
		return nil, Mask(err)
	}

	return res, nil
}

// RemoveAll will remove all tokens for the given user id
func (this *Client) RemoveAll(userId string) (*http.Response, error) {
	res, err := this.delete(this.endpointUrl("/user/" + userId))
	if err != nil {
		return nil, Mask(err)
	}

	return res, nil
}
