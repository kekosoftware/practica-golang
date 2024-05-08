package main

import "fmt"

// Filter takes a slice of any type and a function that defines the filtering criteria.
func Filter[T any](slice []T, criteria func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if criteria(v) {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	// Example usage with a slice of integers
	ints := []int{1, 2, 3, 4, 5}
	even := Filter(ints, func(n int) bool { return n%2 == 0 })
	fmt.Println(even) // Output: [2 4]
	// Example usage with a slice of strings
	strings := []string{"apple", "banana", "cherry", "date"}
	withA := Filter(strings, func(s string) bool { return s[0] == 'a' })
	fmt.Println(withA) // Output: [apple banana]
}