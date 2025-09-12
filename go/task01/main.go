package main

import (
	"fmt"
	"sort"
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

// 26.删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    i := 0
	// 遍历数组，如果当前元素与前一个元素不相同，则将当前元素赋值给nums[i]
    for j := 1; j < len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }
    // 截断数组，移除重复项
    nums = nums[:i+1]
    return i + 1
}

// 66 加一 https://leetcode.cn/problems/plus-one/
func plusOne(digits []int) []int {
	// 从后往前遍历数组, 如果当前元素小于9, 则加一, 并返回
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果当前元素小于9，则加一，并返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 如果当前元素等于9，则将当前元素赋值为0
		digits[i] = 0
	}
	// 如果遍历完数组, 则说明数组所有元素都为9, 则需要扩容, 在数组最前面插入1
	digits = append([]int{1}, digits...)
	return digits
}

// 56. 合并区间 https://leetcode.cn/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return intervals
    }
    
    // 按起始位置排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    result := [][]int{intervals[0]}
	// 遍历区间, 如果当前区间的起始位置小于等于前一个区间的结束位置, 则合并, 否则直接添加到结果数组中
    for i := 1; i < len(intervals); i++ {
        last := result[len(result)-1]
		// 如果当前区间的起始位置小于等于前一个区间的结束位置, 则合并, 否则直接添加到结果数组中
        if intervals[i][0] <= last[1] {
            if intervals[i][1] > last[1] {
                last[1] = intervals[i][1]
            }
        } else {
            result = append(result, intervals[i])
        }
    }
    return result
}

// 1. 两数之和 https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	mapNum := make(map[int]int)
	for i, num := range nums {
		if val, ok := mapNum[target-num]; ok {
			return []int{val, i}
		}
		mapNum[num] = i
	}
	return nil
}

func main() {
	fmt.Println("task01")
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))  // 测试只出现一次的数字
	fmt.Println(isPalindrome(121))  // 测试回文数
	fmt.Println(isValid("()[]{}"))  // 测试有效括号
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // 测试最长公共前缀
	fmt.Println(removeDuplicates([]int{1, 1, 2})) // 测试删除重复项
	fmt.Println(plusOne([]int{1, 2, 3})) // 测试加一
	fmt.Println(merge([][]int{{1, 4}, {2, 5}, {3, 6}})) // 测试合并区间
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // 测试两数之和
}
