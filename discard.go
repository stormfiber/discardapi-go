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

// Client represents the Discard API client
type Client struct {
	APIKey       string
	BaseURL      string
	FullResponse bool
	HTTPClient   *http.Client
}

// Config holds the configuration for the API client
type Config struct {
	APIKey       string
	BaseURL      string
	FullResponse bool
	Timeout      time.Duration
}

// APIResponse represents the standard API response structure
type APIResponse struct {
	Creator string      `json:"creator,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Status  bool        `json:"status,omitempty"`
}

// NewClient creates a new Discard API client
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

// buildURL constructs the full URL with query parameters
func (c *Client) buildURL(endpoint string, params map[string]interface{}) string {
	u, _ := url.Parse(c.BaseURL + endpoint)
	q := u.Query()
	
	// Add API key
	q.Set("apikey", c.APIKey)
	
	// Add other parameters
	for key, value := range params {
		if value != nil {
			q.Set(key, fmt.Sprintf("%v", value))
		}
	}
	
	u.RawQuery = q.Encode()
	return u.String()
}

// makeRequest performs the HTTP request
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

// MakeFormDataRequest performs a multipart form data request
func (c *Client) MakeFormDataRequest(endpoint string, params map[string]interface{}, files map[string]io.Reader) (interface{}, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Add API key
	if err := writer.WriteField("apikey", c.APIKey); err != nil {
		return nil, fmt.Errorf("failed to write apikey field: %w", err)
	}

	// Add parameters
	for key, value := range params {
		if value != nil {
			if err := writer.WriteField(key, fmt.Sprintf("%v", value)); err != nil {
				return nil, fmt.Errorf("failed to write field %s: %w", key, err)
			}
		}
	}

	// Add files
	for fieldName, file := range files {
		part, err := writer.CreateFormFile(fieldName, fieldName)
		if err != nil {
			return nil, fmt.Errorf("failed to create form file: %w", err)
		}
		if _, err := io.Copy(part, file); err != nil {
			return nil, fmt.Errorf("failed to copy file: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

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

// SetFullResponse sets the full response mode
func (c *Client) SetFullResponse(value bool) {
	c.FullResponse = value
}

// GetFullResponse gets the current full response mode
func (c *Client) GetFullResponse() bool {
	return c.FullResponse
}

// SetAPIKey updates the API key
func (c *Client) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

// SetTimeout updates the HTTP client timeout
func (c *Client) SetTimeout(timeout time.Duration) {
	c.HTTPClient.Timeout = timeout
}
