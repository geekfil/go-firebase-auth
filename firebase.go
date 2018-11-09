package firebase

import (
	"fmt"
)

const (
	baseUrl = "https://www.googleapis.com/identitytoolkit/v3/relyingparty"
)

type ErrorResponse struct {
	Error struct {
		Errors []struct {
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
			Message string `json:"message"`
		} `json:"errors"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Client struct {
	apiKey                string
	apiEndpointFirestore  string
	httpHeaderContentType string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey,
		apiEndpointFirestore:  baseUrl,
		httpHeaderContentType: "application/json",
	}
}

func (client *Client) getUrl(path string) (string) {
	return fmt.Sprintf("%s/%s?key=%s", client.apiEndpointFirestore, path, client.apiKey)
}
