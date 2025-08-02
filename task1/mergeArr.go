package task1

import (
	"fmt"
	"sort"
)

// 合并区间
func merge(intervals [][]int) [][]int {
	// 处理空输入
	if len(intervals) == 0 {
		return nil
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用于存储合并后的区间
	merged := [][]int{intervals[0]}

	// 遍历排序后的区间
	for i := 1; i < len(intervals); i++ {
		// 获取合并区间的最后一个区间
		last := &merged[len(merged)-1]

		// 当前区间的起始位置小于等于最后一个区间的结束位置，说明有重叠
		if intervals[i][0] <= (*last)[1] {
			// 合并区间：结束位置取两个区间的最大值
			if intervals[i][1] > (*last)[1] {
				(*last)[1] = intervals[i][1]
			}
		} else {
			// 无重叠，直接添加到合并区间
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

func mergeArrTest() {
	// 测试用例
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {2, 3}},
		{{1, 2}, {3, 4}, {5, 6}},
		{{1, 5}},
		{},
	}

	for _, intervals := range testCases {
		result := merge(intervals)
		fmt.Printf("输入: %v → 合并后: %v\n", intervals, result)
	}
}
