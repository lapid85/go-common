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

// *** 警告 ***
// Set 设置redis的值
func (p *StoreRedis) Set(id string, value string) error {
	siteCode := p.Code
	rdClient := clients.GetRedisBySite(siteCode)
	ctx := context.Background()
	_, err := rdClient.Set(ctx, id, value, 300*time.Second).Result()
	return err
}

// *** 警告 ***
// Get 获取redis的值
func (p *StoreRedis) Get(id string, clear bool) string {
	siteCode := p.Code
	rdClient := clients.GetRedisBySite(siteCode)
	ctx := context.Background()
	v, _ := rdClient.Get(ctx, id).Result()
	if clear {
		rdClient.Del(ctx, id)
	}
	return v
}

// Verify 校验redis的值
func (p *StoreRedis) Verify(id, answer string, clear bool) bool {
	v := p.Get(id, clear)
	return v == answer
}
