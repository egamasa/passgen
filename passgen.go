package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
)

const (
	digitChars     = "0123456789"
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbolChars    = "!@#$%^&*"
)

func randomStr(base []byte, length int) string {
	b := make([]byte, length)
	rand.Read(b)

	result := make([]byte, 0, length)
	for _, v := range b {
		result = append(result, base[int(v)%len(base)])
	}
	return string(result)
}

func main() {
	var baseCharsList = map[int]string{
		1: digitChars,
		2: lowercaseChars,
		4: uppercaseChars,
		8: symbolChars,
	}

	var (
		a = flag.String("a", "", "Add any characters include to passwords")
		l = flag.Int("l", 12, "Password length")
		n = flag.Int("n", 5, "Number of generate")
		s = flag.Int("s", 15, "Character-type stack level")
	)
	flag.Parse()

	if *l < 1 {
		fmt.Printf("Invalid password length value: %d\n", *l)
		os.Exit(0)
	}
	if (*s < 1) || (15 < *s) {
		fmt.Printf("Invalid Character-type stack level value: %d\n", *s)
		os.Exit(0)
	}

	baseChars := make([]byte, 0)
	for bit, chars := range baseCharsList {
		if *s&bit == bit {
			baseChars = append(baseChars, chars...)
		}
	}

	if *a != "" {
		baseChars = append(baseChars, *a...)
	}

	for i := 0; i < *n; i++ {
		result := randomStr(baseChars, *l)
		fmt.Printf("[%d] %s\n", i, result)
	}
}
