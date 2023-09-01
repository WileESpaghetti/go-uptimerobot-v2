package uptime_robot

import (
	"net/http"
)

type Client struct {
	HttpClient *http.Client
}

func New() *Client {
	return &Client{HttpClient: http.DefaultClient}
}
