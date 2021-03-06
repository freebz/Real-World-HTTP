// 예제 6-19 HTTP 규칙을 따라 청크를 다룬다

package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// TCP 소켓 열기
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 요청 보내기
	request, err := http.NewRequest("GET", "http://localhost:18888/chunked", nil)
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}
	// 읽기
	reader := bufio.NewReader(conn)
	// 헤더 읽기
	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}
	if resp.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}
	for {
		// 크기를 구하기
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		// 16진수의 크기를 해석, 크기가 0이면 닫는다
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		// 크기만큼 버퍼를 확보하고 읽어오기
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		log.Println(" ", string(line))
	}
}
