package report

import (
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ReportServiceSuite struct {
	suite.Suite

	http           *mocks.HttpClient
	service            *ReportClient
	integrationService *ReportClient
}

func Test_Report_Service_Suite(t *testing.T) {
	suite.Run(t, new(ReportServiceSuite))
}

func (suite *ReportServiceSuite) SetupTest() {
	suite.service = New("http://localhost:8000", "token")
}