package discard

func (c *Client) AlJazeeraEnglish() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/aljazeera", nil, nil)
}

func (c *Client) AlJazeeraArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/aljazeera/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) AlJazeeraArabic() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/aljazeera/ar", nil, nil)
}

func (c *Client) ArabicArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/aljazeera/article/ar", map[string]interface{}{"url": url}, nil)
}

func (c *Client) TRTWorld() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/trt", nil, nil)
}

func (c *Client) TRTArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/trt/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) TRTAfrika() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/trt/af", nil, nil)
}

func (c *Client) AfrikaArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/trt/article/af", map[string]interface{}{"url": url}, nil)
}

func (c *Client) SkyNews() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/sky", nil, nil)
}

func (c *Client) SkyArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/sky/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) SkySports(sport string) (interface{}, error) {
	return c.makeRequest("GET", "/api/news/skysports", map[string]interface{}{"sport": sport}, nil)
}

func (c *Client) SportsArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/skysports/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) CGTNWorld() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/cgtn", nil, nil)
}

func (c *Client) CGTNArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/cgtn/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) CGTNTech() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/cgtn/tech", nil, nil)
}

func (c *Client) TechArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/cgtn/tech/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) DawnNews() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/dawn", nil, nil)
}

func (c *Client) DawnArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dawn/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) CNNNews() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/cnn", nil, nil)
}

func (c *Client) CNNArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/cnn/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) TheGuardian() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/guardian", nil, nil)
}

func (c *Client) GuardianArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/guardian/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) GeoNews() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/geo/en", nil, nil)
}

func (c *Client) GeoArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/geo/en/article", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ExpressNews() (interface{}, error) {
	return c.makeRequest("GET", "/api/news/express", nil, nil)
}

func (c *Client) ExpressArticle(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/express/article", map[string]interface{}{"url": url}, nil)
}

