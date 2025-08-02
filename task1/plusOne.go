package task1

import "fmt"

// 大整数加1
func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始处理
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果加1后小于10，说明没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}

		// 有进位，当前位置为0，继续处理前一位
		digits[i] = 0
	}

	// 如果所有位都有进位（如999 -> 1000）
	// 需要创建新数组，首位为1，后面跟n个0
	return append([]int{1}, digits...)
}

func plusOneTest() {
	// 测试用例
	testCases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
		{1, 9, 9},
		{0},
	}

	for _, digits := range testCases {
		// 为了不修改原数组，创建副本进行测试
		input := make([]int, len(digits))
		copy(input, digits)

		result := plusOne(digits)
		fmt.Printf("输入: %v → 加1后: %v\n", input, result)
	}
}
