package discard 

func (c *Client) StalkGenshin(id int) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/genshin", map[string]interface{}{"id": id}, nil)
}

func (c *Client) StalkGithub(username string) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/github", map[string]interface{}{"username": username}, nil)
}

func (c *Client) StalkInsta(username string) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/instagram", map[string]interface{}{"username": username}, nil)
}

func (c *Client) StalkPinterest(username string) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/pinterest", map[string]interface{}{"username": username}, nil)
}

func (c *Client) StalkThreads(username string) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/threads", map[string]interface{}{"username": username}, nil)
}

func (c *Client) StalkTwitter(username string) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/twitter", map[string]interface{}{"username": username}, nil)
}

func (c *Client) StalkTikTok(username string) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/tiktok", map[string]interface{}{"username": username}, nil)
}

func (c *Client) StalkTelegram(username string) (interface{}, error) {
	return c.makeRequest("GET", "/api/stalk/telegram", map[string]interface{}{"username": username}, nil)
}
