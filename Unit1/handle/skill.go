package handle

import (
	"context"
	"l_redis/Unit1/def"
	"log"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func GetSkillProduct(ctx context.Context, key string) Product {
	product := Product{
		Name:  "商品1",
		Price: 20,
		Count: 2,
	}
	value, err := def.Get(ctx, key)
	if err != nil {
		log.Fatalf("获取商品时候出现错误:%+v", err)
		return product
	}
	if value == "true" {
		return product
	} else {
		return product
	}
}
