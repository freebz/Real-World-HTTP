// 예제 6-1 Keep-Alive를 위해 바디를 모두 읽는다

resp, err := http.Get("http://...")
if err != nil {
	// 오류 발생
	panic(err)
}
// 이 스코프를 벗어난 곳에서 반드시 닫는다
defer resp.Body.Close()
// ioutil.ReadAll로 서버 응답을 끝까지 일괄적으로 읽는다
body, err := iotuil.ReadAll(resp.Body)
