package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) SetAccountInfo(email, idToken, locale string) (*SetAccountInfoResponse, error) {
	data, err := json.Marshal(&SetAccountInfoRequest{
		Email:             email,
		IdToken:           idToken,
		ReturnSecureToken: true,
	})
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	httpReq, err := http.NewRequest("POST", client.getUrl("setAccountInfo"), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", client.httpHeaderContentType)
	httpReq.Header.Set("X-Firebase-Locale", locale)
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
		resData := &SetAccountInfoResponse{}
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
