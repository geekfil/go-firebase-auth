package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type SetAccountInfoRequest struct {
	IdToken           string `json:"idToken"`
	Email             string `json:"email"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SetAccountInfoResponse struct {
	Kind             string `json:"kind"`
	LocalID          string `json:"localId"`
	Email            string `json:"email"`
	PasswordHash     string `json:"passwordHash"`
	ProviderUserInfo []struct {
		ProviderID  string `json:"providerId"`
		FederatedID string `json:"federatedId"`
	} `json:"providerUserInfo"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) SetAccountInfo(email, idToken, locale string) (*SetAccountInfoResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&SetAccountInfoRequest{
		Email:             email,
		IdToken:           idToken,
		ReturnSecureToken: true,
	}); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.getUrl("setAccountInfo"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	req.Header.Set("X-Firebase-Locale", locale)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &SetAccountInfoResponse{}
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
