package mathutil

import (
	"math/rand"
	"time"
)

// randomInt Generates a random integer using a min/max set.
func randomInt(min, max int) int {
	x := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + x.Intn(max-min)
}

// randomIntSeed Generates a random integer using a min/max set.
func randomIntSeed(min, max int, seed int64) int {
	x := rand.New(rand.NewSource(seed))
	return min + x.Intn(max-min)
}

// RandomInt provides an alias for randomInt
func RandomInt(min, max int) int {
	return randomInt(min, max)
}

// RandomIntSeed provides an alias for randomIntSeed
func RandomIntSeed(min, max int, seed int64) int {
	return randomIntSeed(min, max, seed)
}
