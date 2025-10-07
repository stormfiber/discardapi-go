package discard

func (c *Client) SearchSpotify(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/spotify", map[string]interface{}{"query": query}, nil)
}

func (c *Client) DownloadSpotify(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/spotify", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Lyrics(song string) (interface{}, error) {
	return c.makeRequest("GET", "/api/music/lyrics", map[string]interface{}{"song": song}, nil)
}

func (c *Client) SearchSoundCloud(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/search/soundcloud", map[string]interface{}{"query": query}, nil)
}
