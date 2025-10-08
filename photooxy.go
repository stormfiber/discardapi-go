package discard

func (c *Client) PubgBanner(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/pubg", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) BattleField(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/battle4", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) TikTokEffect(text1, text2 string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/tiktok", map[string]interface{}{
		"text1":  text1,
		"text2":  text2,
	}, nil)
}

func (c *Client) NeoEffect(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/neon", map[string]interface{}{"neon": neon}, nil)
}

func (c *Client) Warface(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/warface", map[string]interface{}{"neon": neon}, nil)
}

func (c *Client) Warface2(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/warface2", map[string]interface{}{"neon": neon}, nil)
}

func (c *Client) League(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/league", map[string]interface{}{"neon": neon}, nil)
}

func (c *Client) DarkMetal(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/photo/darkmetal", map[string]interface{}{"neon": neon}, nil)
}
