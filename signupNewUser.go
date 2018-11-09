package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type SignupNewUserRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignupNewUserResponse struct {
	Kind         string `json:"kind"`
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
}

func (client *Client) SignupNewUser(email, password string) (*SignupNewUserResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&SignupNewUserRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.getUrl("signupNewUser"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &SignupNewUserResponse{}
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
