package discard

func (c *Client) CatImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/cat", nil, nil)
}

func (c *Client) DogImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/dog", nil, nil)
}

func (c *Client) CoupleImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/couple", nil, nil)
}

func (c *Client) PizzaImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/pizza", nil, nil)
}

func (c *Client) BurgerImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/burger", nil, nil)
}

func (c *Client) PastaImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/pasta", nil, nil)
}

func (c *Client) DosaImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/dosa", nil, nil)
}

func (c *Client) BiryaniImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/biryani", nil, nil)
}

func (c *Client) IslamicImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/islamic", nil, nil)
}

func (c *Client) ProgrammingImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/programming", nil, nil)
}

func (c *Client) TechImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/tech", nil, nil)
}

func (c *Client) GameImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/game", nil, nil)
}

func (c *Client) MountainImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/mountain", nil, nil)
}

func (c *Client) CyberImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/cyberspace", nil, nil)
}

func (c *Client) PcImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/wallpc", nil, nil)
}

func (c *Client) MessiImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/messi", nil, nil)
}

func (c *Client) RonaldoImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/ronaldo", nil, nil)
}

func (c *Client) CoffeeImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/coffee", nil, nil)
}

func (c *Client) FoxImage() (interface{}, error) {
	return c.makeRequest("GET", "/api/img/fox", nil, nil)
}
