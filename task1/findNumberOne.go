package task1

import "fmt"

func findNumberOneMethod(numberArray []int) int {
	countMap := make(map[int]int)

	for _, num := range numberArray {
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return 0
}

func findNumberOneTest() {
	param := [][]int{
		{1, 2, 2, 3, 3, 4, 4},
		{1, 2, 3, 4, 1, 2, 3},
	}
	for _, num := range param {
		result := findNumberOneMethod(num)
		fmt.Printf("原数组：%d---只出现一次的数字：%d\n", num, result)
	}
}
