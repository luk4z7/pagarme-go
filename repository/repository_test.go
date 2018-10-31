package repository

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/luk4z7/pagarme-go/auth"
)

type MockedModel struct {
	value string `json:"value"`
}

func setupMockedHttpServer(body []byte, mockedUrl string, mockedMethod string) *httptest.Server {
	os.Setenv("PAGARME_APIKEY", "test")

	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == mockedUrl && r.Method == mockedMethod {
				w.Header().Add("Content-Type", "application/json")
				w.Write(body)
			}
		}),
	)
}

func TestGet(t *testing.T) {
	ts := setupMockedHttpServer([]byte(`{"value": "testing ok"}`), "/mocked_url/1", "GET")

	var mockedModel *MockedModel

	repository := Repository{}
	getResponse, _, _ := repository.Get(url.Values{"route": {ts.URL + "/mocked_url/1"}}, mockedModel, auth.Headers{})

	model := getResponse.(map[string]interface{})

	if model["value"] != "testing ok" {
		t.Errorf("invalid response data")
	}
}

func TestCreate(t *testing.T) {
	ts := setupMockedHttpServer([]byte(`{"status": "created"}`), "/mocked_url", "POST")

	repository := Repository{}

	requestBody := []byte(`{
		"bank_code": "184",
		"agencia": "0808",
		"agencia_dv": "8",
		"conta": "08808",
		"conta_dv": "8",
		"document_number": "80802694594",
		"legal_name": "Hober Mallow"
	}`)

	getResponse, _, _ := repository.Create(url.Values{"route": {ts.URL + "/mocked_url"}}, requestBody, auth.Headers{})

	model := getResponse.(map[string]interface{})

	if model["status"] != "created" {
		t.Errorf("invalid response data")
	}
}

func TestUpdate(t *testing.T) {
	ts := setupMockedHttpServer([]byte(`{"status": "updated"}`), "/mocked_url", "PUT")

	repository := Repository{}

	requestBody := []byte(`{"legal_name": "Hober Mallow"}`)

	getResponse, _, _ := repository.Update(url.Values{"route": {ts.URL + "/mocked_url"}}, requestBody, auth.Headers{})

	model := getResponse.(map[string]interface{})

	if model["status"] != "updated" {
		t.Errorf("invalid response data")
	}
}
