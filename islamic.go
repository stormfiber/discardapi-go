package discard

func (c *Client) QuranSurah(surah string) (interface{}, error) {
	return c.makeRequest("GET", "/api/dl/surah", map[string]interface{}{"surah": surah}, nil)
}

func (c *Client) Hadith(book, number string) (interface{}, error) {
	return c.makeRequest("GET", "/api/get/hadith", map[string]interface{}{
		"book":   book,
		"number": number,
	}, nil)
}

func (c *Client) PrayerTiming(city string) (interface{}, error) {
	return c.makeRequest("GET", "/api/prayer/timing", map[string]interface{}{"city": city}, nil)
}

func (c *Client) Quran(surah, ayat string) (interface{}, error) {
	return c.makeRequest("GET", "/api/islamic/quran", map[string]interface{}{
		"surah": surah,
		"ayat":  ayat,
	}, nil)
}

func (c *Client) IslamicHadit(book, number string) (interface{}, error) {
	return c.makeRequest("GET", "/api/islamic/hadit", map[string]interface{}{
		"book":   book,
		"number": number,
	}, nil)
}

func (c *Client) Tahlil() (interface{}, error) {
	return c.makeRequest("GET", "/api/islamic/tahlil", nil, nil)
}

func (c *Client) Wirid() (interface{}, error) {
	return c.makeRequest("GET", "/api/islamic/wirid", nil, nil)
}

func (c *Client) Dua() (interface{}, error) {
	return c.makeRequest("GET", "/api/islamic/dua", nil, nil)
}

func (c *Client) Ayatkursi() (interface{}, error) {
	return c.makeRequest("GET", "/api/islamic/ayatkursi", nil, nil)
}

func (c *Client) SearchBooks() (interface{}, error) {
	return c.makeRequest("GET", "/api/get/books", nil, nil)
}

func (c *Client) GetBooks(category string) (interface{}, error) {
	return c.makeRequest("GET", "/api/get/books", map[string]interface{}{"category": category}, nil)
}
