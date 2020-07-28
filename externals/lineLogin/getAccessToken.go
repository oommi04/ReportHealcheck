package lineLogin

import (
	"encoding/json"
	"errors"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"net/http"
	"net/url"
	"strings"
)

var getAccessTokenPath = "/token"

type LineGetAccessTokenBody struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectUrl  string `json:"redirect_uri"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type LineGetAccessTest struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type LineGetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (client *LineLoginClient) GetAccessTokenFromCode(code string) (string, error) {
	uri := client.endpoint + getAccessTokenPath
	getAcessTokenBody := LineGetAccessTokenBody{
		"authorization_code",
		code,
		client.redirectUrl,
		client.chanelId,
		client.chanelSecret,
	}

	data := url.Values{}
	data.Set("grant_type", getAcessTokenBody.GrantType)
	data.Set("code", getAcessTokenBody.Code)
	data.Set("redirect_uri", getAcessTokenBody.RedirectUrl)
	data.Set("client_id", getAcessTokenBody.ClientId)
	data.Set("client_secret", getAcessTokenBody.ClientSecret)

	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	if err != nil {
		return "", ErrorUnableCreateRequestGetAccessToken
	}

	resp, err := client.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorMessage{}
		_ = json.NewDecoder(resp.Body).Decode(&errorResp)
		return "",errors.New("error: " + errorResp.ErrorDescription + " status: " + common.IntToString(resp.StatusCode)+ " from " + uri)
	}

	respInfo := LineGetAccessTokenResponse{}

	err = json.NewDecoder(resp.Body).Decode(&respInfo)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return respInfo.AccessToken, nil
}