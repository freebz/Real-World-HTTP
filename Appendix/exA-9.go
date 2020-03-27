// 예제 9 data의 형이 정해지지 않은 이벤트 정보

// 공개용 구조체
type Event struct {
	Created time.Time
	Data    EventData
	Id      string
	Object  string
}

// 소문자로 외부에 공개하지 않는 JSON 읽기 쓰기용 구조체
type tmpEvent struct {
	Created EpochDate       `json:"create"`
	Data    json.RawMessage `json:"data"`
	Id      string          `json:"id"`
	Object  string          `json:"object"`
	Type    string          `json:"type"`
}

func (e *Event) UnmarshalJSON(raw []byte) error {
	// JSON 읽기 쓰기용 구조체로 일단 파싱한다
	tmp := &tmpEvent{}
	err := json.Unmarshal(raw, &tmp)
	if err != nil {
		return err
	}
	e.Created = time.Time(tmp.Created)
	e.Id      = tmp.Id
	e.Object  = tmp.Object
	// Data 내부의 형에 따라 각 구조체의 인스턴스를 생성
	switch tmp.Type {
	case "customer":
		customer := &Customer{}
		json.Unmarshal(tmp.Data, customer)
		e.Data = customer
	case "change":
		charge := &Charge{}
		json.Unmarshal(tmp.Data, charge)
		e.Data = charge
	}
}
