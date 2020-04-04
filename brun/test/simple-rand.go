package main

import (
	"fmt"
	"math/rand"
	"resk/infra/algo"
	"time"
)

func main() {
	count, amount := int64(10), int64(100)
	rand.Seed(time.Now().UnixNano())
	for i := int64(0); i < count; i++ {
		x := algo.SimpleRand(count, amount*100)
		fmt.Print(float64(x)/100, ",")
	}
	fmt.Println()
}
