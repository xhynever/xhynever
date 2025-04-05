package book

import (
	"sort"
	"strconv"
	"unicode"
)

// 给定参数n
// 从1到n会有n个整数 1，2，3，...n
// 这n个数字共有n!种排列 按大小顺序升序列出所有排列情况
// 并一一标记
// 当n = 3 时，所有排列如下
// "123","132","213","231","312","321"
// 给定n和k返回第n个排列

// factorial calculates the factorial of a given number n.
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// getPermutation returns the kth permutation sequence of numbers from 1 to n.
func getPermutation(n int, k int) string {
	k-- // Convert to 0-based index
	var result []byte
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	for n > 0 {
		n--
		// Determine the factorial of (n-1)
		fact := factorial(n)
		// Find the index of current digit
		index := k / fact
		k %= fact
		// Append the current digit to result and remove it from nums
		result = append(result, byte(nums[index]+'0'))
		nums = append(nums[:index], nums[index+1:]...)
	}

	return string(result)
}

// 有N个正整数组成的一个序列
// 给定一个整数sum
// 求长度最长的的连续子序列使他们的和等于sum
// 返回次子序列的长度
// 如果没有满足要求的序列 返回-1
// 要找到一个和为给定值sum的最长连续子序列，可以使用滑动窗口（或双指针）技术。这个方法的时间复杂度是O(n)，其中n是数组的长度。滑动窗口的思想是维护一个窗口，该窗口内的元素之和不超过sum，并且尝试通过调整窗口大小来找到和恰好等于sum的最大窗口。
func longestSubarrayWithSum(nums []int, target int) int {
	maxLength := -1
	windowStart, windowSum := 0, 0

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		windowSum += nums[windowEnd] // Add the next element to the window.

		// Shrink the window as small as possible until the 'windowSum' is smaller than or equal to 'target'
		for windowStart <= windowEnd && windowSum > target {
			windowSum -= nums[windowStart]
			windowStart++
		}

		// Update the maximum length if the current window's sum equals to 'target'.
		if windowSum == target {
			if maxLength < windowEnd-windowStart+1 {
				maxLength = windowEnd - windowStart + 1
			}
		}
	}

	return maxLength
}

// 给定一个正整数数组
// 检查数组中是否存在满足规则的数组组合
// 规则：
// A = B + 2C

// 使用哈希表（在Go语言中是map）来存储已经遍历过的元素及其索引，这样可以在O(1)时间内检查某个值是否存在于哈希表中。通过这种方法，我们可以将问题转换为两层循环，对于每一个元素 B 和 C，我们计算出 A 应该是什么，并检查它是否存在于哈希表中。
// checkCombination checks if there exists a combination of elements in the array that satisfies A = B + 2C.
func checkCombination(nums []int) []int {
	// Create a map to store the elements and their indices.
	elementMap := make(map[int]bool)

	for _, num := range nums {
		elementMap[num] = true
	}
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k1 == k2 {
				continue
			}
			a := v1 + v2*2
			if elementMap[a] && a != k1 && a != k2 {

				return []int{a, v1, v2}
			}
		}
	}

	return []int{}
}

// int n = in.nextInt();
// int[] dp = new int[n+1];
// dp[0] = 1;
// dp[1] = 0;
// for(int i=2; i<=n; i++){
// if(i % 2 == 0){
// dp[i] = dp[i/2]+1;
// }else{
// dp[i] = Math.min(dp[i-1] + 1, dp[(i+1)/2] + 2);
// }
// }

// 小明从糖果盒中随意抓一把糖果
// 每次小明会取出一半的糖果分给同学们
// 当糖果不能平均分配时
// 小明可以从糖果盒中(假设盒中糖果足够)取出一个或放回一个糖果
// 小明至少需要多少次(取出放回和平均分配均记一次)能将手中糖果分至只剩一颗

func dpCandy(num int) int {
	dp := []int{}
	dp[0] = 1
	dp[1] = 0
	dp[2] = 1
	for i := 2; i <= num; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2] + 1
		} else {
			dp[i] = min(dp[i-1]+1, dp[(i+1)/2]+2)
		}

	}
	return dp[num]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 输入一个字符串仅包含大小写字母和数字
// 求字符串中包含的最长的非严格递增连续数字序列长度
// 比如：
// 12234属于非严格递增数字序列

// longestNonDecreasingSequence finds the length of the longest non-decreasing sequence of digits in a string.
func longestNonDecreasingSequence(s string) int {
	maxLen, currentLen := 0, 0
	var prevDigit rune = -1 // Initialize with an invalid digit value.

	for _, char := range s {
		if char >= '0' && char <= '9' { // Check if the character is a digit.
			digit := char - '0'
			if currentLen == 0 || digit >= prevDigit {
				currentLen++
				if currentLen > maxLen {
					maxLen = currentLen
				}
			} else {
				currentLen = 1 // Reset the current sequence length.
			}
			prevDigit = digit
		} else {
			currentLen = 0 // Reset for non-digit characters.
			prevDigit = -1 // Reset the previous digit tracker.
		}
	}

	return maxLen
}

// 翻转字符串
func reverseWordsInSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 为了充分发挥GPU算力，
// 需要尽可能多的将任务交给GPU执行，
// 现在有一个任务数组，
// 数组元素表示在这1s内新增的任务个数，
// 且每秒都有新增任务，
// 假设GPU最多一次执行n个任务，
// 一次执行耗时1s，
// 在保证GPU不空闲的情况下，最少需要多长时间执行完成。

func CalculateMinTime(tasks []int, n int) int {
	if len(tasks) == 0 || n <= 0 {
		return 0
	}

	timeSpent := 0  // Total time spent on processing tasks
	tasksQueue := 0 // Number of tasks waiting in queue
	for v := range tasks {
		timeSpent++
		tasksQueue += v
		if tasksQueue > n {
			tasksQueue -= n
		} else {
			tasksQueue = 0
		}

	}
	if tasksQueue%n > 0 {
		timeSpent += tasksQueue/n + 1
	} else {
		timeSpent += tasksQueue / n
	}
	return timeSpent
}

// 公司用一个字符串来标识员工的出勤信息
// absent: 缺勤
// late: 迟到
// leaveearly:早退
// present: 正常上班
// 现需根据员工出勤信息,判断本次是否能获得出勤奖,
// 能获得出勤奖的条件如下：

// 缺勤不超过1次
// 没有连续的迟到/早退
// 任意连续7次考勤，缺勤/迟到/早退，不超过3次

// CanGetAward checks if the employee can get the attendance award based on their attendance record.
func CanGetAward(record string) bool {
	absentCount := 0
	consecutiveLateOrLeaveEarly := 0
	violationCount := 0
	slidingWindow := make([]int, 0, 7) // To keep track of violations in a sliding window of size 7

	for i := 0; i < len(record); {
		switch {
		case record[i:i+6] == "absent":
			absentCount++
			violationCount++
			slidingWindow = append(slidingWindow, 1)
			i += 6
		case record[i:i+4] == "late":
			consecutiveLateOrLeaveEarly++
			violationCount++
			slidingWindow = append(slidingWindow, 1)
			i += 4
		case record[i:i+9] == "leaveearly":
			consecutiveLateOrLeaveEarly++
			violationCount++
			slidingWindow = append(slidingWindow, 1)
			i += 9
		case record[i:i+7] == "present":
			consecutiveLateOrLeaveEarly = 0
			slidingWindow = append(slidingWindow, 0)
			i += 7
		default:
			return false // Invalid record format
		}

		// Check for conditions to fail
		if absentCount > 1 || consecutiveLateOrLeaveEarly >= 2 {
			return false
		}

		// Maintain the sliding window size of 7
		if len(slidingWindow) > 7 {
			violationCount -= slidingWindow[0]
			slidingWindow = slidingWindow[1:]
		}

		// Check violation count in any continuous 7 records
		if violationCount > 3 {
			return false
		}
	}

	return true
}

// 有一种简易压缩算法：针对全部为小写英文字母组成的字符串，
// 将其中连续超过两个相同字母的部分压缩为连续个数加该字母
// 其他部分保持原样不变.
// 例如字符串aaabbccccd 经过压缩变成字符串 3abb4cd
// 请您编写解压函数,根据输入的字符串,
// 判断其是否为合法压缩过的字符串
// 若输入合法则输出解压缩后的字符串
// 否则输出字符串!error来报告错误
func Decompress(compressed string) string {
	var result []rune
	i := 0
	runes := []rune(compressed)

	for i < len(runes) {
		if unicode.IsDigit(runes[i]) {
			// Read the full number
			numStart := i
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				i++
			}
			count, err := strconv.Atoi(string(runes[numStart:i]))
			if err != nil || i >= len(runes) || !unicode.IsLower(runes[i]) {
				return "!error"
			}
			letter := runes[i]
			for j := 0; j < count; j++ {
				result = append(result, letter)
			}
			i++
		} else if unicode.IsLower(runes[i]) {
			result = append(result, runes[i])
			i++
		} else {
			return "!error"
		}
	}

	return string(result)
}

func shoot(id []int, score []int) []int {
	var res []int
	m := [][]int{}
	for i := 0; i < len(id); i++ {
		m[i] = append(m[i], score[i])
		if i >= len(score) {
			m[i] = append(m[i], 0)
		}
	}
	for k, score := range m {
		if len(score) < 3 {
			break
		}
		temp := score
		sort.Ints(temp)
		sum := 0
		for i := 0; i < 3; i++ {
			sum += temp[i]
		}
		res[k] = sum
	}
	return res
}

func minCars2(nums []int) {
	var count int
	var temp int
	for _, v := range nums {
		if v == 1 {
			temp++
		} else {
			count += temp / 3
			if temp/3 != 0 {
				count++
			}
		}

	}
}

// maxSpending 函数计算在给定预算和必须购买的商品数量下可以花费的最大金额
func maxSpending(budget int, prices []int, mustBuy int) int {
	n := len(prices)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, budget+1)
	}

	// 填充dp数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= budget; j++ {
			if prices[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max2(dp[i-1][j], dp[i-1][j-prices[i-1]]+prices[i-1])
			}
		}
	}

	// 回溯找出恰好选择了mustBuy个商品的情况
	result := 0
	for j := budget; j >= 0; j-- {
		count := 0
		for i := n; i > 0; i-- {
			if dp[i][j] != dp[i-1][j] {
				count++
			}
		}
		if count == mustBuy && dp[n][j] > result {
			result = dp[n][j]
		}
	}
	return result
}

// 启动协程组，监听协程组任务，如果报错。退出协程组。

// func worker(ctx context.Context, wg *sync.WaitGroup, id int, errChan chan<- error, numChan chan int) {
// 	defer wg.Done()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// 当收到取消信号时退出循环
// 			// fmt.Printf("Worker %d: received cancellation signal, stopping...\n", id)
// 			return
// 		default:
// 			// 模拟工作
// 			// temp := <-numChan
// 			// fmt.Printf("Worker %d: working...\n", id)
// 			// fmt.Println(temp%3 + 1)
// 			// printNum(id, numChan)

// 			if id == 2 {
// 				errChan <- errors.New(fmt.Sprintf("Worker %d encountered an error", id))
// 				return
// 			}

// 		}
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithCancel(context.Background())
// 	errChan := make(chan error, 1)
// 	numChan := make(chan int, 1)

// 	// 启动多个worker goroutines
// 	numWorkers := 3
// 	wg.Add(numWorkers)
// 	maxnum := 3 * numWorkers
// 	for i := 1; i <= numWorkers; i++ {
// 		go worker(ctx, &wg, maxnum, errChan, numChan)
// 	}
// 	for i := 1; i < maxnum; i++ {
// 		numChan <- i
// 	}

// 	// 监听错误通道，一旦有错误发生就取消所有goroutines
// 	go func() {
// 		if err := <-errChan; err != nil {
// 			fmt.Println(err)
// 			cancel() // 发送取消信号给所有goroutines
// 		}
// 	}()

// 	wg.Wait()
// }


// 找出N个数中，和大于等于X的数组。
func countSubarraysWithSumGreaterThanX(nums []int, x int) int {
	count := 0
	left := 0
	currentSum := 0

	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]

		// 当前窗口内元素之和大于等于x时，移动左边界缩小窗口，并累加满足条件的子数组数量
		for currentSum >= x {
			count += len(nums) - right // 当前right到数组末尾的所有连续子数组都满足条件
			currentSum -= nums[left]
			left++
		}
	}

	return count
}






// 某公司组织一场公开招聘活动，假设由于人数和场地的限制，
// 每人每次面试的时长不等，并已经安排给定，
// 用(S1,E1)、(S2,E2)、(Sj,Ej)...(Si < Ei，均为非负整数)表示每场面试的开始和结束时间。
// 面试采用一对一的方式，即一名面试官同时只能面试一名应试者，
// 一名面试官完成一次面试后可以立即进行下一场面试，且每个面试官的面试人次不超过m。
// 为了支撑招聘活动高效顺利进行，请你计算至少需要多少名面试官。

type event struct{
	time int
	start bool

}


// 创建队列，入栈出栈

// package main

// import (
// 	"fmt"
// )

// type Node struct {
// 	Val  int
// 	Next *Node
// }

// func top() {
// 	if root.Next == nil {
// 		fmt.Println("error")
// 		return
// 	}
// 	fmt.Println(root.Next.Val)
// }

// func pop() {
// 	if root.Next == nil {
// 		fmt.Println("error")
// 		return
// 	}
// 	fmt.Println(root.Next.Val)
// 	root.Next = root.Next.Next
// }

// func push(x int) {
// 	node := &Node{Val: x, Next: nil}
// 	node.Next = root.Next
// 	root.Next = node
// }

// var root *Node

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	root = &Node{Val: -1, Next: nil}
// 	for n >= 0 {
// 		var met string
// 		var x int
//         fmt.Scan(&met)
// 		switch met {
// 		case "top":
// 			top()
// 		case "pop":
// 			pop()
// 		case "push":
//             fmt.Scan(&x)
// 			push(x)
// 		}
// 		n--
// 	}
// }


// 用interface实现栈
// package main

// import (
// 	"fmt"
// )

// // 定义一个栈结构体
// type Stack struct {
// 	items []interface{}
// }

// // Push 方法用于向栈中添加元素
// func (s *Stack) Push(item interface{}) {
// 	s.items = append(s.items, item)
// }

// // Pop 方法用于从栈顶移除元素并返回它
// func (s *Stack) Pop() interface{} {
// 	if len(s.items) == 0 {
// 		fmt.Println("error")
// 		return nil
// 	}
// 	item := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return item
// }

// // Top 方法返回栈顶元素但不移除它
// func (s *Stack) Top() interface{} {
// 	if len(s.items) == 0 {
// 		fmt.Println("error")
// 		return nil
// 	}
// 	return s.items[len(s.items)-1]
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	stack := &Stack{}

// 	for i := 0; i < n; i++ {
// 		var command string
// 		fmt.Scan(&command)
// 		switch command {
// 		case "pop":
// 			if val := stack.Pop(); val != nil {
// 				fmt.Println(val)
// 			}
// 		case "top":
// 			if val := stack.Top(); val != nil {
// 				fmt.Println(val)
// 			}
// 		default:
// 			// 对于 push 操作，读取后续的整数值
// 			var x int
// 			fmt.Scanf("%s %d", &command, &x)
// 			stack.Push(x)
// 		}
// 	}
// }