// 예제 6-10 클라이언트 인증서를 요청하는 방법

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// handle은 변하지 않으므로 생략
func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	server := &http.Server{
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			MinVersion: tls.VersionTLS12,
		},
		Addr: ":18443",
	}
	http.HandleFunc("/", handler)
	log.Println("start http listening :18443")
	err := server.ListenAndServeTLS("server.crt", "server.key")
	log.Println(err)
}
