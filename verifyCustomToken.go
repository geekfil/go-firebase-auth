package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type VerifyCustomTokenRequest struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type VerifyCustomTokenResponse struct {
	Kind         string `json:"kind"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) VerifyCustomToken(token string) (*VerifyCustomTokenResponse, error) {
	data, err := json.Marshal(&VerifyCustomTokenRequest{
		Token:             token,
		ReturnSecureToken: true,
	})
	if err != nil {
		return nil, err
	}

	httpRes, err := http.Post(client.getUrl("verifyCustomToken"), client.httpHeaderContentType, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	defer httpRes.Body.Close()
	resByte, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode == http.StatusOK {
		resData := &VerifyCustomTokenResponse{}
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
