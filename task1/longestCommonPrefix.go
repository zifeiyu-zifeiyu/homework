package task1

import "fmt"

// 查找字符串数组的最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 处理空数组情况
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准
	for i := 0; i < len(strs[0]); i++ {
		// 取基准字符串的第i个字符
		char := strs[0][i]

		// 与其他字符串的第i个字符比较
		for j := 1; j < len(strs); j++ {
			// 如果超出某个字符串长度，或字符不匹配，返回当前前缀
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	// 第一个字符串就是最长公共前缀
	return strs[0]
}

func longestCommonPrefixTest() {
	// 测试用例
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"abc", "ab", "a"},
		{"same", "same", "same"},
		{""},
		{"a"},
	}

	for _, strs := range testCases {
		result := longestCommonPrefix(strs)
		fmt.Printf("输入: %v → 最长公共前缀: %q\n", strs, result)
	}
}
