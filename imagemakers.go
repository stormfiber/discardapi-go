package discard

func (c *Client) QRCode(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/maker/qrcode", map[string]interface{}{"text": text}, nil)
}

func (c *Client) QRTag(text string, size, color, logo string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	if size != "" {
		params["size"] = size
	}
	if color != "" {
		params["color"] = color
	}
	if logo != "" {
		params["logo"] = logo
	}
	return c.makeRequest("GET", "/api/maker/qrtag", params, nil)
}

func (c *Client) TextToPic(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/maker/ttp", map[string]interface{}{"text": text}, nil)
}

func (c *Client) DesignFont(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/design/font", map[string]interface{}{"text": text}, nil)
}

func (c *Client) WebLogo(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/maker/weblogo", map[string]interface{}{"url": url}, nil)
}

func (c *Client) TextAvatar(text, shape string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	if shape != "" {
		params["shape"] = shape
	}
	return c.makeRequest("GET", "/api/maker/avatar", params, nil)
}

func (c *Client) CarbonImage(code, bg string) (interface{}, error) {
	params := map[string]interface{}{"code": code}
	if bg != "" {
		params["bg"] = bg
	}
	return c.makeRequest("GET", "/api/maker/carbon", params, nil)
}
