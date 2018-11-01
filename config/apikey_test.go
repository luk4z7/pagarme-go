package config

import (
	"os"
	"testing"
)

func TestGetApiKeyPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	apiKey := ApiKey{}
	apiKey.GetApiKey()
}

func TestGetApiKey(t *testing.T) {
	os.Setenv("PAGARME_APIKEY", "test")

	apiKey := ApiKey{}
	if apiKey.GetApiKey() != "test" {
		t.Errorf("GetApiKey was expecting %s but given %s", "test", apiKey.GetApiKey())
	}

	os.Unsetenv("PAGARME_APIKEY")
}
