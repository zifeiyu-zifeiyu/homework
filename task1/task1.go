package task1

import "fmt"

func numCheck(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}

		numMap[num] = i
	}
	return nil
}

func Task1() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}
	for _, tc := range testCases {
		result := numCheck(tc.nums, tc.target)
		fmt.Printf("nums = %v, target = %d → 结果: %v\n", tc.nums, tc.target, result)
	}
}
