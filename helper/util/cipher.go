package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func Md5Hash(key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

func Encrypt(data []byte, s string) []byte {
	block, _ := aes.NewCipher([]byte(Md5Hash(s)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	return gcm.Seal(nonce, nonce, data, nil)
}

func Decrypt(data []byte, s string) []byte {
	key := []byte(Md5Hash(s))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ns := gcm.NonceSize()
	nonce, cipherText := data[:ns], data[ns:]
	text, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err)
	}
	return text
}
