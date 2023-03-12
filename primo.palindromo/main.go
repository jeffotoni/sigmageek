package main

import (
	"fmt"
	"math"
	"math/big"
)

func isPalindrome(n int64) bool {
	s := fmt.Sprintf("%d", n)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	precision := 10000
	pi := big.NewFloat(math.Pi).SetPrec(uint(precision)).SetMode(big.ToNearestEven)
	s := pi.Text('f', precision)
	digitos := 5
	fmt.Println("pi:", pi)

	for i := 0; i < len(s)-digitos; i++ {
		n, _ := big.NewFloat(0).SetPrec(uint(precision)).SetString(s[i : i+3])
		intValue, _ := n.Int64()
		if isPalindrome(intValue) && isPrime(intValue) {
			fmt.Println(intValue)
			break
		}
	}
}
