package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	IDToken      string `json:"idToken"`
	Registered   bool   `json:"registered"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) SigninWithEmailPassword(email, password string) (*SigninWithEmailPasswordResponse, error) {
	data, err := json.Marshal(&SigninWithEmailPasswordRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	})
	if err != nil {
		return nil, err
	}
	httpRes, err := http.Post(client.getUrl("verifyPassword"), client.httpHeaderContentType, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	defer httpRes.Body.Close()
	resByte, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode == http.StatusOK {
		resData := &SigninWithEmailPasswordResponse{}
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
