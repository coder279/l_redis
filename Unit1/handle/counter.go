package handle

import (
	"context"
	"l_redis/Unit1/def"
	"strconv"
)

func Counter(ctx context.Context, timestamp string) (key string, err error) {
	seq, err := def.Incrby(ctx, timestamp, 1)
	if err != nil {
		return "", err
	}
	seqKey := strconv.Itoa(int(seq))
	key = timestamp + ":" + seqKey
	return key, nil
}

func ArticleAddCounter(ctx context.Context, url string) {
	def.Incrby(ctx, url, 1)
}

func ArticleSubCounter(ctx context.Context, url string) {
	def.Incrby(ctx, url, -1)
}

func GetArticleCounter(ctx context.Context, url string) int {
	num, err := def.Get(ctx, url)
	if err != nil {
		return 0
	} else {
		return def.StrToInt(num.(string))
	}
}
