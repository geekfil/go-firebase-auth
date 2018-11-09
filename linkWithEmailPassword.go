package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	IDToken       string `json:"idToken"`
	RefreshToken  string `json:"refreshToken"`
	ExpiresIn     string `json:"expiresIn"`
	EmailVerified bool   `json:"emailVerified"`
}

func (client *Client) LinkWithEmailPassword(idToken, email, password, returnSecureToken string) (*LinkWithEmailPasswordResponse, error) {

	data, err := json.Marshal(&LinkWithEmailPasswordRequest{
		IdToken:           idToken,
		Email:             email,
		Password:          password,
		ReturnSecureToken: returnSecureToken,
	})
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	httpReq, err := http.NewRequest("POST", client.getUrl("getAccountInfo"), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", client.httpHeaderContentType)
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	resByte, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode == http.StatusOK {
		resData := &LinkWithEmailPasswordResponse{}
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
