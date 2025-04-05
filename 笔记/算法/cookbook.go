package book

import (
	"math/rand"
	"strings"
	"time"
)

// 给定⼀个仅包含数字 2-9 的字符串，返回所有它能表示的字⺟组合。给出数字到字⺟的映射如下（与电
// 话按键相同）。注意 1 不对应任何字⺟。  需要用go实现
func letterCombinations(nums string) []string {
	if nums == "" {
		return []string{}
	}
	// 数字到字母的映射表
	mapping := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	// 辅助变量
	var result []string
	var combination string

	// 匿名函数
	var backString func(int)
	backString = func(index int) {
		//    递归出口
		if index == len(nums) {
			// 增加字符
			result = append(result, combination)
			return
		}
		//  字母对应的字符
		num := nums[index : index+1]
		s := mapping[num]
		// 遍历，并且递归调用
		for _, v := range s {
			combination += string(v)
			backString(index + 1)
			combination = combination[:len(combination)-1]
		}
	}
	// 执行
	backString(0)
	//  返回结果
	return result

}

type student struct {
	name string
}

func maxArea(heights []int) int {
	max, left, right := 0, 0, len(heights)
	high, width, temp := 0, 0, 0
	for left < right {

		if heights[left] < heights[right] {
			high = heights[left]
			width = right - left
			left++
		} else {
			high = heights[right]
			width = right - left
			right--
		}
		temp = high * width
		if temp > max {
			max = temp
		}
	}
	return max
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(first *ListNode) *ListNode {
	pre, cur := &ListNode{}, first
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func reversebyindex(first, second *ListNode) *ListNode {
	pre, cur := &ListNode{}, first
	for cur == second {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

// 去重
func removeDuppliactes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	count := 0
	for range m {
		count++
	}
	return count
}

// 找出数组中的index元素，并移动到末尾。返回去除index的数组

func removeElement(nums []int, val int) []int {
	if len(nums) == 0 {
		return nums
	}
	temp := 0
	for i, j := 0, len(nums); i < j; i++ {
		if i != j {
			if nums[j] == val {
				j--
			}
			if nums[i] == val {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		temp = j
	}
	return nums[:temp]
}

// Seed
func removeElement2(nums []int, val int) []int {
	// rand.Seed(2)
	// rand.Seed()
	// New(NewSource(2))
	rand.Seed(time.Now().Unix())
	rand.Int()
	if len(nums) == 0 {
		return nums
	}
	temp := 0
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return nums[:temp]
}

// 找目标子串
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
