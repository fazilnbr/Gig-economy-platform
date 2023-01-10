package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alpabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate randome inteager value between min and max

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)

}

// RandomString generate random string of lenth num

func RandomString(num int) string {
	var sb strings.Builder
	k := len(alpabet)
	for i := 0; i < num; i++ {
		c := alpabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
