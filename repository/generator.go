package repository

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type SourceGenerator struct {
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *SourceGenerator) Generate(n int) (target string, err error) {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b), nil
}
