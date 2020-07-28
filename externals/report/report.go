package report

import (
	"errors"
	"github.com/oommi04/ReportHealcheck/domains/reportHealCheckDomain"
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver"
	"net/http"
)

var (
	ErrorUnableCreateRequestReportHealCheck = errors.New("unable create request from path reportHealCheck")
)

//go:generate mockery -name=ReportClientInterface
type ReportClientInterface interface {
	ReportHealCheck(r reportHealCheckDomain.ReportHealCheck) error
}

type ReportClient struct {
	httpClient httpDriver.HttpClient

	endpoint string

	accessToken string
}

func New(endpoint string, accessToken string) *ReportClient {
	return &ReportClient{
		endpoint:   endpoint,
		httpClient: http.DefaultClient,

		accessToken: accessToken,
	}
}

func (client *ReportClient) setHttpClient(httpClient httpDriver.HttpClient) *ReportClient {
	client.httpClient = httpClient

	return client
}
