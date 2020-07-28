package reportHealCheckUsecase

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/oommi04/ReportHealcheck/domains/reportHealCheckDomain"
	"github.com/oommi04/ReportHealcheck/externals/healcheck/mocks"
	mocks2 "github.com/oommi04/ReportHealcheck/externals/report/mocks"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"github.com/stretchr/testify/assert"
)

func TestReportHealCheckUsecase_Create_Success(t *testing.T) {
	mockHealCheckService := new(mocks.HealcheckInterface)
	mockReportService := new(mocks2.ReportClientInterface)

	u := ReportHealCheckUsecase{
		mockHealCheckService,
		mockReportService,
	}

	r := reportHealCheckDomain.ReportHealCheck{}

	urlsLine := [][]string{{"https://github.com/", "https://github.com/oommi04"}}

	r.TotalWebSites = len(urlsLine[0])

	var wg sync.WaitGroup
	for _, urlLine := range urlsLine[0] {
		wg.Add(1)
		go func(urlFromCSV string, wg *sync.WaitGroup) {
			webStat := reportHealCheckDomain.WebStat{}

			url := common.ParseURL(urlFromCSV)

			mockHealCheckService.On("HealCheckWebsite", url).Once().Return(time.Duration(5), 200, nil)

			totalTime, statusCode, err := u.healCheckInstance.HealCheckWebsite(url)

			r.TotalTime = r.TotalTime + totalTime.Nanoseconds()

			webStat.StatusCode = statusCode
			webStat.Url = url.String()
			webStat.Time = totalTime.Nanoseconds()

			if err != nil {
				r.Failure = r.Failure + 1
				webStat.ErrorMessage = err.Error()
			} else {
				r.Success = r.Success + 1
			}

			r.WebStat = append(r.WebStat, &webStat)
			wg.Done()
		}(urlLine, &wg)
	}

	wg.Wait()

	mockReportService.On("ReportHealCheck", r).Once().Return(nil)

	err := u.reportInsance.ReportHealCheck(r)

	assert.NoError(t, err)
	mockReportService.AssertExpectations(t)
	mockHealCheckService.AssertExpectations(t)

	fmt.Println("Perform website checking...")
	fmt.Println("Done!")
	fmt.Println("Checked webistes: ", r.TotalWebSites)
	fmt.Println("Successful websites: ", r.Success)
	fmt.Println("Failure websites: ", r.Failure)
	fmt.Println("Total times to finished checking website: ", r.TotalTime/int64(time.Millisecond), " ms")
}
