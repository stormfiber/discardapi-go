package discard

func (c *Client) RandomQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quotes/random", nil, nil)
}

func (c *Client) MotivQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/motiv", nil, nil)
}

func (c *Client) IslamicQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/islamic", nil, nil)
}

func (c *Client) TechTip() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/techtips", nil, nil)
}
