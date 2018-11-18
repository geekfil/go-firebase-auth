package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ConfirmEmailVerificationRequest struct {
	OobCode string `json:"oobCode"`
}

type ConfirmEmailVerificationResponse struct {
	Kind             string `json:"kind"`
	LocalID          string `json:"localId"`
	Email            string `json:"email"`
	PasswordHash     string `json:"passwordHash"`
	ProviderUserInfo []struct {
		ProviderID  string `json:"providerId"`
		FederatedID string `json:"federatedId"`
	} `json:"providerUserInfo"`
}

func (client *Client) ConfirmEmailVerification(oobCode string) (*ConfirmEmailVerificationResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&ConfirmEmailVerificationRequest{
		OobCode: oobCode,
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("getOobConfirmationCode"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &ConfirmEmailVerificationResponse{}
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
