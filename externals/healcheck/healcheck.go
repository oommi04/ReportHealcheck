package healcheck

import (
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver"
	"net/http"
	"net/url"
	"time"
)

//go:generate mockery -name=HealcheckInterface
type HealcheckInterface interface {
	HealCheckWebsite(url *url.URL) (time.Duration, int, error)
}

type HealcheckClient struct {
	httpClient httpDriver.HttpClient
}

func New() *HealcheckClient {
	c := HealcheckClient{
		httpClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 5 * time.Second,
		},
	}

	return &c
}

func (client *HealcheckClient) setHttpClient(httpClient httpDriver.HttpClient) *HealcheckClient {
	client.httpClient = httpClient

	return client
}
