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

	fmt.Println(-6 % 10)

	encrypt("0613206176")

	for t := 1_234_567_889; t <= 9_999_999_999; t++ {
		correct1 := check(9, 0, 0, 1, t)
		correct2 := check(5, 6, 0, 2, t)

		//correct3 := check(4, 0, 1, 1, t)
		//correct4 := check(3, 6, 1, 2, t)

		if correct1 && correct2 {
			//options = append(options, t)
			//fmt.Println(t)

			if t == 1_234_567_890 {
				fmt.Println("found")
			}
		}
	}
}

func check(encryptedDigit int, decryptedDigit int, i int, j int, t int) bool {
	s := math.Sin(float64(t + i))
	k := getFloatDigit(s, j)

	return (encryptedDigit+(10-k))%10 == decryptedDigit
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
	i := 0
	//str_t := strconv.Itoa(t)
	//str_dMsg := strconv.Itoa(dMsg)
	str_eMsg := ""

	for j := 0; j < len(str_dMsg); j++ {
		digit := int(str_dMsg[j] - '0')
		s := math.Sin(float64(t + i))
		s_digit := getFloatDigit(s, j+1)
		eDigit := (digit + s_digit) % 10

		str_eMsg += strconv.Itoa(eDigit)
	}

	fmt.Println(str_dMsg)
	fmt.Println(str_eMsg)

	//return input
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
