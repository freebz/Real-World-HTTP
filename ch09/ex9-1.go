// 예제 9-1 HTTP 프로토콜 버전 확인

package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Protocol Version: %s\n", resp.Proto)
}
