package main

import (
	"fmt"
	"testing"
)

func TestEncryptIdentity(t *testing.T) {

	for k := 0; k < 100000; k++ {
		for x := 0; x <= 9; x++ {
			if decryptDigit(encryptDigit(x, k), k) != x {
				t.Fatalf("decryptDigit(encryptDigit(x, k), k) != x. %v %v", x, k)
			}
		}
	}
}

func TestCheckDigit(t *testing.T) {
	tt := 1_234_567_890

	fmt.Println(encrypt("0612345678"))

	correct1 := checkDigit(0, 9, 0, 0, tt)
	correct2 := checkDigit(6, 4, 0, 1, tt)

	if !correct1 {
		t.Fatalf("incorrect 1")
	}

	if !correct2 {
		t.Fatalf("incorrect 2")
	}
}
