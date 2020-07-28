package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oommi04/ReportHealcheck/setup"
	"github.com/oommi04/ReportHealcheck/usecase/reportHealCheckUsecase"
	"github.com/oommi04/ReportHealcheck/utils/common"
)

var accessToken string

func init() {
	flag.StringVar(&accessToken, "accessToken", "", "acessToken from line login")
}

func main() {

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("plz specific .csv")
		os.Exit(0)
	}

	csvPath := flag.Args()[0]
	urlsLine, err := common.ReadCSV(csvPath)
	if err != nil {
		fmt.Println("cannot open csv file")
		panic(err)
	}

	cfgs := setup.SetupConfigs()
	lineLoginInstance := setup.SetupLineLogin(cfgs)

	if accessToken != "" {
		err := lineLoginInstance.VerifyAccessToken(accessToken)

		if err != nil {
			fmt.Println(err)
			fmt.Println("waiting for login")
			token, err := lineLoginInstance.GetAccessTokenFromWebHook()
			if err != nil {
				panic(err)
			}
			accessToken = token
		}

	} else {
		token, err := lineLoginInstance.GetAccessTokenFromWebHook()
		if err != nil {
			panic(err)
		}
		accessToken = token
	}

	reportInsance := setup.SetupReport(cfgs, accessToken)
	healCheckInstance := setup.SetupHealCheck()

	reportHealCheckUsecaseInstance := reportHealCheckUsecase.New(healCheckInstance, reportInsance)
	reportHealCheckUsecaseInstance.Create(urlsLine)
}
