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
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxLineSize), maxLineSize)

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i <= len(line)-digito; i++ {
			n, err := strconv.ParseInt(line[i:i+digito], 10, 64)
			if err == nil && isPalindrome(line[i:i+digito]) && isPrime(n) {
				fmt.Printf("Primeiro primo palindromo encontrado: %d\n", n)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
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
