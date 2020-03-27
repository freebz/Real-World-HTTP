// 예제 11-2 http.Client 작성

client := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))
