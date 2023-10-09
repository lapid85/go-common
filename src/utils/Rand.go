package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// ServerKey 服务器端的key
const ServerKey = "a!k@9#d%9*&aDQK-=bv<M+3456*8c0}@"

// McryptLetters  用于生成随机字符串的值
var McryptLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// McryptLettersLower 用于生成随机字符串的值
var McryptLettersLower = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// Letters 默认的可用字符
var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678")

// GetPassword 生成密码
func GetPassword(password, secret string) string {
	return MD5(password + "-" + ServerKey + "-" + secret)
}

// MD5 生成md5的值
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// GetSecret 随机生成密钥
func GetSecret() string {
	return GetRandString(32)
}

// GetRandString 生成随机字符串
func GetRandString(n uint) string {
	rand.NewSource(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = McryptLetters[rand.Intn(len(McryptLetters))]
	}
	return string(b)
}

// RandInt64 随机int64
func RandInt64(min, max int64) int64 {
	rand.NewSource(time.Now().Unix())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// RandString 生成随机字符串
func RandString(n int) string {
	b := make([]rune, n)
	total := len(Letters)
	for i := range b {
		b[i] = Letters[rand.Intn(total)]
	}
	return string(b)
}

// GenerateRangeNum 随机生成指定范围的数字比如6-9
func GenerateRangeNum(min, max int) int {
	if min > max {
		return 0
	}
	rand.NewSource(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}
