package uptime_robot

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, err := os.ReadFile(path.Join("mock_responses", r.URL.Path)) //read the content of file
		if err != nil {
			t.Errorf("could not read file: %s : %s", r.URL.Path, err)
			return
		}

		_, err = w.Write(v)
		if err != nil {
			t.Errorf("could not write request: %s", err)
		}
	}))
	defer server.Close()

	ur := &Client{
		Url:        server.URL + "/",
		HttpClient: server.Client(),
	}

	tests := map[string]struct {
		response api.Envelope
		wantErr  bool
	}{
		"error_not_authorized.json": {
			response: api.Envelope{},
			wantErr:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := ur.Get(name, &test.response, nil)
			gotErr := err != nil
			if test.wantErr != gotErr {
				t.Errorf("unexpected response error: wanted %t, got %t, error: %s", test.wantErr, gotErr, err)
			}

			t.Skipf("response: %+v", test.response)
		})
	}

}
