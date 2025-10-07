package discard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	APIKey       string
	BaseURL      string
	FullResponse bool
	HTTPClient   *http.Client
}

type Config struct {
	APIKey       string
	BaseURL      string
	FullResponse bool
	Timeout      time.Duration
}

type APIResponse struct {
	Creator string      `json:"creator,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Status  bool        `json:"status,omitempty"`
}

func NewClient(config Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://discardapi.dpdns.org"
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &Client{
		APIKey:       config.APIKey,
		BaseURL:      config.BaseURL,
		FullResponse: config.FullResponse,
		HTTPClient: &http.Client{
			Timeout: config.Timeout,
		},
	}, nil
}

func (c *Client) buildURL(endpoint string, params map[string]interface{}) string {
	u, _ := url.Parse(c.BaseURL + endpoint)
	q := u.Query()
	
	q.Set("apikey", c.APIKey)
	
	for key, value := range params {
		if value != nil {
			q.Set(key, fmt.Sprintf("%v", value))
		}
	}
	
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Client) makeRequest(method, endpoint string, params map[string]interface{}, body interface{}) (interface{}, error) {
	var req *http.Request
	var err error

	if method == "GET" {
		fullURL := c.buildURL(endpoint, params)
		req, err = http.NewRequest("GET", fullURL, nil)
	} else {
		fullURL := c.buildURL(endpoint, params)
		
		var jsonData []byte
		if body != nil {
			jsonData, err = json.Marshal(body)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal request body: %w", err)
			}
		}
		
		req, err = http.NewRequest(method, fullURL, bytes.NewBuffer(jsonData))
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: status %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(responseBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if c.FullResponse {
		return apiResp, nil
	}
	
	if apiResp.Result != nil {
		return apiResp.Result, nil
	}
	
	return apiResp, nil
}

func (c *Client) makeFormDataRequest(endpoint string, params map[string]interface{}, files map[string]io.Reader) (interface{}, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("apikey", c.APIKey)
  
	for key, value := range params {
		if value != nil {
			writer.WriteField(key, fmt.Sprintf("%v", value))
		}
	}

	for fieldName, file := range files {
		part, err := writer.CreateFormFile(fieldName, fieldName)
		if err != nil {
			return nil, fmt.Errorf("failed to create form file: %w", err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, fmt.Errorf("failed to copy file: %w", err)
		}
	}

	writer.Close()

	req, err := http.NewRequest("POST", c.BaseURL+endpoint, &buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: status %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(responseBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if c.FullResponse {
		return apiResp, nil
	}
	return apiResp.Result, nil
}

func (c *Client) SetFullResponse(value bool) {
	c.FullResponse = value
}

func (c *Client) GetFullResponse() bool {
	return c.FullResponse
}

func (c *Client) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.HTTPClient.Timeout = timeout
}
