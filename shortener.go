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

func (c *Client) ShortenTinycc(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/tinycc", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenReurl(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/reurl", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenSurl(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/surl", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenVurl(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/vurl", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenVgd(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/vgd", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShortenClean(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/clean", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Unshort(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/short/unshort", map[string]interface{}{"url": url}, nil)
}
