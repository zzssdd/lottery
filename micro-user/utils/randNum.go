package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandNum(len int) string {
	rand.Seed(time.Now().UnixMicro())
	ret := ""
	for i := 0; i < len; i++ {
		ret += strconv.Itoa(rand.Intn(10))
	}
	return ret
}
