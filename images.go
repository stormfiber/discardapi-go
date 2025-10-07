package discard

func (c *Client) CatImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/cat", nil, nil)
}

func (c *Client) DogImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/dog", nil, nil)
}

func (c *Client) CoupleImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/couple", nil, nil)
}

func (c *Client) IslamicImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/islamic", nil, nil)
}

func (c *Client) ProgrammingImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/programming", nil, nil)
}
