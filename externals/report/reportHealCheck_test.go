package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/oommi04/ReportHealcheck/domains/reportHealCheckDomain"
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver/mocks"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"io/ioutil"
	"net/http"
)

func(suite *ReportServiceSuite) TestReport_ReportHealCheck_Success() {
	r := reportHealCheckDomain.ReportHealCheck{}

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

	err := suite.service.ReportHealCheck(r)

	suite.NoError(err)
}

func(suite *ReportServiceSuite) TestReport_ReportHealCheck_Fail() {
	uri := suite.service.endpoint
	r := reportHealCheckDomain.ReportHealCheck{}

	var respBody interface{}
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

	errorExpect := errors.New("error status: " + common.IntToString(resp.StatusCode)+ " from " + uri)

	err := suite.service.ReportHealCheck(r)

	suite.Equal(errorExpect, err)
}

