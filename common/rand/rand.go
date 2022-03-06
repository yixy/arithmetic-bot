package rand

import "math/rand"

func GetInt(i int) int {
	return rand.Int() % (i + 1)
}
