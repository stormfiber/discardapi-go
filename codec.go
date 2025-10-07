package discard

func (c *Client) Base64(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base64", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Base32(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base32", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Binary(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/binary", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}
