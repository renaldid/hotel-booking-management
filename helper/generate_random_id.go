package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomId(prefix string) string {
	min := 100000
	max := 500000

	rand.Seed(time.Now().UnixMilli())

	return prefix + strconv.Itoa(rand.Intn(max-min)+min)
}
