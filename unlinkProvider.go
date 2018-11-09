package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type UnlinkProviderRequest struct {
	IdToken        string   `json:"idToken"`
	DeleteProvider []string `json:"requestUri"`
}

type UnlinkProviderResponse struct {
	Kind             string `json:"kind"`
	LocalID          string `json:"localId"`
	Email            string `json:"email"`
	DisplayName      string `json:"displayName"`
	PhotoURL         string `json:"photoUrl"`
	PasswordHash     string `json:"passwordHash"`
	ProviderUserInfo []struct {
		ProviderID  string `json:"providerId"`
		FederatedID string `json:"federatedId"`
		DisplayName string `json:"displayName,omitempty"`
		PhotoURL    string `json:"photoUrl,omitempty"`
	} `json:"providerUserInfo"`
	EmailVerified string `json:"emailVerified"`
}

func (client *Client) UnlinkProvider(idToken string, deleteProvider []string) (*UnlinkProviderResponse, error) {
	data, err := json.Marshal(&UnlinkProviderRequest{
		IdToken:        idToken,
		DeleteProvider: deleteProvider,
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
		resData := &UnlinkProviderResponse{}
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
