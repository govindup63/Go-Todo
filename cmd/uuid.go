package cmd

import (
	"math/rand"
	"time"
)

func NewId() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(9000) + 1000
	return num
}
