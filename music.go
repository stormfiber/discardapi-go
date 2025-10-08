package discard

func (c *Client) SearchSpotify(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/spotify", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlSpotify(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/spotify", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Lyrics(song string) (interface{}, error) {
	return c.makeRequest("GET", "/api/music/lyrics", map[string]interface{}{"song": song}, nil)
}

func (c *Client) SearchSoundCloud(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/soundcloud", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlSoundCloud(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/soundcloud", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Ringtones(title string) (interface{}, error) {
	return c.makeRequest("GET", "/api/music/lyrics", map[string]interface{}{"title": title}, nil)
}

func (c *Client) SearchSound(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/sound", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlSound(id int) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/sound", map[string]interface{}{"id": id}, nil)
}

func (c *Client) SearchDeezer(track string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/deezer", map[string]interface{}{"track": track}, nil)
}

func (c *Client) dlDeezer(id int) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/deezer", map[string]interface{}{"id": id}, nil)
}
