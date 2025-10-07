package discard

func (c *Client) FakeUser() (interface{}, error) {
	return c.makeRequest("GET", "/api/fake/user", nil, nil)
}

func (c *Client) FakeUsers(quantity int, locale, gender string, seed int) (interface{}, error) {
	params := map[string]interface{}{}
	if quantity > 0 {
		params["_quantity"] = quantity
	}
	if locale != "" {
		params["_locale"] = locale
	}
	if gender != "" {
		params["_gender"] = gender
	}
	if seed > 0 {
		params["_seed"] = seed
	}
	return c.makeRequest("GET", "/api/fake/users", params, nil)
}

func (c *Client) FakeAddresses(quantity int, locale, countryCode string, seed int) (interface{}, error) {
	params := map[string]interface{}{}
	if quantity > 0 {
		params["_quantity"] = quantity
	}
	if locale != "" {
		params["_locale"] = locale
	}
	if countryCode != "" {
		params["_country_code"] = countryCode
	}
	if seed > 0 {
		params["_seed"] = seed
	}
	return c.makeRequest("GET", "/api/fake/addresses", params, nil)
}

func (c *Client) FakeCompanies(quantity int, locale string, seed int) (interface{}, error) {
	params := map[string]interface{}{}
	if quantity > 0 {
		params["_quantity"] = quantity
	}
	if locale != "" {
		params["_locale"] = locale
	}
	if seed > 0 {
		params["_seed"] = seed
	}
	return c.makeRequest("GET", "/api/fake/companies", params, nil)
}
