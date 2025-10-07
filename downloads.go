package discard

func (c *Client) DownloadFacebook(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/facebook", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DownloadGitClone(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/gitclone", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DownloadInstagram(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/instagram", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DownloadMediafire(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/mediafire", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DownloadPinterest(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/pinterest", map[string]interface{}{"text": text}, nil)
}

func (c *Client) DownloadTikTok(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/tiktok", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DownloadTwitter(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/twitter", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DownloadYouTube(url, format string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/youtube", map[string]interface{}{
		"url":    url,
		"format": format,
	}, nil)
}

func (c *Client) DownloadSnapChat(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/snapchat", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DownloadReddit(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/reddit", map[string]interface{}{"url": url}, nil)
}
