package handle

import (
	"context"
	"l_redis/Unit1/def"
)

const TOKEN = "@yQ3rPlYJxkRwg4UC$0Miz&vG9Ij*AQu3rgHY5*N4*Jul0lo2Ws8nKMXbdG#ds-$OXr!K-Z5AgoDqzX0gV#OjYJ6AK_aNfcTePIP"

func SetToken(ctx context.Context, key, value string) {
	def.Lpush(ctx, key, TOKEN)
}

func GetToken(ctx context.Context, key string) string {
	tokens, err := def.RBpop(ctx, key)
	if err != nil {
		return ""
	}
	return tokens[1]
}
