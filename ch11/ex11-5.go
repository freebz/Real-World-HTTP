// 예제 11-5 context를 사용한 타임아웃

import (
	"context"
)

func timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 타임아웃 기능이 있는 클라이언트
	oauth2.NewClient(ctx, conf.TokenSource(ctx, token))
	:
	// 긴 처리
}
