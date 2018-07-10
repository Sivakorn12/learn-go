package main

import (
	"fmt"
	"math"
)

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	var total int
	var countPrime int
	for index := 1; index <= 100; index++ {
		if IsPrime(index) {
			total += index
			countPrime++
		}
	}
	avg := total / countPrime
	fmt.Println("Total:", total)
	fmt.Println("Avg:", avg)
}
