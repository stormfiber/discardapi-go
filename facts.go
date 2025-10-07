package discard

func (c *Client) DateFact(month, day int) (interface{}, error) {
	if month > 0 && day > 0 {
		return c.makeRequest("GET", "/api/fact/date", map[string]interface{}{
			"month": month,
			"day":   day,
		}, nil)
	}
	return c.makeRequest("GET", "/api/date/fact", nil, nil)
}

func (c *Client) YearFact(year int) (interface{}, error) {
	if year > 0 {
		return c.makeRequest("GET", "/api/fact/year", map[string]interface{}{"year": year}, nil)
	}
	return c.makeRequest("GET", "/api/year/fact", nil, nil)
}

func (c *Client) MathFact(number int) (interface{}, error) {
	if number > 0 {
		return c.makeRequest("GET", "/api/fact/math", map[string]interface{}{"number": number}, nil)
	}
	return c.makeRequest("GET", "/api/math/fact", nil, nil)
}

func (c *Client) UselessFact() (interface{}, error) {
	return c.makeRequest("GET", "/api/fact/useless", nil, nil)
}
