package hello

import (
	"fmt"
	"hello/sum"

	"rsc.io/quote/v3"
)

// Hello returns hello world
func Hello() string {
	return quote.HelloV3()
}

// Proverb returns a proverb
func Proverb() string {
	return quote.Concurrency()
}

// DoSum sums 1 + 1
func DoSum() int {
	return sum.Sum(1, 1)
}

// SayHi echos hi to terminal
func SayHi() {
	fmt.Println("hi")
}
