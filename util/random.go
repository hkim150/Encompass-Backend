package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns a random integer from range [min, max]
func RandomNumber(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// returns a random price from range [min, max)
func RandomPrice(min, max int) string {
	price := strconv.Itoa(RandomNumber(min, max-1))
	if i := rand.Intn(5); i < 2 {
		price += ".99"
	} else if i < 4 {
		price += ".00"
	} else {
		decimal := strconv.Itoa(RandomNumber(1, 99))
		if len(decimal) == 1 {
			decimal += "0"
		}
		price += "." + decimal
	}

	return price
}

func RandomUsername() string {
	length := RandomNumber(3, 7)
	return RandomString(length)
}

// generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	length := RandomNumber(6, 10)
	return RandomString(length) + "@gmail.com"
}
