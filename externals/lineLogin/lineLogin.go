package lineLogin

import (
	"errors"
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver"
	"net/http"
)

var (
	ErrorUnableCreateRequestGetAccessToken       = errors.New("unable create request from path getAccessToken")
	ErrorUnableCreateRequestVerifyAccessToke      = errors.New("unable create request from path verifyAccessToke")
)

//go:generate mockery -name=LineLoginClientInterface
type LineLoginClientInterface interface {
	LineAuth()
	VerifyAccessToken(accessToken string) error
	GetAccessTokenFromCode(code string) (string, error)
	GetAccessTokenFromWebHook() (string ,error)
}

type ErrorMessage struct {
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type LineLoginClient struct {
	httpClient httpDriver.HttpClient

	endpoint string

	chanelId     string
	chanelSecret string
	redirectUrl  string
	lineAuthEndpoint string
	portWebHook string
}

func New(endpoint string, chanelId string, chanelSecret string, redirectUrl string, lineAuthEndpoint string, portWebHook string) *LineLoginClient {
	return &LineLoginClient{
		endpoint:   endpoint,
		httpClient: http.DefaultClient,

		chanelId:     chanelId,
		chanelSecret: chanelSecret,
		redirectUrl:  redirectUrl,
		lineAuthEndpoint: lineAuthEndpoint,
		portWebHook: portWebHook,
	}
}

func (client *LineLoginClient) setHttpClient(httpClient httpDriver.HttpClient) *LineLoginClient {
	client.httpClient = httpClient

	return client
}
