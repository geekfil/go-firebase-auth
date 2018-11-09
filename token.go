package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	IDToken      string `json:"id_token"`
	UserID       string `json:"user_id"`
	ProjectID    string `json:"project_id"`
}

func (client *Client) Token(refreshToken string) (*TokenResponse, error) {
	data, err := json.Marshal(&TokenRequest{
		Token: refreshToken,
		GrantType:    "refresh_token",
	})
	if err != nil {
		return nil, err
	}

	httpRes, err := http.Post(fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%s", client.apiKey), "application/x-www-form-urlencoded", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	defer httpRes.Body.Close()
	resByte, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode == http.StatusOK {
		resData := &TokenResponse{}
		if json.Unmarshal(resByte, resData) != nil {
			return nil, err
		}
		return resData, nil
	} else {
		resData := &ErrorResponse{}
		if json.Unmarshal(resByte, resData) != nil {
			return nil, err
		}
		return nil, errors.New(resData.Error.Message)
	}

}
