package utils

import "math/rand"

type Random struct {
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
