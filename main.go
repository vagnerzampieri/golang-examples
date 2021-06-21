package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dinners := []string{
		"tacos",
		"pizza",
		"ramen",
	}

	rand.Seed(time.Now().Unix())
	fmt.Printf("%s dinner tonight\n", dinners[int(rand.Intn(len(dinners)))])
}
