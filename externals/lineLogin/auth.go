package lineLogin

import (
	"context"
	"fmt"
	"github.com/oommi04/ReportHealcheck/utils/common"
	"net/http"
)

func (client *LineLoginClient) LineAuth() {
	lineAuthUrl := client.lineAuthEndpoint
	chanelId := client.chanelId
	state := "random_state_str"
	redirectUrl := client.redirectUrl
	redirectUrlAuth := lineAuthUrl + "?response_type=code&client_id=" + chanelId + "&redirect_uri=" + redirectUrl + "&state=" + state + "&scope=openid profile"
	common.Openbrowser(redirectUrlAuth)
}

func (client *LineLoginClient) GetAccessTokenFromWebHook() (string, error) {
	channelCodeForAccessToken := make(chan string)
	code := ""

	//setup server for webhook
	addr := fmt.Sprintf(":%s", client.portWebHook)
	m := http.NewServeMux()
	s := http.Server{Addr: addr, Handler: m}
	m.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if code == "" && len(channelCodeForAccessToken) == 0 {
			channelCodeForAccessToken <- r.FormValue("code")
			close(channelCodeForAccessToken)
		}
		fmt.Fprintf(w, "Line Login Success")
	})

	go func(channelCodeForAccessToken chan string) {
		for c := range channelCodeForAccessToken {
			code = c
			s.Shutdown(context.Background())
		}
	}(channelCodeForAccessToken)

	done := make(chan bool)
	go func(done chan bool) {
		s.ListenAndServe()
		done <- true
	}(done)

	client.LineAuth()
	<-done

	token, err := client.GetAccessTokenFromCode(code)
	if err != nil {
		return "", err
	}
	return token, err
}
