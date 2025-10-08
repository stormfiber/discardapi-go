package discard

func (c *Client) PickupLine() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/pickup", nil, nil)
}

func (c *Client) RandomQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quotes/random", nil, nil)
}

func (c *Client) CommitMsg() (interface{}, error) {
	return c.makeRequest("GET", "/api/commit/message", nil, nil)
}

func (c *Client) StrangerQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/stranger", nil, nil)
}

func (c *Client) MotivQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/motiv", nil, nil)
}

func (c *Client) StoicQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/stoic", nil, nil)
}

func (c *Client) LuciferQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/lucifer", nil, nil)
}

func (c *Client) IslamicQuote() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/islamic", nil, nil)
}

func (c *Client) TechTip() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/techtips", nil, nil)
}

func (c *Client) CodingTip() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/coding", nil, nil)
}

func (c *Client) WhyQuiz() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/why", nil, nil)
}

func (c *Client) FunFact() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/funfacts", nil, nil)
}

func (c *Client) WYRather() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/wyr", nil, nil)
}

func (c *Client) LifeHack() (interface{}, error) {
	return c.makeRequest("GET", "/api/quote/lifehacks", nil, nil)
}
