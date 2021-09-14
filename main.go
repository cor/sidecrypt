package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello World")

	// t = sequence number
	// j = digit

	//correct1 := check(5, 0, 42, 0)
	//correct2 := check(9, 6, 42, 0)

	//fmt.Println(correct1)
	//fmt.Println(correct2)

	for t := 0; t < 1000000000; t++ {
		correct1 := check(5, 0, t, 1)
		correct2 := check(9, 6, t, 2)

		//fmt.Println(getFloatDigit(math.Pi, t))

		if correct1 && correct2 {

			fmt.Println(t)
		}
		//fmt.Println(correct1)
		//fmt.Println(correct2)
	}
}

func check(encryptedDigit int, decryptedDigit int, t int, i int) bool {
	s := math.Sin(float64(t + i))
	k := getFloatDigit(s, i)

	return ((encryptedDigit + k) % 10) == decryptedDigit
}

func getFloatDigit(number float64, digit int) int {
	return (int)(math.Abs(number)*math.Pow10(digit)) % 10
}
