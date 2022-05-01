package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var floatNum float32 = 254.5972
	fmt.Printf("number is %.2f", floatNum)

	string1 := "test string1"
	byte_string := []byte(string1)

	hex.EncodeToString(byte_string)
	fmt.Printf("hex string1: %v\n", byte_string)
}
