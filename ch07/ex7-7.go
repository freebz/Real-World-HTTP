// 예제 7-7 Key와 Accept를 생성한다

package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	clientKeySrc := "the sample nonce"
	key := base64.StdEncoding.EncodeToString([]byte(clientKeySrc))
	fmt.Printf("Sec-WebSocket-Key: %s\n", key)
	// Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==

	salt := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	hash := sha1.Sum([]byte(key + salt))
	accept := base64.StdEncoding.EncodeToString(hash[:])
	fmt.Printf("Sec-WebSocket-Accept: %s\n", accept)
	// Sec-WebSocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=
}
