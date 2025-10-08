package Discard 

func (c *Client) Android1(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/search/android1", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlAndroid1(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/dl/android1", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ApkMirror(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/search/apkmirror", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlApkMirror(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/dl/apkmirror", map[string]interface{}{"url": url}, nil)
}

func (c *Client) ApkPure(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/search/apkpure", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlApkPure(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/dl/apkpure", map[string]interface{}{"url": url}, nil)
}

func (c *Client) PlayStore(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/search/playstore", map[string]interface{}{"query": query}, nil)
}

func (c *Client) dlPlayStore(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/dl/playstore", map[string]interface{}{"url": url}, nil)
}

func (c *Client) HappyMod(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/search/happymod", map[string]interface{}{"query": query}, nil)
}

func (c *Client) SfileMobi(query string) (interface{}, error) {
	return c.makeRequest("GET", "/api/apk/search/sfile", map[string]interface{}{"query": query}, nil)
}
