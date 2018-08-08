// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Feel free to use the main function for testing your functions
	//fmt.Println("Hello, دنيا!")
	//fmt.Print(ParsePhone("1234567890"))
	//fmt.Print(Anagram("binary", "brainy"))
	//fmt.Print(FindEvens([]int{1,2,3,4,5,6,7,8,9,10}))
	//fmt.Print(SliceProduct([]int{1,2,3,4,5,6,7,8,9}))
	fmt.Print(Unique([]int{1,1,1,2,3,4,5,5,4,3,2}))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	digits := make([]byte, 10)
	index := 0

	for _, r := range phone{
		if unicode.IsDigit(r){
			digits[index] = byte(r)
			index++
		}
	}

	fst := string(digits[:3])
	mid := string(digits[3:6])
	end := string(digits[6:10])
	return fmt.Sprintf("(%v) %v-%v", fst, mid, end)
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}

	m1 := make(map[rune]int)
	for _, r := range s1{
		m1[r]++
	}

	m2 := make(map[rune]int)
	for _, r := range s2{
		m2[r]++
	}

	for key, value := range m1 {
		if value != m2[key]{
			return false
		}
	}

	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var evens []int
	for _, num := range e{
		if num % 2 == 0{
			evens = append(evens, num)
		}
	}
	return evens
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	prod := 1
	for _, num := range e {
		prod *= num
	}
	return prod
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	//这是把map当set用啊。。。
	uniqueMap := make(map[int]bool)
	var uniques []int
	for _, num := range e{
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniques = append(uniques, num)
		}
	}

	return uniques
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	m := make(map[int]string)

	for key, value := range kv {
		m[value] = key
	}

	return m
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	m := make(map[rune]int)

	for _, r := range s {
		m[r] += 1
	}

	for key, value := range m{
		if value <= k {
			delete(m, key)
		}
	}

	return m
}
