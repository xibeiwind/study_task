package taskone

import "sort"

// 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

// 有效的括号
func isValid(s string) bool {
	// 定义括号映射
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 使用切片模拟栈
	var stack []rune

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 左括号入栈
			stack = append(stack, char)
		case ')', '}', ']':
			// 右括号检查
			if len(stack) == 0 {
				return false // 栈空，无匹配左括号
			}
			// 弹出栈顶
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != pairs[char] {
				return false // 不匹配
			}
		}
	}

	// 栈空则有效
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 以第一个字符串为基准
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// 逐个字符比较，直到不匹配或越界
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = prefix[:j] // 截断
		if prefix == "" {   // 提前退出
			return ""
		}
	}
	return prefix
}

// 加一
func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始处理
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果加1后小于10，没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}

		// 如果等于10，当前位设为0，继续处理前一位
		digits[i] = 0
	}

	// 如果所有位都处理完还有进位，需要在数组前面插入1
	return append([]int{1}, digits...)
}

// 移除重复元素
func removeDuplicates(nums []int) int {
	s, l := 1, len(nums)
	for f := 1; f < l; f++ {
		if nums[f] != nums[f-1] {
			nums[s] = nums[f]
			s++
		}
	}
	return s
}

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 1. 按照区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化结果集
	result := [][]int{intervals[0]}

	// 3. 遍历并合并区间
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1] // 结果集中最后一个区间
		current := intervals[i]       // 当前区间

		// 如果当前区间与最后一个区间重叠
		if current[0] <= last[1] {
			// 合并区间：取结束位置的较大值
			if current[1] > last[1] {
				last[1] = current[1] //last修改后即修改了result
			}
		} else {
			// 不重叠，直接添加到结果集
			result = append(result, current)
		}
	}

	return result
}

// 两数之和
func twoSum(nums []int, target int) []int {
	// value -> index 的映射
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		need := target - v
		if j, ok := seen[need]; ok { // 伙伴已出现过
			return []int{j, i}
		}
		// 先查后插，保证不会重复使用同一元素
		seen[v] = i
	}
	return nil // 题目保证有解，这里不会走到
}
