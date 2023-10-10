package userhash

import (
	"consts/consts"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hash/crc32"
	"strconv"
)

// UserNameToIdHashNum 通过用户名按id取模分配的hash缓存，通过用户名来获取用户id
var UserNameToIdHashNum = 2000

// HashCodeByString 加码以字符串
func HashCodeByString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

// SetUserNameToId 通过用户名的hashcode来写入到hash缓存
func SetUserNameToId(username string, userId int64, rClient *redis.Client) {
	hc := HashCodeByString(username)
	modId := hc % UserNameToIdHashNum                       //通过用户名按id取模分配的hash缓存，通过用户名来获取用户id //设置hash缓存
	cacheKey := consts.UserNameFindId + strconv.Itoa(modId) //设置hash缓存
	//写入的时候，可以不存在当前的cacheKey
	ctx := context.Background()
	hasInt, err := rClient.HSet(ctx, cacheKey, username, strconv.Itoa(int(userId))).Result()
	if hasInt > 0 && err != nil {
		rClient.Expire(ctx, cacheKey, consts.ForeverExpiration)
	}
}

// GetIdByUserName 通过用户名获取hash的缓存的key来获取值
func GetIdByUserName(username string, rClient *redis.Client) (int64, error) {
	hc := HashCodeByString(username)
	modId := hc % UserNameToIdHashNum
	//通过用户名按id取模分配的hash缓存，通过用户名来获取用户id
	cacheKey := consts.UserNameFindId + strconv.Itoa(modId)
	ctx := context.Background()
	isExists, err := rClient.HExists(ctx, cacheKey, username).Result()
	if err != nil {
		fmt.Println("无法获取用户ID: ", err)
		return 0, err
	}
	//也就是检查exists的时候，即使没值，err 是nil, isExists是0
	if !isExists {
		return 0, nil
	}

	//设置hash缓存
	v, err := rClient.HGet(ctx, cacheKey, username).Result()
	//这种直接取值，就要 注意err == redis.Nil
	if err != nil {
		return 0, err
	}
	if v == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return i, nil
}
