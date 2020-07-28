package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/oommi04/ReportHealcheck/domains/reportHealCheckDomain"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"net/http"
)

func (client *ReportClient) ReportHealCheck(r reportHealCheckDomain.ReportHealCheck) error {
	uri := client.endpoint
	accessToken := "Bearer " + client.accessToken

	body, _ := json.Marshal(r)

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", accessToken)

	if err != nil {
		return ErrorUnableCreateRequestReportHealCheck
	}

	resp, err := client.httpClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("error status: " + common.IntToString(resp.StatusCode) + " from " + uri)
	}

	defer resp.Body.Close()

	return nil
}
