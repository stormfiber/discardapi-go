package discard

func (c *Client) BankLogo(domain string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/banklogo", map[string]interface{}{"domain": domain}, nil)
}

func (c *Client) DetectLang(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/detect-lang", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Dictionary(word string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/dictionary", map[string]interface{}{"word": word}, nil)
}

func (c *Client) Dictionary2(word string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/dict", map[string]interface{}{"word": word}, nil)
}

func (c *Client) Mathematics(expr string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/math", map[string]interface{}{"expr": expr}, nil)
}

func (c *Client) Screenshot(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/ssweb", map[string]interface{}{"url": url}, nil)
}

func (c *Client) SsPc(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/sspc", map[string]interface{}{"url": url}, nil)
}

func (c *Client) SsMobile(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/ssmobile", map[string]interface{}{"url": url}, nil)
}

func (c *Client) StyleText(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/styletext", map[string]interface{}{"text": text}, nil)
}

func (c *Client) ReverseText(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/reverse", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Translate(text, to string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/translate", map[string]interface{}{
		"text": text,
		"to":   to,
	}, nil)
}

func (c *Client) Translate2(text, lang string) (interface{}, error) {
	return c.makeRequest("GET", "/api/go/translate", map[string]interface{}{
		"text": text,
		"lang": lang,
	}, nil)
}

func (c *Client) SimplePing(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/simple/ping", map[string]interface{}{"url": url}, nil)
}

func (c *Client) Counter(count int) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/count", map[string]interface{}{"count": count}, nil)
}

func (c *Client) Whois(domain string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/whois", map[string]interface{}{"domain": domain}, nil)
}

func (c *Client) HandWriting(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/handwrite", map[string]interface{}{"text": text}, nil)
}

func (c *Client) TextStats(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/string", map[string]interface{}{"text": text}, nil)
}

func (c *Client) TextStats(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/string", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Wordcount(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/word/count", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Converter(from, to string, value int) (interface{}, error) {
	return c.makeRequest("GET", "/api/convert/unit", map[string]interface{}{
		"from":  from,
		"to":    to,
		"value": value,
	}, nil)
}
