// 예제 5 DueDate를 이용한다

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// 이곳에 앞에서 설명한 DueDate 형의 정의를 기술한다

type ToDo struct {
	Task string  `json:"task"`
	Time DueDate `json:"due"`
}

var jsonString = []byte(`[
   {"task": "유치원등원", "due": 1486600200},
   {"task": "에릭슨연구회에 간다", "due": 1486634400}
]`)

func main() {
	var todos []ToDo
	err := json.Unmarshal(jsonString, &todos)
	if err != nil {
		panic(err)
	}
	for _, todo := range todos {
		fmt.Printf("%s: %v\n", todo.Task, todo.Time)
	}
}
