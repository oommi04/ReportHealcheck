package healcheck

import (
	"net/http"
	"net/http/httptrace"
	"net/url"
	"time"
)

func (client *HealcheckClient) HealCheckWebsite(url *url.URL) (time.Duration, int, error) {
	var t1, t2, tEnd, tnh time.Time
	//statusCodeNoHost := 11001

	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return 0, 0, err
	}

	trace := &httptrace.ClientTrace{
		DNSStart: func(i httptrace.DNSStartInfo) {
			t1 = time.Now()
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) { t2 = time.Now() },
	}

	ctx := httptrace.WithClientTrace(req.Context(), trace)
	req = req.WithContext(ctx)
	// for no such host
	tnh = time.Now()

	resp, err := client.httpClient.Do(req)

	tEnd = time.Now()

	// when skip dns
	if t1.IsZero() {
		t1 = t2
	}

	if err != nil {
		return tEnd.Sub(tnh), 0, err
	}

	defer resp.Body.Close()

	return tEnd.Sub(t1), resp.StatusCode, nil
}
