package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SendPasswordResetEmailRequest struct {
	RequestType string `json:"requestType"`
	Email       string `json:"email"`
}

type SendPasswordResetEmailResponse struct {
	Kind  string `json:"kind"`
	Email string `json:"email"`
}

func (client *Client) SendPasswordResetEmail(email, locale string) (*SendPasswordResetEmailResponse, error) {
	data, err := json.Marshal(&SendPasswordResetEmailRequest{
		Email:       email,
		RequestType: "PASSWORD_RESET",
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
		resData := &SendPasswordResetEmailResponse{}
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
