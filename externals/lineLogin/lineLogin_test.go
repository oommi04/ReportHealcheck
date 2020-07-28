package lineLogin

import (
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LineLoginServiceSuite struct {
	suite.Suite

	http               *mocks.HttpClient
	service            *LineLoginClient
	integrationService *LineLoginClient
}

func Test_LineLogin_Service_Suite(t *testing.T) {
	suite.Run(t, new(LineLoginServiceSuite))
}

func (suite *LineLoginServiceSuite) SetupTest() {
	suite.http = &mocks.HttpClient{}
	suite.service = New("http://localhost:8000", "123456789", "123456789", "google.com", "localhost", "8000").setHttpClient(suite.http)
}
