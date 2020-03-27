// 예제 3 포인터를 사용해 값이 있었는지 명시적으로 판정할 수 있게 한다

type EditHistory struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

if history.Price != nil {
	// JSON에 값이 들어왔다
	fmt.Println(*history.Price)
}
