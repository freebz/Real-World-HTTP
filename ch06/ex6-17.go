// 예제 6-17 서버가 청크 형식으로 응답을 반환한다

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlerChunkedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

// main 안에서 /chunked라는 경로로 청크 전송하로독 핸들러를 등록한다
func main() {
	http.HandleFunc("/chunked", handlerChunkedResponse)
	log.Println("start http listening :18888")
	err := http.ListenAndServe(":18888", nil)
	log.Println(err)
}
