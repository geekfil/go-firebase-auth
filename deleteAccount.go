package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type DeleteAccountRequest struct {
	IdToken string `json:"idToken"`
}

type DeleteAccountResponse struct {
	Kind string `json:"kind"`
}

func (client *Client) DeleteAccount(idToken string) (*DeleteAccountResponse, error) {
	data, err := json.Marshal(&DeleteAccountRequest{
		IdToken: idToken,
	})
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}

	httpReq, err := http.NewRequest("POST", client.getUrl("deleteAccount"), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", client.httpHeaderContentType)
	httpRes, err := httpClient.Do(httpReq)

	defer httpRes.Body.Close()
	resByte, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode == http.StatusOK {
		resData := &DeleteAccountResponse{}
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
