package discard

func (c *Client) DadJoke() (interface{}, error) {
	return c.makeRequest("GET", "/api/joke/dad", nil, nil)
}

func (c *Client) ProgrammingJoke() (interface{}, error) {
	return c.makeRequest("GET", "/api/joke/programming", nil, nil)
}

func (c *Client) RandomJoke() (interface{}, error) {
	return c.makeRequest("GET", "/api/joke/random", nil, nil)
}

func (c *Client) DarkJoke() (interface{}, error) {
	return c.makeRequest("GET", "/api/joke/dark", nil, nil)
}
