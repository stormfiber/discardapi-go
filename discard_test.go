/*
package discard

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

type mockRoundTripper struct {
	lastRequest *http.Request
	response    *http.Response
	err         error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.lastRequest = req
	if m.err != nil {
		return nil, m.err
	}
	if m.response != nil {
		return m.response, nil
	}
	body := io.NopCloser(strings.NewReader(`{"creator":"Qasim Ali 🩷","status":true,"result":"ok"}`))
	return &http.Response{
		StatusCode: 200,
		Body:       body,
		Header:     make(http.Header),
	}, nil
}

// updated to accept *testing.T for error handling
func newMockClient(t *testing.T) *Client {
	c, err := NewClient(Config{APIKey: "test"})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	mock := &mockRoundTripper{}
	c.HTTPClient.Transport = mock
	c.SetFullResponse(true)
	return c
}

func TestMakeRequestFlexible(t *testing.T) {
	client := newMockClient(t)
	resp, err := client.makeRequest("GET", "/api/test", nil, nil)
	if err != nil {
		t.Fatalf("makeRequest failed: %v", err)
	}
	ar, ok := resp.(APIResponse)
	if !ok {
		t.Fatalf("expected APIResponse, got %T", resp)
	}
	if !ar.Status || ar.Creator != "Qasim Ali 🩷" {
		t.Fatalf("unexpected APIResponse: %+v", ar)
	}
}

func TestMakeFormDataRequest(t *testing.T) {
	client := newMockClient(t)
	buf := bytes.NewBufferString("dummy data")
	files := map[string]io.Reader{"file": buf}
	resp, err := client.MakeFormDataRequest("/api/upload", map[string]interface{}{"type": "x"}, files)
	if err != nil {
		t.Fatalf("MakeFormDataRequest failed: %v", err)
	}
	ar, ok := resp.(APIResponse)
	if !ok {
		t.Fatalf("expected APIResponse, got %T", resp)
	}
	if !ar.Status {
		t.Fatalf("expected success APIResponse, got %+v", ar)
	}
}

func TestReflectionCoversAllMethods(t *testing.T) {
	client := newMockClient(t)
	val := reflect.ValueOf(client)
	typ := reflect.TypeOf(client)

	for i := 0; i < typ.NumMethod(); i++ {
		m := typ.Method(i)
		if !m.IsExported() || strings.HasPrefix(m.Name, "Set") || strings.HasPrefix(m.Name, "Get") {
			continue
		}

		args := []reflect.Value{val}
		for j := 1; j < m.Type.NumIn(); j++ {
			argType := m.Type.In(j)
			switch argType.Kind() {
			case reflect.String:
				args = append(args, reflect.ValueOf("test"))
			case reflect.Int:
				args = append(args, reflect.ValueOf(1))
			case reflect.Slice, reflect.Map, reflect.Interface:
				args = append(args, reflect.Zero(argType))
			default:
				args = append(args, reflect.Zero(argType))
			}
		}

		func() {
			defer func() { _ = recover() }()
			m.Func.Call(args)
		}()
	}
}

func TestTimeoutAndConfig(t *testing.T) {
	client, err := NewClient(Config{APIKey: "x"})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	client.SetTimeout(10 * time.Second)
	if client.HTTPClient.Timeout != 10*time.Second {
		t.Errorf("timeout not applied properly")
	}

	client.SetFullResponse(true)
	if !client.GetFullResponse() {
		t.Errorf("FullResponse not set")
	}
}

func TestBuildURL(t *testing.T) {
	client, err := NewClient(Config{APIKey: "x"})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	urlStr := client.buildURL("/api/test", map[string]interface{}{"q": "1", "x": 5})
	u, err := url.Parse(urlStr)
	if err != nil {
		t.Fatalf("failed to parse URL: %v", err)
	}

	query := u.Query()
	if query.Get("apikey") != "x" {
		t.Errorf("expected apikey=x, got %s", query.Get("apikey"))
	}
	if query.Get("q") != "1" {
		t.Errorf("expected q=1, got %s", query.Get("q"))
	}
	if query.Get("x") != "5" {
		t.Errorf("expected x=5, got %s", query.Get("x"))
	}
}
*/


package discard

import (
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	// Test with valid config
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

	// Test without API key
	_, err = NewClient(Config{})
	if err == nil {
		t.Error("Expected error when API key is missing")
	}
}

func TestSetFullResponse(t *testing.T) {
	client, err := NewClient(Config{APIKey: "test"})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

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
	client, err := NewClient(Config{APIKey: "test"})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	newTimeout := 60 * time.Second
	client.SetTimeout(newTimeout)

	if client.HTTPClient.Timeout != newTimeout {
		t.Errorf("Expected timeout to be %v, got %v", newTimeout, client.HTTPClient.Timeout)
	}
}

func TestBuildURL(t *testing.T) {
	client, err := NewClient(Config{APIKey: "test-key"})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	url := client.buildURL("/api/test", map[string]interface{}{
		"param1": "value1",
		"param2": 123,
	})

	if url == "" {
		t.Error("Expected non-empty URL")
	}
}

// Integration tests (require valid API key from environment)
func TestIntegrationRandomQuote(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	apiKey := os.Getenv("DISCARD_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test: DISCARD_API_KEY not set")
	}

	client, err := NewClient(Config{
		APIKey: apiKey,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.RandomQuote()
	if err != nil {
		t.Errorf("RandomQuote failed: %v", err)
	}
}

func TestIntegrationDadJoke(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	apiKey := os.Getenv("DISCARD_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test: DISCARD_API_KEY not set")
	}

	client, err := NewClient(Config{
		APIKey: apiKey,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DadJoke()
	if err != nil {
		t.Errorf("DadJoke failed: %v", err)
	}
}
