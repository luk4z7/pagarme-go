package auth

import (
	"net/http"
	_ "net/http"
	_ "net/http/httptest"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	encode := BasicAuth("user", "12345")
	if encode != "dXNlcjoxMjM0NQ==" {
		t.Fatal("error")
	}
	err := requireInMethodBasicAuthReturn(encode)
	if err != nil {
		t.Fatal("error")
	}
}

func TestInit(t *testing.T) {
	resp := Init("https://google.com", Headers{
		"page":  "1",
		"count": "10",
	})
	_, err := requireInMethodInitReturn(resp)
	if err != nil {
		t.Fatal("error")
	}

	if resp.StatusCode == 400 {
		t.Fatal("error")
	}
}

func requireInMethodInitReturn(d *http.Response) (*http.Response, error) {
	return d, nil
}

func requireInMethodBasicAuthReturn(s string) error {
	return nil
}
