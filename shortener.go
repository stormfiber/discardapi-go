package discard

func (c *Client) ShortenIsgd(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/isgd", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenClck(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/clck", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenTiny(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/tiny", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenBitly(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/bitly", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Unshort(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/unshort", map[string]interface{}{"url": url}, nil)
}
