package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (client *Client) UpdateProfile(dataProfile *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	dataProfile.ReturnSecureToken = true
	data, err := json.Marshal(dataProfile)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	httpReq, err := http.NewRequest("POST", client.getUrl("setAccountInfo"), bytes.NewBuffer(data))
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
		resData := &UpdateProfileResponse{}
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
