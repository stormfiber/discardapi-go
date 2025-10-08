package discard

func (c *Client) DeadPool(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/deadpool", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) WolfLogo(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/wolf", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) FootballShirt(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/shirt", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) PencilSketch(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/sketch", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) Avengers(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/avengers", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) ThorLogo(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/thor", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) Captain(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/captain", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) MascotLogo(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/mascot", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) FootballLogo(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/football", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) SteelText(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/steel", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) RoyalText(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/royal", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Comic3D(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/comic", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Gaming(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/game", map[string]interface{}{"text": text}, nil)
}

func (c *Client) SandText(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/sand", map[string]interface{}{"text": text}, nil)
}

func (c *Client) JewelEffect(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/jewel", map[string]interface{}{"text": text}, nil)
}

func (c *Client) GlitchText(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/glitch", map[string]interface{}{"text": text}, nil)
}

func (c *Client) MAvatar(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ephoto/mavatar", map[string]interface{}{"text": text}, nil)
}
