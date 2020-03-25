// 예제 6-13 프로토콜 업그레이드를 하는 서버의 핸들러

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handlerUpgrade(w http.ResponseWriter, r *http.Request) {
	// 이 엔드포인트에서는 변경 외는 받아들이지 않는다
	if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "MyProtocol" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to MyProtocol")

	// 소켓을 획득
	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()

	// 프로토콜이 바뀐다는 응답을 보낸다
	response := http.Response{
		StatusCode: 101,
		Header:     make(http.Header),
	}
	response.Header.Set("Upgrade", "MyProtocol")
	response.Header.Set("Connection", "Upgrade")
	response.Write(conn)

	// 오리지널 통신 시작
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(readWriter, "%d\n", i)
		fmt.Println("->", i)
		readWriter.Flush() // Trigger "chunked" encoding and send a chunk...
		recv, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("<- %s", string(recv))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	http.HandleFunc("/upgrade", handlerUpgrade)
	log.Println("start http listening :18888")
	err := http.ListenAndServe(":18888", nil)
	log.Println(err)
}
