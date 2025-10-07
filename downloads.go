package discard

func (c *Client) dlFacebook(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/facebook", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlGitClone(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/gitclone", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlInstagram(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/instagram", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlMediafire(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/mediafire", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlPinterest(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/pinterest", map[string]interface{}{"text": text}, nil)
}

func (c *Client) dlTikTok(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/tiktok", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlTwitter(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/twitter", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlLikee(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/likee", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlThreads(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/threads", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlTwitch(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/twitch", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlWallpaper(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/wallcraft", map[string]interface{}{"text": text}, nil)
}

func (c *Client) dlWallpaper2(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/wallbest", map[string]interface{}{"text": text}, nil)
}

func (c *Client) dlBaiduImg(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/img/baidu", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlWikimedia(title string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/wikimedia", map[string]interface{}{"title": title}, nil)
}

func (c *Client) dlYouTube(url, format string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/youtube", map[string]interface{}{
		"url":    url,
		"format": format,
	}, nil)
}

func (c *Client) dlSnapChat(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/snapchat", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlLinkedin(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/linkedin", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlShareChat(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/sharechat", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlSnackvideo(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/snack", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Pinterestvideo(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/pinterest", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlReddit(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/reddit", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlVideezy(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/videezy", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlVidsplay(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/vidsplay", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlIMDb(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/imdb", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlIFunny(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/ifunny", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlGetty(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/getty", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlPastebin(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/pastebin", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlTenor(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/tenor", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlOdysee(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/odysee", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlAlamy(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/alamy", map[string]interface{}{"url": url}, nil)
}

func (c *Client) dlIstock(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/istock", map[string]interface{}{"url": url}, nil)
}
