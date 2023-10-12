package captchas

import (
	"common/clients"
	"context"
	"time"
)

// StoreRedis 生成图片保存redis
type StoreRedis struct {
	Code string // 平台名称
}

// Set 设置redis的值
func (ths *StoreRedis) Set(id string, value string) error {
	siteCode := ths.Code
	rdClient := clients.GetRedisBySite(siteCode)
	ctx := context.Background()
	_, err := rdClient.Set(ctx, id, value, 300*time.Second).Result()
	return err
}

// Get 获取redis的值
func (ths *StoreRedis) Get(id string, clear bool) string {
	siteCode := ths.Code
	rdClient := clients.GetRedisBySite(siteCode)
	ctx := context.Background()
	v, _ := rdClient.Get(ctx, id).Result()
	if clear {
		rdClient.Del(ctx, id)
	}
	return v
}

// Verify 校验redis的值
func (ths *StoreRedis) Verify(id, answer string, clear bool) bool {
	v := ths.Get(id, clear)
	return v == answer
}
