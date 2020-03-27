// 예제 8 data 형이 정해지지 않은 이벤트 정보

type Event struct {
	Created EpochDate       `json:"create"`
	Data    json.RawMessage `json:"data"`
	Id      string          `json:"id"`
	Object  string          `json:"object"`
	Types   string          `json:"type"`
}

// data를 고객으로 간주해서 분석하고 반환한다
func (e Event) GetCustomerData() (*Customer, error) {
	if e.Type !== "customer" {
		return nil, fmt.Errorf("자료형은 customer가 아니라 %s", e.Type)
	}
	customer := &Customer{}
	err := json.Unmarshal(e.data, customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
