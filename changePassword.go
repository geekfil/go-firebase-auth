package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) ChangePassword(newPassword, idToken, locale string) (*ChangePasswordResponse, error) {
	data, err := json.Marshal(&ChangePasswordRequest{
		Password:          newPassword,
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
		resData := &ChangePasswordResponse{}
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
