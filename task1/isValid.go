package task1

import "fmt"

// 判断括号字符串是否有效
func isValid(s string) bool {
	// 创建一个栈用于存储左括号
	stack := []rune{}

	// 创建括号映射表，右括号对应左括号
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是右括号
		if val, ok := bracketMap[char]; ok {
			// 栈为空或栈顶元素不匹配，返回false
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			// 弹出栈顶元素（匹配成功）
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	// 栈为空说明所有括号都匹配成功
	return len(stack) == 0
}

func isValidTest() {
	// 测试用例
	testCases := []string{
		"()",
		"(][)",   // 无效
		"()[]{}", // 有效
		"(]",     // 无效
		"([)]",   // 无效
		"{[]}",   // 有效
		"",       // 有效（空字符串）
		"(",      // 无效
		")",      // 无效
	}

	for _, s := range testCases {
		result := isValid(s)
		fmt.Printf("字符串: %q → 是否有效: %v\n", s, result)
	}
}
