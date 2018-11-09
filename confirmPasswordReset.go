package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ConfirmPasswordResetRequest struct {
	OobCode     string `json:"oobCode"`
	NewPassword string `json:"newPassword"`
}

type ConfirmPasswordResetResponse struct {
	Kind        string `json:"kind"`
	Email       string `json:"email"`
	RequestType string `json:"requestType"`
}

func (client *Client) ConfirmPasswordReset(oobCode string) (*ConfirmPasswordResetResponse, error) {
	data, err := json.Marshal(&ConfirmPasswordResetRequest{
		OobCode: oobCode,
	})
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}

	httpReq, err := http.NewRequest("POST", client.getUrl("resetPassword"), bytes.NewBuffer(data))
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
		resData := &ConfirmPasswordResetResponse{}
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
