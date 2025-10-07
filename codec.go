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

func (c *Client) Base16(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base16", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Base36(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base36", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Base45(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base45", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Base58(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base58", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Base62(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base62", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Base85(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base85", map[string]interface{}{
		"mode": mode,
		"data": data,
	}, nil)
}

func (c *Client) Base91(mode, data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/base91", map[string]interface{}{
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

func (c *Client) Brainfuck(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/brainfuck", map[string]interface{}{
		"text": text,
	}, nil)
}

func (c *Client) Interpreter(code string) (interface{}, error) {
	return c.makeRequest("GET", "/api/interpreter", map[string]interface{}{
		"code": code,
	}, nil)
}
