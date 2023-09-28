package utils

import "math/rand"

// RandString 获取随机的字符串
func RandString(n int) string {
	var Letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	b := make([]rune, n)
	total := len(Letters)
	for i := range b {
		b[i] = Letters[rand.Intn(total)]
	}
	return string(b)
}
