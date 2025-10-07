package discard

func (c *Client) GeminiPro(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/gemini/pro", map[string]interface{}{"text": text}, nil)
}

func (c *Client) GeminiFlash(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/gemini/flash", map[string]interface{}{"text": text}, nil)
}

func (c *Client) GoogleGemma(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/gemini/gemma", map[string]interface{}{"text": text}, nil)
}

func (c *Client) GeminiEmbed(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/gemini/embed", map[string]interface{}{"text": text}, nil)
}

func (c *Client) LlamaAI(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ai/llama", map[string]interface{}{"text": text}, nil)
}

func (c *Client) Mythomax(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ai/mythomax", map[string]interface{}{"text": text}, nil)
}

func (c *Client) MistralAI(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ai/mistral", map[string]interface{}{"text": text}, nil)
}

func (c *Client) QwenCoder(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ai/qwen", map[string]interface{}{"text": text}, nil)
}

func (c *Client) KimiAI(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ai/kimi", map[string]interface{}{"text": text}, nil)
}

func (c *Client) GemmaAI(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/ai/gemma", map[string]interface{}{"text": text}, nil)
}

func (c *Client) FluxSchnell(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/imagen/schnell", map[string]interface{}{"text": text}, nil)
}

func (c *Client) FluxDev(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/imagen/flux", map[string]interface{}{"text": text}, nil)
}

func (c *Client) StableDiffusion(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/imagen/diffusion", map[string]interface{}{"text": text}, nil)
}

func (c *Client) BlackForest(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/imagen/sdxlb", map[string]interface{}{"text": text}, nil)
}

func (c *Client) DallE(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/imagen/dalle", map[string]interface{}{"text": text}, nil)
}
