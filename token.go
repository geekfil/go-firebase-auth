package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	Token string `json:"refresh_token"`
}

type TokenResponse struct {
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Token string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
	UserID       string `json:"user_id"`
	ProjectID    string `json:"project_id"`
}

func (client *Client) Token(refreshToken string) (*TokenResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&TokenRequest{
		Token: refreshToken,
		GrantType:    "refresh_token",
	}); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%s", client.apiKey), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &TokenResponse{}
		if json.NewDecoder(res.Body).Decode(data) != nil {
			return nil, err
		}
		return data, nil
	} else {
		data := &ErrorResponse{}
		if json.NewDecoder(res.Body).Decode(data) != nil {
			return nil, err
		}
		return nil, errors.New(data.Error.Message)
	}
}
