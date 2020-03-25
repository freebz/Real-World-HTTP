// 예제 6-16 이 코드로도 청크 형식의 응답을 다룰 수 있다

resp, _ := http.Get("http://localhost:18888")
defer resp.Body.Close()
body, _ := ioutil.ReadAll(resp.Body)
