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

	t := 1_234_567_890

	fmt.Println(encrypt(strconv.Itoa(t)))

	correct1 := checkDigit(0, 1, 0, 0, t)
	correct2 := checkDigit(6, 1, 0, 1, t)

	fmt.Println(correct1)
	fmt.Println(correct2)

	//var options []int

	//fmt.Println(-6 % 10)
	//
	//encrypt("0613206176")
	//
	//for t := 1_234_567_889; t <= 9_999_999_999; t++ {
	//	correct1 := check(9, 0, 0, 1, 1_234_567_890)
	//	correct2 := check(5, 6, 0, 2, t)
	//
	//	//correct3 := check(4, 0, 1, 1, t)
	//	//correct4 := check(3, 6, 1, 2, t)
	//
	//	if correct1 && correct2 {
	//		//options = append(options, t)
	//		//fmt.Println(t)
	//
	//		if t == 1_234_567_890 {
	//			fmt.Println("found")
	//		}
	//	}
	//}

	//k := 1234
	//for i := 0; i <= 9; i++ {
	//	fmt.Printf("Unencrypted: %v encrypted: %v decrypted: %v\n", i, encryptDigit(i, k), decryptDigit(encryptDigit(i, k), k))
	//}
}

func digitKey(t int, i int, j int) int {
	s := math.Sin(float64(t + i))
	return getFloatDigit(s, j)
}

func encryptDigit(x int, k int) int {
	return (x + k) % 10
}

func decryptDigit(y int, k int) int {
	return (y - (k % 10) + 10) % 10
}

func checkDigit(x int, y int, i int, j int, t int) bool {

	k := digitKey(t, i, j)

	d := decryptDigit(y, k)

	fmt.Printf("k: %v, d: %v\n", k, d)
	return d == x
}

func getFloatDigit(number float64, digit int) int {
	return (int)(math.Abs(number)*math.Pow10(digit)) % 10
}

//func getIntDigit(num, place int) int {
//	r := num % int(math.Pow(10, float64(place)))
//	return r / int(math.Pow(10, float64(place-1)))
//}

func encrypt(str_dMsg string) string {
	t := 1_234_567_890
	i := 0
	//str_t := strconv.Itoa(t)
	//str_dMsg := strconv.Itoa(dMsg)
	str_eMsg := ""

	for j := 0; j < len(str_dMsg); j++ {
		digit := int(str_dMsg[j] - '0')

		k := digitKey(t, i, j)

		eDigit := encryptDigit(digit, k)

		str_eMsg += strconv.Itoa(eDigit)
	}

	return str_eMsg
	//fmt.Println(str_dMsg)
	//fmt.Println(str_eMsg)

	//return input
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
