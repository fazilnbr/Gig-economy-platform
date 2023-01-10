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

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
	k := len(alpabet)
	for i := 0; i < num; i++ {

	}
}

//

func RandomString(num int) string {
	var sb strings.Builder
}
