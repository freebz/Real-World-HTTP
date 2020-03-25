// 예제 6-6 HTTPS 통신을 하는 클라이언트 코드

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
