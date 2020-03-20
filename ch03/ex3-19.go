// 예제 3-19 url.Values를 io.Reader로 변환

import (
	"strings"    // 이것들을 추가
	"net/url"
)

values := url.Values{"test": {"value"}}
reader := strings.NewReader(values.Encode())
