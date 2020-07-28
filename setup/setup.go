package setup

import (
	"github.com/oommi04/ReportHealcheck/configs"
	"github.com/oommi04/ReportHealcheck/externals/healcheck"
	"github.com/oommi04/ReportHealcheck/externals/lineLogin"
	"github.com/oommi04/ReportHealcheck/externals/report"
)

func SetupConfigs() *configs.Configs {
	return configs.New()
}

func SetupLineLogin(cfg *configs.Configs) lineLogin.LineLoginClientInterface {
	return lineLogin.New("https://api.line.me/oauth2/v2.1", cfg.CHANELID, cfg.CHANELSECRET, cfg.REDIRECTURL, "https://access.line.me/oauth2/v2.1/authorize", cfg.PORT)
}

func SetupReport(cfg *configs.Configs, accessToken string) report.ReportClientInterface {
	return report.New(cfg.REPORTHEALCHECLURL, accessToken)
}

func SetupHealCheck() healcheck.HealcheckInterface {
	return healcheck.New()
}
