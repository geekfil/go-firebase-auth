package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type LinkWithOAuthCredentialRequest struct {
	IdToken             string `json:"idToken"`
	RequestUri          string `json:"requestUri"`
	PostBody            string `json:"postBody"`
	ReturnIdpCredential string `json:"returnIdpCredential"`
	ReturnSecureToken   bool   `json:"returnSecureToken"`
}

type LinkWithOAuthCredentialResponse struct {
	Kind          string `json:"kind"`
	FederatedID   string `json:"federatedId"`
	ProviderID    string `json:"providerId"`
	LocalID       string `json:"localId"`
	EmailVerified bool   `json:"emailVerified"`
	Email         string `json:"email"`
	OauthIDToken  string `json:"oauthIdToken"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	FullName      string `json:"fullName"`
	DisplayName   string `json:"displayName"`
	IDToken       string `json:"idToken"`
	PhotoURL      string `json:"photoUrl"`
	RefreshToken  string `json:"refreshToken"`
	ExpiresIn     string `json:"expiresIn"`
	RawUserInfo   string `json:"rawUserInfo"`
}

func (client *Client) LinkWithOAuthCredential(idToken, requestUri, postBody string) (*LinkWithOAuthCredentialResponse, error) {
	data, err := json.Marshal(&LinkWithOAuthCredentialRequest{
		IdToken:             idToken,
		RequestUri:          requestUri,
		PostBody:            postBody,
		ReturnSecureToken:   true,
		ReturnIdpCredential: "EMAIL_EXISTS",
	})
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	httpReq, err := http.NewRequest("POST", client.getUrl("verifyAssertion"), bytes.NewBuffer(data))
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
		resData := &LinkWithOAuthCredentialResponse{}
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
