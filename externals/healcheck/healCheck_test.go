package healcheck

import (
	"github.com/oommi04/ReportHealcheck/drivers/httpDriver/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HealCheckServiceSuite struct {
	suite.Suite

	http               *mocks.HttpClient
	service            *HealcheckClient
	integrationService *HealcheckClient
}

func Test_HealCheck_Service_Suite(t *testing.T) {
	suite.Run(t, new(HealCheckServiceSuite))
}

func (suite *HealCheckServiceSuite) SetupTest() {
	suite.service = New()
}
