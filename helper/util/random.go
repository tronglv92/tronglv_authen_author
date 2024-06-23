package util

import (
	"crypto/rand"
	"math/big"
)

func RandomString(length int) (string, error) {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-0123456789"
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		bytes[i] = charset[num.Int64()]
	}
	return string(bytes), nil
}
func GenSalt(length int) (string, error) {
	if length < 0 {
		length = 50
	}
	return RandomString(length)
}
