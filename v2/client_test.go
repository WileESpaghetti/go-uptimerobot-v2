package uptimerobot

import (
	"testing"
)

func TestClient_NewRequest(t *testing.T) {
	testApiKey := "xxxx-xxxxxxxxxxxxxxxxxx"
	client := New(testApiKey)
	testRequest, _ := client.NewRequest("TestClient_NewRequest")

	actualApiKey := testRequest.FormValue(_POST_API_KEY)
	if actualApiKey != testApiKey {
		t.Errorf("API Key is not set properly\nExpected = \"%s\"\nActual = \"%s\"", testApiKey, actualApiKey)
	}
}
