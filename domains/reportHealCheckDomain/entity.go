package reportHealCheckDomain

type WebStat struct {
	Url          string
	ErrorMessage string
	Time         int64
	StatusCode   int
}

type ReportHealCheck struct {
	TotalWebSites int   `json:"total_websites"`
	Success       int   `json:"success"`
	Failure       int   `json:"failure"`
	TotalTime     int64 `json:"total_time"`
	WebStat       []*WebStat
}
