// 예제 1 JSON을 파싱해서 구조체로 변환

package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var jsonString = []byte(`
[
    {"title": "The Art of Community", "author": "Jono Bacon"},
    {"title": "Mithril", "author": "Yoshiki Shibukawa"}
]`)

func main() {
	var books []Book
	err := json.Unmarshal(jsonString, &books)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}
}
