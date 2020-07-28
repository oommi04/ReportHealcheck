package lineLogin

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver/mocks"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"io/ioutil"
	"net/http"
)

func (suite *LineLoginServiceSuite) TestLineLoginClient_GetAccessToken_Success() {
	respBody := LineGetAccessTokenResponse{
		"token",
	}
	creatorJSON, _ := json.Marshal(respBody)

	resp := &http.Response{}
	resp.StatusCode = http.StatusOK
	resp.Body = ioutil.NopCloser(bytes.NewReader(creatorJSON))

	mock := &mocks.HttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return resp, nil
		},
	}
	suite.service.setHttpClient(mock)

	token, err := suite.service.GetAccessTokenFromCode("code")

	suite.NoError(err)
	suite.Equal(respBody.AccessToken, token)
}

func (suite *LineLoginServiceSuite) TestLineLoginClient_GetAccessToken_Fail() {
	uri := suite.service.endpoint + getAccessTokenPath

	respBody := ErrorMessage{
		ErrorDescription: "invalid authorization code",
	}
	creatorJSON, _ := json.Marshal(respBody)

	resp := &http.Response{}
	resp.StatusCode = http.StatusBadRequest
	resp.Body = ioutil.NopCloser(bytes.NewReader(creatorJSON))

	mock := &mocks.HttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return resp, nil
		},
	}
	suite.service.setHttpClient(mock)

	_, err := suite.service.GetAccessTokenFromCode("code")

	errorExpect := errors.New("error: " + respBody.ErrorDescription + " status: " + common.IntToString(resp.StatusCode) + " from " + uri)

	suite.Equal(errorExpect, err)
}
