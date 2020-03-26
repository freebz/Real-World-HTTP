// 예제 9-10 Pong 이벤트 핸들러

conn.SetPongHandler(
	func(string) error {
		conn.SetReadDeadline(
			time.Now().Add(pongWait));
	return nil}
)
