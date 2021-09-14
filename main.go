package main

import (
	"math"
	"strconv"
)

func main() {
	//fmt.Println("Hello World")

	// t = sequence number
	// j = digit

	//var options []int

	//fmt.Println(-6 % 10)
	//
	//encrypt("0613206176")
	//
	//for t := 1_234_567_889; t <= 9_999_999_999; t++ {
	//	correct1 := checkDigit(0, 9, 0, 0, t)
	//	correct2 := checkDigit(6, 4, 0, 1, t)
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
	k := getFloatDigit(s, j+1)
	return k
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

	return d == x
}

func getFloatDigit(number float64, digit int) int {
	return (int)(math.Abs(number)*math.Pow10(digit)) % 10
}

func encrypt(str_dMsg string) string {
	t := 1_234_567_890
	i := 0

	str_eMsg := ""

	for j := 0; j < len(str_dMsg); j++ {
		x := int(str_dMsg[j] - '0')

		k := digitKey(t, i, j)

		y := encryptDigit(x, k)

		str_eMsg += strconv.Itoa(y)
	}

	return str_eMsg
}
