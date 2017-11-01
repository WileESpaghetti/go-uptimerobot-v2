package uptimerobot

import (
	"bytes"
	"net/http"
	"net/url"
	"github.com/gorilla/schema"
)

// Client constants
const (
	defaultUserAgent = "go-uptimerobot-v2/1.0.0 (Language=Go)"
)

// API constants
const (
	baseUrl = "https://api.uptimerobot.com/v2/"
)

type Client struct {
	ApiKey string `schema:"api_key"`
	HttpClient http.Client `schema:"-"`
}

func New(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
		HttpClient: http.Client{},
	}
}


func (c *Client) NewRequest(apiMethod string) (*http.Request, error) {
	apiUrl := baseUrl + apiMethod
	formEncoder := schema.NewEncoder()

	form := url.Values{}
	err := formEncoder.Encode(c, form)
	if err != nil {
		return nil, err
	}

	encodedForm := bytes.NewBufferString(form.Encode())
	req, err := http.NewRequest("POST", apiUrl, encodedForm)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)


	return req, nil
}


//func (*c Client) Do(req *http.Request) (*http.Response, error) {
//	return nil, nil
//}
