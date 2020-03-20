// 예제 3-11 전송할 파일에 임의의 MIME 타입 설정하기

import (
	"net/textproto"
)

:

part := make(textproto.MIMEHeader)
part.Set("Content-Type", "image/jpeg")
part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
if err != nil {
	panic(err)
}
readFile, err := os.Open("photo.jpg")
if err != nil {
	panic(err)
}
io.Copy(fileWriter, readFile)

:
