package firebase_auth

import (
	"bytes"
	"encoding/json"
	"errors"
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
	OauthIdToken  string `json:"oauthIdToken"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	FullName      string `json:"fullName"`
	DisplayName   string `json:"displayName"`
	IdToken       string `json:"idToken"`
	PhotoURL      string `json:"photoUrl"`
	RefreshToken  string `json:"refreshToken"`
	ExpiresIn     string `json:"expiresIn"`
	RawUserInfo   string `json:"rawUserInfo"`
}

func (client *Client) LinkWithOAuthCredential(idToken, requestUri, postBody string) (*LinkWithOAuthCredentialResponse, error) {
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(&LinkWithOAuthCredentialRequest{
		IdToken:             idToken,
		RequestUri:          requestUri,
		PostBody:            postBody,
		ReturnSecureToken:   true,
		ReturnIdpCredential: "EMAIL_EXISTS",
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("verifyAssertion"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &LinkWithOAuthCredentialResponse{}
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
