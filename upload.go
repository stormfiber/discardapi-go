package discard

import (
	"io"
)

func (c *Client) UpCatbox(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/catbox", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpGofile(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/gofile", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpGyazo(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/gyazo", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpHastebin(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/hastebin", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpPastebin(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/pastebin", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpImgBB(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/imgbb", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpPixelDrain(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/pixeldrain", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpUguuse(file io.Reader) (interface{}, error) {
	return c.MakeFormDataRequest("/api/uguuse", nil, map[string]io.Reader{"file": file})
}

func (c *Client) UpMap360(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/upload/360", map[string]interface{}{"url": url}, nil)
}

func (c *Client) UpXianImage(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/upload/xian", map[string]interface{}{"url": url}, nil)
}
