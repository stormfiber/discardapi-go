package discard

import (
	"io"
)

func (c *Client) Catbox(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/catbox", nil, map[string]io.Reader{"file": file})
}

func (c *Client) Gofile(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/gofile", nil, map[string]io.Reader{"file": file})
}

func (c *Client) Gyazo(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/gyazo", nil, map[string]io.Reader{"file": file})
}

func (c *Client) Hastebin(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/hastebin", nil, map[string]io.Reader{"file": file})
}

func (c *Client) Pastebin(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/pastebin", nil, map[string]io.Reader{"file": file})
}

func (c *Client) ImgBB(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/imgbb", nil, map[string]io.Reader{"file": file})
}

func (c *Client) PixelDrain(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/pixeldrain", nil, map[string]io.Reader{"file": file})
}

func (c *Client) Uguuse(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/uguuse", nil, map[string]io.Reader{"file": file})
}

func (c *Client) Map360(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/upload/360", map[string]interface{}{"url": url}, nil)
}

func (c *Client) XianImage(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/upload/xian", map[string]interface{}{"url": url}, nil)
}
