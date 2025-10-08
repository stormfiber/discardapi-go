package discard

func (c *Client) CustomMeme(templateID, text1, text2, text3, text4, text5 string) (interface{}, error) {
	params := map[string]interface{}{"template_id": templateID, "text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	if text4 != "" {
		params["text4"] = text4
	}
	if text5 != "" {
		params["text5"] = text5
	}
	return c.makeRequest("GET", "/api/meme/custom", params, nil)
}

func (c *Client) TwoButtons(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/buttons", params, nil)
}

func (c *Client) YellingWoman(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/yelling", params, nil)
}

func (c *Client) SuccessKid(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/success", params, nil)
}

func (c *Client) PuppetMonkey(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/puppet", params, nil)
}

func (c *Client) ThinkingCouple(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/couple", params, nil)
}

func (c *Client) WinniePooh(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/pooh", params, nil)
}

func (c *Client) SquidGame(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/squid", params, nil)
}

func (c *Client) RockDriving(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/rock", params, nil)
}

func (c *Client) BlackGuy(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/disappointed", params, nil)
}

func (c *Client) DisasterGirl(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/disaster", params, nil)
}

func (c *Client) DrakeHotline(text1, text2 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	return c.makeRequest("GET", "/api/meme/drake", params, nil)
}

func (c *Client) ArgumentMeme(text1, text2, text3, text4, text5 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	if text4 != "" {
		params["text4"] = text4
	}
	if text5 != "" {
		params["text5"] = text5
	}
	return c.makeRequest("GET", "/api/meme/argument", params, nil)
}

func (c *Client) MaskReveal(text1, text2, text3, text4 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	if text4 != "" {
		params["text4"] = text4
	}
	return c.makeRequest("GET", "/api/meme/mask", params, nil)
}

func (c *Client) ExpandingBrain(text1, text2, text3, text4 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	if text4 != "" {
		params["text4"] = text4
	}
	return c.makeRequest("GET", "/api/meme/brain", params, nil)
}

func (c *Client) DrowningKid(text1, text2, text3, text4 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	if text4 != "" {
		params["text4"] = text4
	}
	return c.makeRequest("GET", "/api/meme/drowning", params, nil)
}

func (c *Client) DistractedBF(text1, text2, text3 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	return c.makeRequest("GET", "/api/meme/boyfriend", params, nil)
}

func (c *Client) LeftExit(text1, text2, text3 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	return c.makeRequest("GET", "/api/meme/exit", params, nil)
}

func (c *Client) AnakinPadme(text1, text2, text3 string) (interface{}, error) {
	params := map[string]interface{}{"text1": text1}
	if text2 != "" {
		params["text2"] = text2
	}
	if text3 != "" {
		params["text3"] = text3
	}
	return c.makeRequest("GET", "/api/meme/padme", params, nil)
}

func (c *Client) TrendingMemes() (interface{}, error) {
	return c.makeRequest("GET", "/api/meme/memes", nil, nil)
}
