package util

import (
	"math/rand"
	"time"
)

func CalculateRandom(n int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randVal := rand.Intn(n)
	return randVal
}
