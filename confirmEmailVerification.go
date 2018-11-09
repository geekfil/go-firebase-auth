package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	data, err := json.Marshal(&ConfirmEmailVerificationRequest{
		OobCode: oobCode,
	})
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}

	httpReq, err := http.NewRequest("POST", client.getUrl("getOobConfirmationCode"), bytes.NewBuffer(data))
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
		resData := &ConfirmEmailVerificationResponse{}
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
