package healcheck

import (
	"bytes"
	"encoding/json"
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver/mocks"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"io/ioutil"
	"net/http"
)

func(suite *HealCheckServiceSuite) TestHealCheck_HealCheckWebsite_Success() {
	uri := common.ParseURL("http://localhost:8000/")

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

	_, statusCode, err := suite.service.HealCheckWebsite(uri)

	suite.NoError(err)
	suite.Equal(resp.StatusCode, statusCode)
}
