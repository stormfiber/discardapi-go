package discard

func (c *Client) Facebook(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/facebook", map[string]interface{}{"url": url}, nil)
}

func (c *Client) GitClone(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/gitclone", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Instagram(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/instagram", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Mediafire(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/mediafire", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Pinterest(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/pinterest", map[string]interface{}{"text": text}, nil)
}

func (c *Client) TikTok(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/tiktok", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Twitter(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/twitter", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Likee(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/likee", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Threads(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/threads", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Twitch(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/twitch", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Wallpaper(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/wallcraft", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Wallpaper2(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/wallbest", map[string]interface{}{"text": text}, nil)
}

func (c *Client) BaiduImg(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/img/baidu", map[string]interface{}{"query": query}, nil)
}

func (c *Client) Wikimedia(title string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/wikimedia", map[string]interface{}{"title": title}, nil)
}

func (c *Client) YouTube(url, format string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/youtube", map[string]interface{}{
		"url":    url,
		"format": format,
	}, nil)
}

func (c *Client) SnapChat(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/snapchat", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Linkedin(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/linkedin", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ShareChat(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/sharechat", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Snackvideo(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/snack", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Pinterest2(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/pinterest", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Reddit(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/reddit", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Videezy(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/videezy", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Vidsplay(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/vidsplay", map[string]interface{}{"url": url}, nil)
}

func (c *Client) IMDbVideo(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/imdb", map[string]interface{}{"url": url}, nil)
}

func (c *Client) IFunny(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/ifunny", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Getty(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/getty", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Pastebin(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/pastebin", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Tenor(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/tenor", map[string]interface{}{"query": query}, nil)
}

func (c *Client) Odysee(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/odysee", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Alamy(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/alamy", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Istock(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/istock", map[string]interface{}{"url": url}, nil)
}
