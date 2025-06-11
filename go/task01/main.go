package main

import (
	"fmt"
	"strconv"
)

/*
 singleNumber 只出现过一次的数字 https://leetcode.cn/problems/single-number/
*/
func singleNumber(nums []int) int {
    mapNum := make(map[int]int)
	for _, num := range nums {
		if val, ok := mapNum[num]; ok {
			mapNum[num] = val + 1
		} else {
			mapNum[num] = 1
		}
	}
	for num, count := range mapNum {
		if count == 1 {
			return num
		}
	}
	return 0
}

// 判断是否是回文数 https://leetcode.cn/problems/palindrome-number/
func isPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	for i := 0; i < len(numStr)/2; i++ {
		if numStr[i] != numStr[len(numStr)-i-1] {
			return false
		}
	}
	return true
}

// 有效括号 https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
	for _, char := range s {
		// 如果是左括号，则入栈
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			// 如果栈顶元素不匹配，则返回 false
			if stack[len(stack)-1] != mapping[char] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	// 如果字符串数组为空，则返回空字符串
    if len(strs) == 0 {
        return ""
    }
	// 初始化前缀为第一个字符串
    prefix := strs[0]
    for _, str := range strs[1:] {
        i := 0
		// 比较前缀和当前字符串, 且字符串长度不超过前缀长度
		for i < len(prefix) && i < len(str) && prefix[i] == str[i] {
			i++
		}
		prefix = prefix[:i]
		// 如果前缀为空，则返回空字符串
		if prefix == "" {
			return ""
		}
    }
	return prefix
}

func main() {
	fmt.Println("task01")
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))  // 测试只出现一次的数字
	fmt.Println(isPalindrome(121))  // 测试回文数
	fmt.Println(isValid("()[]{}"))  // 测试有效括号
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // 测试最长公共前缀
}
