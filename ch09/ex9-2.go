// 예제 9-2 서버 푸시 이용 방법

func handler(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/style.css", nil)
	}
	
	// 일반 이벤트 핸들러 처리
}
