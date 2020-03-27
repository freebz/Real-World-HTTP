// 예제 11-4 gist에 투고

// gist 투고
gist := `{
  "description": "API example",
  "public": true,
  "files": {
    "hello_from_rest_api.txt": {
      "content": "Hello World"
    }
  }
}`
// 투고
resp2, err := client.Post("https://api.github.com/gists", "application/json", strings.NewReader(gist))
if err != nil {
	panic(err)
}
fmt.Println(resp2.Status)
defer resp2.Body.Close()
// 결과를 해석한다
type GistResult struct {
	Url string `json:"html_url"`
}
gistResult := &GistResult{}
err = json.NewDecoder(resp2.Body).Decode(&gistResult)
if err != nil {
	panic(err)
}
if gistResult.Url != "" {
	// 브라우저로 연다
	open.Start(gistResult.Url)
}
