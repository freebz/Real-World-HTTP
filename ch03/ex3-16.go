// 예제 3-16 프록시에 BASIC 인증 정보 설정하기

http.DefaultTransport = &http.Transport{
	Proxy: http.ProxyURL(proxyUrl),
}
