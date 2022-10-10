package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Pw(password string) string {
	h := md5.New()
	md5Code := "23gasf3df"
	h.Write([]byte(password + md5Code))
	return hex.EncodeToString(h.Sum(nil))
}
