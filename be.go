package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello, world\n")

	rand.Seed(time.Now().UnixNano())

	computeProbability(0.5)
}

func computeProbability(probability float64) bool {
	random := rand.Float64()
	fmt.Println(random)
	if random < probability {
		return true
	} else {
		return false
	}
}