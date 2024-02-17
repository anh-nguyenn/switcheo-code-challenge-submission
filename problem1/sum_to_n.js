/*
Problem 1: Three ways to sum to n using javascript
Author: Anh Nguyen
Date: 14/02/2024
*/

// Iterative approach with complexity: O(n)
function sum_to_n_a(n) {
  let sum = 0;
  for (let i = 1; i <= n; i++) {
    sum += i;
  }
  return sum;
}

// Mathematical approach (most efficient) with complexity: O(1)
function sum_to_n_b(n) {
  return (n * (n + 1)) / 2;
}

// Recursive approach with complexity: O(n)
function sum_to_n_c(n) {
  if (n === 1) {
    return 1;
  }
  return n + sum_to_n_c(n - 1);
}

// Testing the functions
console.log(sum_to_n_a(15));
console.log(sum_to_n_b(15));
console.log(sum_to_n_c(15));
