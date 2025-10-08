package discard

func (c *Client) GoogleSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/google", map[string]interface{}{"query": query}, nil)
}

func (c *Client) BingSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/bing", map[string]interface{}{"query": query}, nil)
}

func (c *Client) ImgurSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/imgur", map[string]interface{}{"query": query}, nil)
}

func (c *Client) TimeSearch(place string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/time", map[string]interface{}{"place": place}, nil)
}

func (c *Client) FlickerSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/flicker", map[string]interface{}{"query": query}, nil)
}

func (c *Client) ItuneSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/itunes", map[string]interface{}{"query": query}, nil)
}

func (c *Client) WattpadSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/wattpad", map[string]interface{}{"query": query}, nil)
}

func (c *Client) StickerSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/stickers", map[string]interface{}{"query": query}, nil)
}

func (c *Client) YouTubeSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/youtube2", map[string]interface{}{"query": query}, nil)
}

func (c *Client) TrackSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/youtube2", map[string]interface{}{"query": query}, nil)
}

func (c *Client) GifSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/klipy/gif", map[string]interface{}{"query": query}, nil)
}

func (c *Client) MemeSearch(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/klipy/meme", map[string]interface{}{"query": query}, nil)
}
