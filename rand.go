package str

import (
    "fmt"
    "time"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const size = 62

func RandSeq(n int) string {

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(size)]
	}

	return fmt.Sprintf("%d_%s", seed, string(b))

}
