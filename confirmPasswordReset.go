package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
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
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&ConfirmPasswordResetRequest{
		OobCode: oobCode,
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("resetPassword"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &ConfirmPasswordResetResponse{}
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
