package math

import (
	"fmt"
	"math/rand"
	"time"
)

// Random Generates a random number.
func Random() {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(7)
	fmt.Println(x)
}
