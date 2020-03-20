// 예제 3-13 쿠키 송수신을 위해 기본 클라이언트를 치환한다

http.DefaultClient = &http.Client{
	Jar: jar,
}

:

resp, err := http.Get("http://localhost:18888/cookie")
