package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ChangePasswordRequest struct {
	IdToken           string `json:"idToken"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type ChangePasswordResponse struct {
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

func (client *Client) ChangePassword(newPassword, idToken, locale string) (*ChangePasswordResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&ChangePasswordRequest{
		Password:          newPassword,
		IdToken:           idToken,
		ReturnSecureToken: true,
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("setAccountInfo"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &ChangePasswordResponse{}
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
