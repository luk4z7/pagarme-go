package auth

import (
	"net/http"
	_ "net/http"
	_ "net/http/httptest"
	"os"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	setPagarmeAPIKey()

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
	setPagarmeAPIKey()

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

func setPagarmeAPIKey() {
	os.Setenv("PAGARME_APIKEY", "TEST_VALUE")
}
