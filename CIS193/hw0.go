// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import (
	"fmt"
	"math"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, दुनिया!")
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	str := ""
	if n % 3 == 0 {
		str += "Fizz"
	}
	if n % 5 == 0{
		str += "Buzz"
	}
	return str
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if n == 2{
		return true
	}
	if n % 2 == 0{
		return false
	} else {
		for i := 3; i < int(math.Sqrt(float64(n))); i += 2{
			if n % i == 0{
				return false
			}
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length - i - 1]{
			return false
		}
	}

	return true
}
