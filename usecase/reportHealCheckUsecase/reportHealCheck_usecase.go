package reportHealCheckUsecase

import (
	"fmt"
	"github.com/oommi04/ReportHealcheck/domains/reportHealCheckDomain"
	"github.com/oommi04/ReportHealcheck/externals/healcheck"
	"github.com/oommi04/ReportHealcheck/externals/report"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"sync"
)

type ReportHealCheckUsecaseInterface interface {
	Create(urlsLine *[][]string)
}

type ReportHealCheckUsecase struct {
	healCheckInstance healcheck.HealcheckInterface
	reportInsance report.ReportClientInterface
}

func New(h healcheck.HealcheckInterface, r report.ReportClientInterface) ReportHealCheckUsecaseInterface {
	return &ReportHealCheckUsecase{h,r }
}

func (rc *ReportHealCheckUsecase) Create(urlsLine *[][]string) {
	r := reportHealCheckDomain.ReportHealCheck{}

	r.TotalWebSites = len(*urlsLine)

	var wg sync.WaitGroup
	for _, urlLine := range *urlsLine {
		wg.Add(1)
		go func(urlFromCSV string, wg *sync.WaitGroup) {
			webStat := reportHealCheckDomain.WebStat{}

			url := common.ParseURL(urlFromCSV)

			totalTime, statusCode, err := rc.healCheckInstance.HealCheckWebsite(url)

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
		}(urlLine[0], &wg)
	}

	wg.Wait()

	err := rc.reportInsance.ReportHealCheck(r)
	if err != nil {
		fmt.Println(err, "cannot reporthealcheck")
	}

	fmt.Println("Perform website checking...")
	fmt.Println("Done!")
	fmt.Println("Checked webistes: ", r.TotalWebSites)
	fmt.Println("Successful websites: ", r.Success)
	fmt.Println("Failure websites: ", r.Failure)
	fmt.Println("Total times to finished checking website: ", r.TotalTime)
}
