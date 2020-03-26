// 예제 9-9 Upgrade() 메서드의 내용

h, ok := w.(http.Hijacker)
var rw *bufio.ReadWriter
netConn, rw, err = h.Hijack()
