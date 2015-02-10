package client

func (this *Client) Remove(token string) error {
	_, err := this.delete(this.endpointUrl("/token/" + token))
	if err != nil {
		return Mask(err)
	}

	return nil
}

// RemoveAll will remove all tokens for the given user id
func (this *Client) RemoveAll(userId string) error {
	_, err := this.delete(this.endpointUrl("/user/" + userId))
	if err != nil {
		return Mask(err)
	}

	return nil
}
