// 예제 10 Go 언어의 세 가지 캐스팅 방법

// 특정형으로 캐스팅하기 (1)
bookList := books.([]interface{})

// 특정형으로 캐스팅하기 (2)
bookMap, ok := book.(map[string]interface{})

// switch문
switch v := value.(type) {
case bool:
	log.Println("bool", value)
case float64:
	log.Println("float64", value)
case string:
	log.Println("string", value)
case map[string]interface{}:
	log.Println("map", value)
case []interface{}:
	log.Println("array", value)
}
