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
		x := algo.BeforeShuffle(count, amount*100)
		fmt.Print(x, ",")
	}
	fmt.Println()
}
