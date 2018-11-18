package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type LinkWithEmailPasswordRequest struct {
	IdToken           string `json:"idToken"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken string `json:"returnSecureToken"`
}

type LinkWithEmailPasswordResponse struct {
	Kind             string `json:"kind"`
	LocalID          string `json:"localId"`
	Email            string `json:"email"`
	DisplayName      string `json:"displayName"`
	PhotoURL         string `json:"photoUrl"`
	PasswordHash     string `json:"passwordHash"`
	ProviderUserInfo []struct {
		ProviderID  string `json:"providerId"`
		FederatedID string `json:"federatedId"`
	} `json:"providerUserInfo"`
	IdToken       string `json:"idToken"`
	RefreshToken  string `json:"refreshToken"`
	ExpiresIn     string `json:"expiresIn"`
	EmailVerified bool   `json:"emailVerified"`
}

func (client *Client) LinkWithEmailPassword(idToken, email, password, returnSecureToken string) (*LinkWithEmailPasswordResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&LinkWithEmailPasswordRequest{
		IdToken:           idToken,
		Email:             email,
		Password:          password,
		ReturnSecureToken: returnSecureToken,
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("getAccountInfo"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &LinkWithEmailPasswordResponse{}
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
