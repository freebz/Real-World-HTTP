// 예제 3-14 Go 언어의 오브젝트 생성

// 초기값을 지정해서 생성
a := Struct{
	Member: "Value",
}

// new 함수로 초기화
a := new(Struct)

// make 함수로 초기화
// 배열의 슬라이스, map, 채널 전용
a := make(map[string]string)

// 라이브러리가 준비한 빌더 함수로 생성
// 내부에서는 이 중 어느 하나를 이용
a := library.New()
