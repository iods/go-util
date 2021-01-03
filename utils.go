package util

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Average Calculates the average of several numbers. */
func Average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

/*
Random Generates a random number. */
func Random() {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(7)
	fmt.Println(x)
}