package task1

import "fmt"

// 原地有序数组中的重复项，返回新长度
func removeDuplicates(nums []int) int {
	// 处理空数组情况
	if len(nums) == 0 {
		return 0
	}

	// 慢指针i：记录不重复元素的位置（初始指向第一个元素）
	i := 0

	// 快指针j：遍历整个数组（从第二个元素开始）
	for j := 1; j < len(nums); j++ {
		// 当快慢指针指向的元素不同时
		if nums[j] != nums[i] {
			// 慢指针后移一位
			i++
			// 将快指针元素赋值给慢指针位置（覆盖重复值）
			nums[i] = nums[j]
		}
		// 若元素相同，快指针继续后移（跳过重复项）
	}

	// 新长度为慢指针索引+1（因为索引从0开始）
	return i + 1
}

func removeDuplicatesTest() {
	// 测试用例
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{},
		{5},
		{2, 2, 2, 2},
		{2, 1, 2, 2},
	}

	for _, nums := range testCases {
		// 创建输入副本用于输出展示（避免原数组被修改）
		input := make([]int, len(nums))
		copy(input, nums)

		length := removeDuplicates(nums)
		// 截取有效长度的结果
		result := nums[:length]
		fmt.Printf("输入: %v → 新长度: %d, 结果数组: %v\n", input, length, result)
	}
}
