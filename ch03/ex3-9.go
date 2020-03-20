// 예제 3-9 Go 언어로 임의의 문자열을 POST로 전송하기

package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("텍스트")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		// 전송 실패
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
