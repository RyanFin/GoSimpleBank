package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random intesget between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generate a random owner
func RandomOwner() string {
	return RandomString(6)
}

// generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// generate a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "GBP", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// generate random account IDs for FromAccount and ToAccount
func RandomAccountIDs() (int64, int64) {
	fromAccountID, toAccountID := RandomInt(1, 81), RandomInt(1, 81)
	return fromAccountID, toAccountID
}
