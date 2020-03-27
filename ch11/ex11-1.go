// 예제 11-1 OAuth2 인가

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.x/oauth2/github"
	"io"
	"net/http"
	"os"
	"strings"
)

// 깃허브에서 가져온 내용을 붙여 넣는다
var clientID = "xxxxxxxxxxxxxxx"
var clientSecret = "xxxxxxxxxxxxxxxxxxxxxxxxxx"
var redirectURL = "https://localhost:18888"
var state = "your state"

func main() {
	// OAuth2 접속 정보
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"user:email", "gist"},
		Endpoint:     github.Endpoint,
	}
	// 이것을 이제부터 초기화한다
	var token *oauth2.Token

	// 로컬에 이미 저장됐는가?
	file, err := os.Open("access_token.json")
	if os.IsNotExist(err) {
		// 첫 액세스
		// 우선 인가 화면의 URL을 취득
		url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)

		// 콜백을 받을 웹 서버를 설정
		code := make(chan string)
		var server *http.Server
		server = &http.Server{
			Addr: ":18888",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// 쿼리 매개변수에서 code를 추출하고 브라우저를 닫는다
				w.Head().Set("Content-Type", "text/html")
				io.WriteString(w, "<html><script>window.open('abut:blank','_self').close()</script></html>")
				w.(http.Flusher).Flush()
				code <- r.URL.Query().Get("code")
				// 서버도 닫는다
				server.Shutdown(context.Background())
			}),
		}
		go server.ListenAndServe()

		// 브라우저로 인가 화면을 연다
		// 깃허브의 인가가 완료되면 상기 서버로 리디렉트되고
		// Handler가 실행된다
		open.Start(url)

		// 가져온 코드를 액세스 토큰으로 교환
		token, err = conf.Exchange(oauth2.NoContext, <-code)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(file).Encode(token)
	} else if err == nil {
		// 한 번 인가를 하고 로컬에 자정 완료
		token = &oauth2.Token{}
		json.NewDecoder(file).Decode(token)
	} else {
		panic(err)
	}
	client := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))
	// 이곳에서 다양한 처리를 한다
}
