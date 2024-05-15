package recommendation

import (
	"math/rand"
)

func GenerateRandomNumber() int {
	return rand.Intn(3) + 1
}