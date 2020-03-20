// 예제 3-18 Go 언어로 DELETE 메서드 전송하기

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
