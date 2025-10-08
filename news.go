package discard

func (c *Client) AlJazeeraEnglish() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/aljazeera", nil, nil)
}

func (c *Client) AlJazeeraArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/aljazeera/article", map[string]interface{}{"url": url}, nil)
}
