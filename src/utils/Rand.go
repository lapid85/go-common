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

// ServerKey 用于生成随机字符串的值
var McryptLettersLower = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

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
