package main

import (
	"golang.org/x/net/context"
	"l_redis/Unit1/def"
	"l_redis/Unit1/handle"
	"log"
	"time"
)

func getUUid(ctx context.Context) {
	key, err := handle.Counter(ctx, def.IntToStr(int(time.Now().UnixMicro())))
	if err != nil {
		log.Println(err)
	}
	log.Printf("生成key:%s", key)

}

func calcUrlCounter(ctx context.Context) {
	go func() {
		for {
			// 添加url浏览量
			handle.ArticleAddCounter(ctx, "www.baidu.com")
		}
	}()

	go func() {
		for {
			// 减少url浏览量
			handle.ArticleSubCounter(ctx, "www.baidu.com")
		}
	}()
	for {
		count := handle.GetArticleCounter(ctx, "www.baidu.com")
		log.Println(count)
	}
}

// todo 目前会出现 redis: connection pool timeout
func getSkillProduct(ctx context.Context) {
	go func() {
		for {
			def.SetNX(ctx, "product", true)
			p := handle.GetSkillProduct(ctx, "product")
			log.Printf("test : %+v", p)
		}
	}()

	go func() {
		for {
			def.SetNX(ctx, "product", true)
			p := handle.GetSkillProduct(ctx, "product")
			log.Printf("test2 : %+v", p)
		}
	}()
}

func getLuckyNumber(ctx context.Context) {
	redNumber := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33"}
	blueNumber := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"}
	handle.AddLuckyNumber(ctx, "redLucky", redNumber)
	handle.AddLuckyNumber(ctx, "blueLucky", blueNumber)
	time.Sleep(30 * time.Second)
	log.Println("Begin...")
	handle.GetLuckyNumber(ctx, "redLucky", 6)
	handle.GetLuckyNumber(ctx, "blueLucky", 1)
}

func main() {
	ctx := context.Background()
	//获取唯一id
	//getUUid(ctx)
	////计数问题
	//calcUrlCounter(ctx)
	////获取秒杀商品
	//getSkillProduct(ctx)
	//双色球随机抽奖
	//getLuckyNumber(ctx)
	//令牌桶
	go func() {
		IntervalTime := 30 * time.Second       // 觸發間隔時間
		ticker := time.NewTicker(IntervalTime) // 設定 2 秒
		for {
			select {
			case <-ticker.C:
				handle.SetToken(ctx, "key", "value")
			}
		}
	}()

	go func() {
		for {
			token := handle.GetToken(ctx, "key")
			log.Printf("token:%s", token)
		}
	}()

	for {
	}
}
