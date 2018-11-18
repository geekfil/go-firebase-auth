package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type DeleteAccountRequest struct {
	IdToken string `json:"idToken"`
}

type DeleteAccountResponse struct {
	Kind string `json:"kind"`
}

func (client *Client) DeleteAccount(idToken string) (*DeleteAccountResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&DeleteAccountRequest{
		IdToken: idToken,
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("deleteAccount"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &DeleteAccountResponse{}
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
