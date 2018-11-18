package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type VerifyCustomTokenRequest struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type VerifyCustomTokenResponse struct {
	Kind         string `json:"kind"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) VerifyCustomToken(token string) (*VerifyCustomTokenResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&VerifyCustomTokenRequest{
		Token:             token,
		ReturnSecureToken: true,
	}); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.getUrl("verifyCustomToken"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &VerifyCustomTokenResponse{}
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
