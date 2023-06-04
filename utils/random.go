package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomTitle generates a random item title
func RandomTitle() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() float32 {
	return float32(RandomInt(100, 1000))
}

// RandomQuantity generates a random quantity
func RandomQuantity() int64 {
	return RandomInt(0, 20)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return RandomString(6) + "@email.com"
}

// RandomPassword generates a random password
func RandomPassword() string {
	return RandomString(6)
}
