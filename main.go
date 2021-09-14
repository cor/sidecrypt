package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println("Hello World")

	// t = sequence number
	// j = digit

	//correct1 := check(5, 0, 42, 0)
	//correct2 := check(9, 6, 42, 0)

	//fmt.Println(correct1)
	//fmt.Println(correct2)

	//var options []int

	encrypt("0613206176")

	for t := 1_000_000_000; t <= 9_999_999_999; t++ {
		correct1 := check(5, 0, 0, 1, t)
		correct2 := check(9, 6, 0, 2, t)

		correct3 := check(4, 0, 1, 1, t)
		correct4 := check(3, 6, 1, 2, t)

		if correct1 && correct2 && correct3 && correct4 {
			//options = append(options, t)
			fmt.Println(t)
		}
	}
}

func check(encryptedDigit int, decryptedDigit int, i int, j int, t int) bool {
	s := math.Sin(float64(t + i))
	k := getFloatDigit(s, j)

	return ((encryptedDigit + k) % 10) == decryptedDigit
}

func getFloatDigit(number float64, digit int) int {
	return (int)(math.Abs(number)*math.Pow10(digit)) % 10
}

func getIntDigit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func encrypt(str_dMsg string) {
	t := 1_234_567_890
	//str_t := strconv.Itoa(t)
	//str_dMsg := strconv.Itoa(dMsg)
	str_eMsg := ""

	for j := 0; j < len(str_dMsg); j++ {
		digit := int(str_dMsg[j] - '0')
		s := math.Sin(float64(t + j))
		s_digit := getFloatDigit(s, j)
		eDigit := (digit + s_digit) % 10

		str_eMsg += strconv.Itoa(eDigit)
	}

	fmt.Println(str_dMsg)
	fmt.Println(str_eMsg)

	//return input
}
