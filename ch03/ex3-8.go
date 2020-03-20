// 예제 3-8 Go 언어로 임의의 바디를 POST로 전송하기

package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		// 전송 실패
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
