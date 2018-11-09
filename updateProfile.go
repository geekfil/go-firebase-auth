package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type UpdateProfileRequest struct {
	IdToken           string `json:"idToken"`
	DisplayName       string `json:"displayName"`
	PhotoUrl          string `json:"photoUrl"`
	DeleteAttribute   string `json:"deleteAttribute"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type UpdateProfileResponse struct {
	Kind             string `json:"kind"`
	LocalID          string `json:"localId"`
	Email            string `json:"email"`
	DisplayName      string `json:"displayName"`
	PhotoURL         string `json:"photoUrl"`
	PasswordHash     string `json:"passwordHash"`
	ProviderUserInfo []struct {
		ProviderID  string `json:"providerId"`
		FederatedID string `json:"federatedId"`
		DisplayName string `json:"displayName"`
		PhotoURL    string `json:"photoUrl"`
	} `json:"providerUserInfo"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) UpdateProfile(dataProfile *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	dataProfile.ReturnSecureToken = true
	buff := &bytes.Buffer{}
	if err := json.NewEncoder(buff).Encode(dataProfile); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.getUrl("setAccountInfo"), buff)
	req.Header.Set("Content-Type", client.httpHeaderContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		data := &UpdateProfileResponse{}
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
