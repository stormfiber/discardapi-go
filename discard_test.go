package discard

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient(Config{
		APIKey: "test-key",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if client.APIKey != "test-key" {
		t.Errorf("Expected APIKey to be 'test-key', got '%s'", client.APIKey)
	}
	if client.BaseURL != "https://discardapi.dpdns.org" {
		t.Errorf("Expected default BaseURL, got '%s'", client.BaseURL)
	}
	_, err = NewClient(Config{})
	if err == nil {
		t.Error("Expected error when API key is missing")
	}
}

func TestSetFullResponse(t *testing.T) {
	client, _ := NewClient(Config{APIKey: "test"})
	
	client.SetFullResponse(true)
	if !client.GetFullResponse() {
		t.Error("Expected FullResponse to be true")
	}

	client.SetFullResponse(false)
	if client.GetFullResponse() {
		t.Error("Expected FullResponse to be false")
	}
}

func TestSetTimeout(t *testing.T) {
	client, _ := NewClient(Config{APIKey: "test"})
	
	newTimeout := 60 * time.Second
	client.SetTimeout(newTimeout)
	
	if client.HTTPClient.Timeout != newTimeout {
		t.Errorf("Expected timeout to be %v, got %v", newTimeout, client.HTTPClient.Timeout)
	}
}

func TestBuildURL(t *testing.T) {
	client, _ := NewClient(Config{APIKey: "test-key"})
	
	url := client.buildURL("/api/test", map[string]interface{}{
		"param1": "value1",
		"param2": 123,
	})
	
	if url == "" {
		t.Error("Expected non-empty URL")
	}
}

func TestIntegrationRandomQuote(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	client, err := NewClient(Config{
		APIKey: "YOUR_API_KEY_FOR_TESTING",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.RandomQuote()
	if err != nil {
		t.Errorf("RandomQuote failed: %v", err)
	}
}
