// 예제 6 MarshalJSON의 정의

// 날짜 시리얼라이즈
func (d *DueDate) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(d.Unix()))), nil
}

type ToDoList []ToDo

// 리스트를 필터링해서 시리얼라이즈
func (l ToDoList) MarshalJSON() ([]byte, error) {
	tmpList := make([]ToDo, 0, len(l))
	for _, todo := range l {
		if !todo.Done {
			tmpList = append(tmpList, todo)
		}
	}
	return json.Marshal(tmpList)
}
