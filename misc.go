package discard

import (
	"os"
)

// Compress handles text or URL compression requests
func (c *Client) Compress(t, text, url string) (interface{}, error) {
	params := map[string]interface{}{"type": t}
	if t == "text" && text != "" {
		params["text"] = text
	}
	if t == "url" && url != "" {
		params["url"] = url
	}
	return c.makeRequest("GET", "/api/compress", params, nil)
}

// Decompress handles text or URL decompression requests
func (c *Client) Decompress(t, data, url string) (interface{}, error) {
	params := map[string]interface{}{"type": t}
	if t == "text" && data != "" {
		params["data"] = data
	}
	if t == "url" && url != "" {
		params["url"] = url
	}
	return c.makeRequest("GET", "/api/decompress", params, nil)
}

// MathQuiz calls the math quiz API
func (c *Client) MathQuiz(difficulty, steps, allowNegative, numQuestions string) (interface{}, error) {
	params := map[string]interface{}{}
	if difficulty != "" {
		params["difficulty"] = difficulty
	}
	if steps != "" {
		params["steps"] = steps
	}
	if allowNegative != "" {
		params["allow_negative"] = allowNegative
	}
	if numQuestions != "" {
		params["num_questions"] = numQuestions
	}
	return c.makeRequest("GET", "/api/tools/math", params, nil)
}

// MorseText encodes or decodes morse text
func (c *Client) MorseText(text, mode string) (interface{}, error) {
	params := map[string]interface{}{"text": text, "mode": mode}
	return c.makeRequest("GET", "/api/tools/morse", params, nil)
}

// ReadQR reads QR code contents from an image
func (c *Client) ReadQR(imagePath string) (interface{}, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	files := map[string]io.Reader{"image": file}
	return c.MakeFormDataRequest("/api/tools/readqr", nil, files)
}

// TextToSpeech converts text into speech
func (c *Client) TextToSpeech(text, lang string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	if lang != "" {
		params["lang"] = lang
	}
	return c.makeRequest("GET", "/api/tools/tts", params, nil)
}

// WebsiteSEO fetches SEO details for a website
func (c *Client) WebsiteSEO(url string) (interface{}, error) {
	params := map[string]interface{}{"url": url}
	return c.makeRequest("GET", "/api/tools/seo", params, nil)
}

// StandardPing pings a website and returns details
func (c *Client) StandardPing(url, lang string) (interface{}, error) {
	params := map[string]interface{}{"url": url}
	if lang != "" {
		params["lang"] = lang
	}
	return c.makeRequest("GET", "/api/tools/ping", params, nil)
}

// ToASCII converts image to ASCII art
func (c *Client) ToASCII(imagePath string) (interface{}, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	files := map[string]io.Reader{"file": file}
	return c.MakeFormDataRequest("/api/tools/ascii", nil, files)
}

// RandomUUID generates a random UUID
func (c *Client) RandomUUID() (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/uuid", nil, nil)
}

// GenerateHash hashes input data with specified algorithm
func (c *Client) GenerateHash(data, algo string) (interface{}, error) {
	params := map[string]interface{}{"data": data, "algo": algo}
	return c.makeRequest("GET", "/api/tools/hash", params, nil)
}

// StrongPassword generates a strong random password
func (c *Client) StrongPassword() (interface{}, error) {
	return c.makeRequest("GET", "/api/tools/password", nil, nil)
}

// CompressFile uploads and compresses a file
func (c *Client) CompressFile(filePath, t string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	params := map[string]interface{}{"type": t}
	files := map[string]io.Reader{"file": file}
	return c.MakeFormDataRequest("/api/compress", params, files)
}

// DecompressFile uploads and decompresses a file
func (c *Client) DecompressFile(filePath, t string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	params := map[string]interface{}{"type": t}
	files := map[string]io.Reader{"file": file}
	return c.MakeFormDataRequest("/api/decompress", params, files)
}

// MarkdownToHTML converts markdown text to HTML
func (c *Client) MarkdownToHTML(markdown string) (interface{}, error) {
	body := map[string]interface{}{"markdown": markdown}
	return c.makeRequest("POST", "/api/tools/markdown", nil, body)
}

// ExifReader reads EXIF metadata from image
func (c *Client) ExifReader(imagePath string) (interface{}, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	files := map[string]io.Reader{"image": file}
	return c.MakeFormDataRequest("/api/tools/exif", nil, files)
}

// MinifyCSS minifies CSS input
func (c *Client) MinifyCSS(css string) (interface{}, error) {
	body := map[string]interface{}{"css": css}
	return c.makeRequest("POST", "/api/tools/minifycss", nil, body)
}

// JsonBeautify beautifies or minifies JSON input
func (c *Client) JsonBeautify(jsonStr string) (interface{}, error) {
	body := map[string]interface{}{"json": jsonStr}
	return c.makeRequest("POST", "/api/json/format", nil, body)
}
