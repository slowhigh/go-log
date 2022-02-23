package main

// Given an integer n, return the number of trailing zeroes in n!.

// Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

// Example 1:
// Input: n = 3
// Output: 0
// Explanation: 3! = 6, no trailing zero.

// Example 2:
// Input: n = 5
// Output: 1
// Explanation: 5! = 120, one trailing zero.

// Example 3:
// Input: n = 0
// Output: 0

// Constraints:
// 0 <= n <= 104

// Follow up: Could you write a solution that works in logarithmic time complexity?

import (
	"fmt"
)

func trailingZeroes(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		target := i
		for ; target%5 == 0; {
			target /= 5
			count++
		}
	}

	return count
}

// 2와 5 묶음의 개수가 뒤에 따라오는 0의 개수를 의미한다.
// 그런데 5의 개수는 항상 2의 개수보다 모자른다. 따라서 5의 개수를 구하면 된다.
// 5!일 때 1*2*3*4*5 이므로 5는 1개 이다.
// 5는 1개, 10도 1개, 15도 1개,,,,,25는 2개, 30도 2개,,,,,125는 3개, 130는 3개

// 따라서 130! 일 때 ( 5배수 개수 + 25배수 개수 + 125배수 개수)
//                 =(   130/5    +    130/25  +   130/125  )
//                 =(   130/5    +    130/5*5 +  130/5*5*5 )

func trailingZeroes2(n int) int {
	zeroes := 0
	for n != 0 {
		zeroes += n/5
		n /= 5
	}

	return zeroes
}

func main() {
	fmt.Println(trailingZeroes(130))
	fmt.Println(trailingZeroes2(130))
}
