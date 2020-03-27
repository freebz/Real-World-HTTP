// 예제 11-3 이메일 표시

// Email 취득
resp, err := client.Get("https://api.github.com/user/emails")
if err != nil {
	panic(err)
}
defer resp.Body.Close()
emails, err := ioutil.ReadAll(resp.Body)
if err != nil {
	panic(err)
}
fmt.Println(string(emails)
$ go run oauth2_sample.go
[{"email":"yoshiki@shibu.jp","primary":true,"verified":true}]
