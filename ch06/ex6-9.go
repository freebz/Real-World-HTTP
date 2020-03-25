// 예제 6-9 인증서를 확인하지 않는 설정

tlsConfig := &tls.Config{
	InsecureSkipVerify: true,
}
