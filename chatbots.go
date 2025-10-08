package discard

func (c *Client) LlamaBot(text string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	return c.makeRequest("GET", "/api/bot/llama", params, nil)
}

func (c *Client) QwenBot(text string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	return c.makeRequest("GET", "/api/bot/qwen", params, nil)
}

func (c *Client) BaiduBot(text string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	return c.makeRequest("GET", "/api/bot/baidu", params, nil)
}

func (c *Client) GemmaBot(text string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	return c.makeRequest("GET", "/api/bot/gemma", params, nil)
}

func (c *Client) SparkBot(text string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	return c.makeRequest("GET", "/api/chat/spark", params, nil)
}

func (c *Client) QuarkBot(text string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	return c.makeRequest("GET", "/api/chat/quark", params, nil)
}

func (c *Client) GLMBot(text string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	return c.makeRequest("GET", "/api/chat/glm", params, nil)
}
