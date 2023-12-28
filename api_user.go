package filen

func (c *APIClient) AuthInfo() (*User, error) {
	
	_, err := c.postJSONWithBody(
		AuthInfoEndpoint,
		&user,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}