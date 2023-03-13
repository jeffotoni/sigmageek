// @autor: jeffotoni
// @date: 2023-03-13
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int64) bool {
	if n <= 1 || n%2 == 0 {
		return false
	}
	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	digito := NumeroDigito()
	file, err := os.Open("pi.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	maxLineSize := 1000000
	var maxPrimePalindrome int64
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxLineSize), maxLineSize)
	var start bool
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i <= len(line)-digito; i++ {
			n, err := strconv.ParseInt(line[i:i+digito], 10, 64)
			if err == nil && isPalindrome(line[i:i+digito]) && isPrime(n) {
				if n > maxPrimePalindrome {
					maxPrimePalindrome = n
				}
				if !start {
					start = true
					fmt.Printf("Primeiro primo palindromo encontrado: %d\n", n)
				}
				next := findNextPrimePalindrome(n)
				fmt.Printf("Proximo primo palindromo encontrado: %d\n", next)
				break
			}
		}
	}
	if maxPrimePalindrome > 0 {
		fmt.Printf("Maior número primo palíndromo encontrado: %d\n", maxPrimePalindrome)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func findNextPrimePalindrome(current int64) int64 {
	for i := current + 1; ; i++ {
		if isPalindrome(strconv.FormatInt(i, 10)) && isPrime(i) {
			return i
		}
	}
}

func NumeroDigito() int {
	if len(os.Args) != 2 {
		fmt.Println("Use: go run main.go <numero_de_digitos>")
		os.Exit(0)
	}
	digito, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: inválido número de digitos")
		os.Exit(0)
	}
	if digito == 0 {
		return 9
	}
	return digito
}
