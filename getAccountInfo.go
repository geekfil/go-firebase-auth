package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetAccountInfoRequest struct {
	IdToken string `json:"idToken"`
}

type GetAccountInfoResponse struct {
	Kind  string `json:"kind"`
	Users []struct {
		LocalID          string `json:"localId"`
		Email            string `json:"email"`
		EmailVerified    bool   `json:"emailVerified"`
		DisplayName      string `json:"displayName"`
		ProviderUserInfo []struct {
			ProviderID  string `json:"providerId"`
			DisplayName string `json:"displayName"`
			PhotoURL    string `json:"photoUrl"`
			FederatedID string `json:"federatedId"`
			Email       string `json:"email"`
			RawID       string `json:"rawId"`
			ScreenName  string `json:"screenName"`
		} `json:"providerUserInfo"`
		PhotoURL          string `json:"photoUrl"`
		PasswordHash      string `json:"passwordHash"`
		PasswordUpdatedAt int64  `json:"passwordUpdatedAt"`
		ValidSince        string `json:"validSince"`
		Disabled          bool   `json:"disabled"`
		LastLoginAt       string `json:"lastLoginAt"`
		CreatedAt         string `json:"createdAt"`
		CustomAuth        bool   `json:"customAuth"`
	} `json:"users"`
}

func (client *Client) GetAccountInfo(dataProfile *GetAccountInfoRequest) (*GetAccountInfoResponse, error) {

	data, err := json.Marshal(dataProfile)
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
		resData := &GetAccountInfoResponse{}
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
