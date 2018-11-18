package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
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
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&UnlinkProviderRequest{
		IdToken:        idToken,
		DeleteProvider: deleteProvider,
	}); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.getUrl("setAccountInfo"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &UnlinkProviderResponse{}
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
