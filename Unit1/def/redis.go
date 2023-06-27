package def

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var Client *redis.Client

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	Client = rdb
}

func Set(ctx context.Context, key string, value interface{}) {
	Client.Set(ctx, key, value, time.Second*60)
}

func MGet(ctx context.Context, key ...string) ([]interface{}, error) {
	return Client.MGet(ctx, key...).Result()
}

func Incrby(ctx context.Context, key string, step int64) (int64, error) {
	return Client.IncrBy(ctx, key, step).Result()
}

func Get(ctx context.Context, key string) (interface{}, error) {
	return Client.Get(ctx, key).Result()
}

func SetNX(ctx context.Context, key string, tag bool) (bool, error) {
	return Client.SetNX(ctx, key, tag, time.Second*5).Result()
}

func Sadds(ctx context.Context, key string, member []string) (int64, error) {
	return Client.SAdd(ctx, key, member).Result()
}

func SrandMember(ctx context.Context, key string, count int64) ([]string, error) {
	return Client.SRandMemberN(ctx, key, count).Result()
}

// Sunion
func Sunion(ctx context.Context, key1, key2 string) ([]string, error) {
	return Client.SUnion(ctx, key1, key2).Result()
}

func Sinter(ctx context.Context, key1, key2 string) ([]string, error) {
	return Client.SInter(ctx, key1, key2).Result()
}

func Lpush(ctx context.Context, key, value string) (int64, error) {
	return Client.LPush(ctx, key, value).Result()
}

func RBpop(ctx context.Context, key string) ([]string, error) {
	return Client.BRPop(ctx, 10*time.Second, key).Result()
}
