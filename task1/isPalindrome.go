package task1

import "fmt"

func isPalindrome(x int) bool {
	// 特殊情况：
	// 1. 负数不是回文数（因为有负号）
	// 2. 末位是0的非零数不是回文数（如1200，反转后是0021即21，不等于原数）
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedNum := 0
	for x > reversedNum {
		reversedNum = reversedNum*10 + x%10
		x /= 10
	}

	return x == reversedNum || x == reversedNum/10
}

func isPalindromeTest() {
	// 测试用例
	testCases := []int{121, -121, 10, 12321, 0}

	for _, num := range testCases {
		result := isPalindrome(num)
		fmt.Printf("%d 是回文数吗? %v\n", num, result)
	}
}
