package handle

import (
	"context"
	"l_redis/Unit1/def"
	"log"
	"time"
)

func AddLuckyNumber(ctx context.Context, key string, luckyNumber []string) {
	def.Sadds(ctx, key, luckyNumber)

}

func GetLuckyNumber(ctx context.Context, key string, count int64) {
	luckyNumber, err := def.SrandMember(ctx, key, count)
	if err != nil {
		log.Fatalf("failed:%+v", err)
		return
	}
	for k, v := range luckyNumber {
		time.Sleep(1 * time.Second)
		log.Printf("第"+def.IntToStr(k+1)+"个中将号码:%s", v)
	}
	return
}
