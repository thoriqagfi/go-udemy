package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var result string
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(alphabet))
		result += string(alphabet[randomIndex])
	}
	return result
}

func RandomOwnerName() string {
	return RandomString(6)
}

func RandomMoney() int {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}