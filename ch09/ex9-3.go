// 예제 9-3 서버 푸시 구현

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var image []byte

// 이미지 파일 준비
func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

// HTML을 브라우저로 송신
// 이미지를 푸시한다
func handlerHtml(w http.ResponseWriter, r *http.Request) {
	// Pusher로 캐스팅 가능하면(HTTP/2로 접속했다면) 푸시한다
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

// 이미지 파일을 브라우저로 송신
func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
	fmt.Println(err)
}
