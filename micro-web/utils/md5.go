package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Pw(password string) string {
	hash := md5.New()
	key := "414hsdffszd"
	hash.Write([]byte(password + key))
	return hex.EncodeToString(hash.Sum(nil))
}
