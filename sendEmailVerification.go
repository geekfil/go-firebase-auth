package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type SendEmailVerificationRequest struct {
	RequestType string `json:"requestType"`
	IdToken     string `json:"idToken"`
}

type SendEmailVerificationResponse struct {
	Kind  string `json:"kind"`
	Email string `json:"email"`
}

func (client *Client) SendEmailVerification(idToken, locale string) (*SendEmailVerificationResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&SendEmailVerificationRequest{
		IdToken:     idToken,
		RequestType: "VERIFY_EMAIL",
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("getOobConfirmationCode"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	req.Header.Set("X-Firebase-Locale", locale)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &SendEmailVerificationResponse{}
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
