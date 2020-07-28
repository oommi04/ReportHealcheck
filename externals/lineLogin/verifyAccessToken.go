package lineLogin

import (
	"encoding/json"
	"errors"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"net/http"
)

var verifyAccessTokenPath = "/verify"

func(client *LineLoginClient)VerifyAccessToken(accessToken string) error {
	uri := client.endpoint + verifyAccessTokenPath
	req, err := http.NewRequest("GET", uri, nil)

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return ErrorUnableCreateRequestVerifyAccessToke
	}

	resp, err := client.httpClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorMessage{}
		_ = json.NewDecoder(resp.Body).Decode(&errorResp)
		return errors.New("error: " + errorResp.ErrorDescription + " status: " + common.IntToString(resp.StatusCode)+ " from " + uri)
	}

	defer resp.Body.Close()

	return nil
}