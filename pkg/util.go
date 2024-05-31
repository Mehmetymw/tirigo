package util

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.NewSource(time.Now().Unix())
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandomUsername() string {
	return RandomString(8)
}

func RandomPassword() string {
	return RandomString(12)
}
func RandomEmail() string {
	return RandomString(8) + "@example.com"
}
func RandomID() string {
	return uuid.NewString()
}
