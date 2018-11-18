package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
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

	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(dataProfile); err != nil {
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
		data := &GetAccountInfoResponse{}
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
