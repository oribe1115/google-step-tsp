package lib

import (
	"math/rand"
	"time"
)

func Rand(under int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(under)
}
