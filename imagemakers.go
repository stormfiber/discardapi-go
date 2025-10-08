package discard

import "io"

func (c *Client) QRCode(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/maker/qrcode", map[string]interface{}{"text": text}, nil)
}

func (c *Client) QRTag(text, size, color, logo string) (interface{}, error) {
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

func (c *Client) CaptchaImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/maker/captcha", map[string]interface{}{}, nil)
}

func (c *Client) QRCustom(text, size, color string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	if size != "" {
		params["size"] = size
	}
	if color != "" {
		params["color"] = color
	}
	return c.makeRequest("GET", "/api/maker/customqr", params, nil)
}

func (c *Client) TextAvatar(text, shape string) (interface{}, error) {
	params := map[string]interface{}{"text": text}
	if shape != "" {
		params["shape"] = shape
	}
	return c.makeRequest("GET", "/api/maker/avatar", params, nil)
}

func (c *Client) WebLogo(url string) (interface{}, error) {
	return c.makeRequest("GET", "/api/maker/weblogo", map[string]interface{}{"url": url}, nil)
}

func (c *Client) WhoWins(url1, url2 string) (interface{}, error) {
	params := map[string]interface{}{"url1": url1, "url2": url2}
	return c.makeRequest("GET", "/api/maker/whowin", params, nil)
}

func (c *Client) QuotedLyo(text, name, profile, color string) (interface{}, error) {
	params := map[string]interface{}{
		"text":    text,
		"name":    name,
		"profile": profile,
	}
	if color != "" {
		params["color"] = color
	}
	return c.makeRequest("GET", "/api/maker/quoted", params, nil)
}

func (c *Client) QRPro(text, size, color, logo, caption string) (interface{}, error) {
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
	if caption != "" {
		params["caption"] = caption
	}
	return c.makeRequest("GET", "/api/qr/pro", params, nil)
}

func (c *Client) Img2Base64(file io.Reader) (interface{}, error) {
	files := map[string]io.Reader{"file": file}
	return c.MakeFormDataRequest("/api/img2base64", nil, files)
}

func (c *Client) Base642Img(data string) (interface{}, error) {
	return c.makeRequest("GET", "/api/img2base64", map[string]interface{}{"data": data}, nil)
}

func (c *Client) Barcode128(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/barcode/code", map[string]interface{}{"text": text}, nil)
}

func (c *Client) BarcodeEAN(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/barcode/ean", map[string]interface{}{"text": text}, nil)
}

func (c *Client) BarcodeQR(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/barcode/qr", map[string]interface{}{"text": text}, nil)
}

func (c *Client) EmojiMosaic(file io.Reader, width, palette, format string) (interface{}, error) {
	params := map[string]interface{}{}
	if width != "" {
		params["width"] = width
	}
	if palette != "" {
		params["palette"] = palette
	}
	if format != "" {
		params["format"] = format
	}
	files := map[string]io.Reader{"file": file}
	return c.MakeFormDataRequest("/api/emoji/mosaic", params, files)
}

func (c *Client) EmojiTranslate(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/emoji/translate", map[string]interface{}{"text": text}, nil)
}

func (c *Client) EmojiReplace(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/emoji/replace", map[string]interface{}{"text": text}, nil)
}

func (c *Client) EmojiMirror(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/emoji/mirror", map[string]interface{}{"text": text}, nil)
}

func (c *Client) EmojiRainbow(text string) (interface{}, error) {
	return c.makeRequest("GET", "/api/emoji/rainbow", map[string]interface{}{"text": text}, nil)
}

func (c *Client) EmojiMix(e1, e2 string) (interface{}, error) {
	params := map[string]interface{}{"e1": e1, "e2": e2}
	return c.makeRequest("GET", "/api/emoji/mix", params, nil)
}

func (c *Client) CarbonImage(code, bg string) (interface{}, error) {
	params := map[string]interface{}{"code": code}
	if bg != "" {
		params["bg"] = bg
	}
	return c.makeRequest("GET", "/api/maker/carbon", params, nil)
}

func (c *Client) WelcomeImage(background, avatar, text1, text2, text3 string) (interface{}, error) {
	params := map[string]interface{}{
		"background": background,
		"avatar":     avatar,
		"text1":      text1,
		"text2":      text2,
		"text3":      text3,
	}
	return c.makeRequest("GET", "/api/maker/welcome", params, nil)
}
