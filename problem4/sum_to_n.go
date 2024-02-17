/*
Problem 1: Three ways to sum to n using Golang
Author: Anh Nguyen
Date: 14/02/2024
*/

package main

import "fmt"

func main() {
	// Testing the functions
	fmt.Println(sum_to_n_a(680))
	fmt.Println(sum_to_n_b(680))
	fmt.Println(sum_to_n_c(680))

}

func sum_to_n_a(n int) int {
	//Mathematical approach with complexity O(1)
	return n * (n + 1) / 2
}

func sum_to_n_b(n int) int {
	//Using loop approach with complexity: O(n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sum_to_n_c(n int) int {
	//Recursive approach with complexity: O(n)
	if n == 1 {
		return 1
	}
	return n + sum_to_n_c(n-1)
}
