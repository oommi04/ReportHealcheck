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

func (suite *LineLoginServiceSuite) TestLineLoginClient_VerifyAccessToken_Success() {
	var respBody interface{}
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

	err := suite.service.VerifyAccessToken("token")

	suite.NoError(err)
}

func (suite *LineLoginServiceSuite) TestLineLoginClient_VerifyAccessToken_Fail() {
	uri := suite.service.endpoint + verifyAccessTokenPath

	respBody := ErrorMessage{
		ErrorDescription: "access token expired",
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

	err := suite.service.VerifyAccessToken("token")

	errorExpect := errors.New("error: " + respBody.ErrorDescription + " status: " + common.IntToString(resp.StatusCode) + " from " + uri)

	suite.Equal(errorExpect, err)
}
