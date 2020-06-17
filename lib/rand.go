package lib

import (
	"math/rand"
	"time"
)

// Rand [0,under)の範囲で乱数を返す
func Rand(under int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(under)
}
