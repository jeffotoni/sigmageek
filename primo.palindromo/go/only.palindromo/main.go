// @autor: jeffotoni
// @date: 2023-03-13
package main

import (
	"fmt"
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
	for i := int64(100000000); i <= 999999999; i++ {
		if isPalindrome(i) && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
