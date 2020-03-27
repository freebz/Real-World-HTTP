// 예제 4 JSON을 로드할 때 형 변환을 해서 로드하는 날짜형

type DueDate struct {
	time.Time
}

func (d *DueDate) UnmarshalJSON(raw []byte) error {
	epoch, err := strconv.Atoi(string(raw))
	if err != nil {
		return err
	}
	d.Time = time.Unix(int64(epoch), 0)
	return nil
}
