package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type SigninWithEmailPasswordRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SigninWithEmailPasswordResponse struct {
	Kind         string `json:"kind"`
	LocalID      string `json:"localId"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	IdToken      string `json:"idToken"`
	Registered   bool   `json:"registered"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) SigninWithEmailPassword(email, password string) (*SigninWithEmailPasswordResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&SigninWithEmailPasswordRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.getUrl("verifyPassword"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &SigninWithEmailPasswordResponse{}
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
