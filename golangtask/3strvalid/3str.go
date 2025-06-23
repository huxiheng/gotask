package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}

	// 括号对应关系 map：右括号 -> 左括号
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char) // 压栈
		case ')', ']', '}':
			if len(stack) == 0 {
				return false // 没有匹配的左括号
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 出栈
			if top != pairs[char] {
				return false // 括号不匹配
			}
		default:
			return false // 非法字符（如果只允许括号）
		}
	}

	return len(stack) == 0 // 栈为空，说明全部匹配
}

func main() {
	tests := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
		"[",
	}

	for _, test := range tests {
		fmt.Printf("isValid(%q) = %v\n", test, isValid(test))
	}
}
