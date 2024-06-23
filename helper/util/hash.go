package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
