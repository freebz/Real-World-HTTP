// 예제 11-6 rate 패키지를 이용한 액세스 수 제한

package main

import (
	"time"
	"context"
	"golang.org/x/time/rate"
)

func main() {
	// 1초당 상한 횟수
	RateLimit := 10
	// 토큰의 최대 보유 수
	BucketSize := 10
	ctx := context.Background()
	e := rate.Every(time.Second/RateLimit)
	l := rate.NewLimiter(e, BucketSize)

	for _, task := range tasks {
		err := l.Wait(ctx)
		if err != nil {
			panic(err)
		}
		// 여기서 태스크를 실행한다
	}
}
