// 예제 3-7 x-www-form-urlencoded 형식의 POST 메서드 전송

package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		// 전송 실패
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
