// 예제 3-20 국제화 도메인을 아스키로 변환

package main

import (
	"fmt"
	"golang.org/x/net/idna"
)

func main() {
	src := "악력왕"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", src, ascii)
}

// 라이브러리 오류 시: go get -u golang.org/x/net/idna
