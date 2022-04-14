package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GenerateMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
