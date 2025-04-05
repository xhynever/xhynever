package book

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 随机数

func num() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	sz := make([]int, 505)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		sz[n] = 1
	}
	for i := 1; i <= 500; i++ {
		if sz[i] == 1 {
			fmt.Println(i)
		}
	}

}

// 字符串最后单词的长度
func LastWordTest() {

	reader := bufio.NewReader(os.Stdin)
	char, _ := reader.ReadString('\n')
	char1 := strings.Fields(char)
	fmt.Println(len(char1[len(char1)-1]))

}
func LastWord(s string) int {
	words := strings.Fields(s)
	n := len(words)
	if n == 0 {
		return 0
	}
	LastWord := words[n-1]
	return len(LastWord)
}

//     给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换。
// 输入描述:
// 输入为两行，第一行为参数K，第二行为字符串S。
// 输出描述:
// 输出转换后的字符串。
// 示例1
// 输入
// 3
// 12abc-abCABc-4aB@
// 输出
// 12abc-abc-ABC-4aB-@
// 说明
// 子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每3个字符一组为abC、ABc、4aB、@，abC中小写字母较多，转换为abc，ABc中大写字母较多，转换为ABC，4aB中大小写字母都为1个，不做转换，@中没有字母，连起来即12abc-abc-ABC-4aB-@
// 示例2
// 输入
// 12
// 12abc-abCABc-4aB@
// 输出
// 12abc-abCABc4aB@
// 说明
// 子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每12个字符一组为abCABc4aB@，这个子串中大小写字母都为4个，不做转换，连起来即12abc-abCABc4aB@

// countUpperAndLower 统计字符串中大写字母和小写字母的数量
func TestcountUpperAndLower() {
	var k int
	fmt.Scanln(&k)
	var s string
	fmt.Scanln(&s)
	fmt.Println(transformString(s, k))
}
func countUpperAndLower(s string) (int, int) {
	upnumber, lownumber := 0, 0
	for char := range s {
		if 'A' <= char && 'Z' >= char {
			upnumber++
		} else if 'a' <= char && 'z' >= char {
			lownumber++
		}
	}
	return upnumber, lownumber
}

func transformString(s string, k int) string {
	// 拆分
	parts := strings.Split(s, "-")
	var result []string
	// 放入第一个不变的单词
	result = append(result, parts[0])

	// 遍历余下单词。若k的值大于单个单词，temp,tempindex
	for _, part := range parts[1:] {
		subParts := []string{}
		// 将k个字符打包为一组,边界条件单独一组

		for i := 0; i < len(part); i += k {

			end := i + k
			if end > len(part) {
				end = len(part)

			} else {

				subParts = append(subParts, part[i:end])
			}

		}

		for _, subPart := range subParts {
			upperCount, lowerCount := countUpperAndLower(subPart)
			if upperCount > lowerCount {
				subPart = strings.ToUpper(subPart)
			} else if lowerCount > upperCount {
				subPart = strings.ToLower(subPart)
			}
			result = append(result, subPart)
		}
	}
	return strings.Join(result, "-")
}

// 组成最大数  思路通常是对给定的一组数字进行排序，自定义排序规则让数字按照能够组成最大数的顺序排列，然后将这些数字拼接起来形成最终的最大数。

// largestNumber函数接收一个整数切片，返回由这些整数组成的最大数字字符串
func largestNumber(nums []int) string {
	// 将整数切片转换为字符串切片，方便后续操作
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	//  自定义排序规则，按照能够组成最大数的方式排序
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	// 如果排序后的第一个字符串是"0"，意味着最大数就是0，直接返回"0"
	if strNums[0] == "0" {
		return "0"
	}

	// 将排序后的字符串切片拼接成最终的最大数字字符串
	return strings.Join(strNums, "")
}

func largestNumberTest() {
	nums := []int{3, 30, 34, 5, 9}
	result := largestNumber(nums)
	fmt.Println(result)
}

// 题目描述
// 为了充分发挥GPU算力，
// 需要尽可能多的将任务交给GPU执行，
// 现在有一个任务数组，
// 数组元素表示在这1s内新增的任务个数，
// 且每秒都有新增任务，
// 假设GPU最多一次执行n个任务，
// 一次执行耗时1s，
// 在保证GPU不空闲的情况下，最少需要多长时间执行完成。

// 输入描述
// 第一个参数为GPU最多执行的任务个数，取值范围1 ~ 10000
// 第二个参数为任务数组的长度，取值范围1 ~ 10000
// 第三个参数为任务数组，数字范围1 ~ 10000

// 输出描述
// 执行完所有任务需要多少秒

// 示例一
// 输入
// 3
// 5
// 1 2 3 4 5
func minTimeToFinishTasks(maxTasks int, taskLen int, tasks []int) int {
	totalTasks := 0
	for _, task := range tasks {
		totalTasks += task
	}
	time := 0
	for totalTasks > 0 {
		if totalTasks <= maxTasks {
			time++
			totalTasks = 0
		} else {
			totalTasks -= maxTasks
			time++
		}
	}
	return time
}

func minTimeToFinishTasksTest() {
	var maxTasks, taskLen int
	fmt.Scan(&maxTasks)
	fmt.Scan(&taskLen)
	tasks := make([]int, taskLen)
	for i := 0; i < taskLen; i++ {
		fmt.Scan(&tasks[i])
	}
	result := minTimeToFinishTasks(maxTasks, taskLen, tasks)
	fmt.Println(result)
}

// 猴子爬山
// 题目描述
// 一天一只顽猴想要从山脚爬到山顶，
// 途中经过一个有n个台阶的阶梯，
// 但是这个猴子有个习惯，每一次只跳1步或3步
// 试问？猴子通过这个阶梯有多少种不同的跳跃方式

func climbStairs(n int) int {
	//边界情况处理
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	// 动态规划数组初始化及递推逻辑：
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	// 开始动态规划
	for i := 4; i <= n; i++ {
		// 核心代码
		dp[i] = dp[i-1] + dp[i-3]
	}
	return dp[n]
}

// 身高排序
// 题目描述
// 小明今年升学到了小学一年级，
// 来到新班级后，发现其他小朋友身高参差不齐，
// 然后就想基于各小朋友和自己的身高差，对他们进行排序，
// 请帮他实现排序

// 输入描述
// 第一行为正整数H和N
// 0 < H < 200 为小明的身高
// 0 < N < 50 为新班级其他小朋友个数
// 第二行为N个正整数
// H1 ~ Hn分别是其他小朋友的身高
// 取值范围0 < Hi < 200
// 且N个正整数各不相同

// 输出描述
// 输出排序结果，各正整数以空格分割
// 和小明身高差绝对值最小的小朋友排在前面
// 和小明身高差绝对值最大的小朋友排在后面
// 如果两个小朋友和小明身高差一样
// 则个子较小的小朋友排在前面

type Child struct {
	Height int
	Diff   int
}

func HeightSorting() {
	// 处理输入
	var H, N int
	fmt.Scan(&H, &N)
	// 定制化 child
	children := make([]Child, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&children[i].Height)
		children[i].Diff = abs(children[i].Height - H)
	}
	//   排序
	sort.Slice(children, func(i, j int) bool {
		if children[i].Diff == children[j].Diff {
			return children[i].Height < children[j].Height
		}
		return children[i].Diff < children[j].Diff
	})
	//    遍历输出结果
	for _, child := range children {
		fmt.Printf("%d ", child.Height)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 单词反转
// 题目描述
// 输入一个英文文章片段，
// 翻转指定区域的单词顺序，
// 标点符号和普通字母一样处理，
// 例如输入字符串

func reverseWordsInRange(s string, start, end int) string {
	words := strings.Fields(s)
	if start >= len(words) || end >= len(words) || start > end {
		return s
	}

	// 反转指定范围内的单词
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseWordsInRangeTest() {
	var input, startStr, endStr string
	fmt.Scanln(&input)
	fmt.Scanln(&startStr)
	fmt.Scanln(&endStr)
	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)
	result := reverseWordsInRange(input, start, end)
	fmt.Println(result)
}

func maxSpend(M []int, R int) int {
	maxCost := -1
	for i := 0; i < len(M)-2; i++ {
		for j := i + 1; j < len(M)-1; j++ {
			for k := j + 1; k < len(M); k++ {
				sum := M[i] + M[j] + M[k]
				if sum <= R && sum > maxCost {
					maxCost = sum
				}
			}
		}
	}
	return maxCost
}

func maxSpendTest() {
	var n int
	fmt.Scan(&n)
	M := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&M[i])
	}
	var R int
	fmt.Scan(&R)
	result := maxSpend(M, R)
	fmt.Println(result)
}

// 给定两个字符串，
// 从字符串2中找出字符串1中的所有字符，
// 去重并按照ASCII码值从小到大排列。

// 输入描述
// 字符范围满足ASCII编码要求，
// 输入字符串1长度不超过1024，
// 字符串2长度不超过100。

// 输出描述
// 按照ASCII由小到大排序

// 示例一
// 输入
// bach
// bbaaccddfg
func findAndSort(s1, s2 string) string {
	// 使用map记录出现过的字符，起到去重作用
	exist := make(map[rune]bool)
	for _, char := range s1 {
		exist[char] = true
	}
	var result []rune
	for _, char := range s2 {
		if exist[char] {
			result = append(result, char)
			delete(exist, char)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return string(result)
}

func findAndSortTest() {
	var s1, s2 string
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	output := findAndSort(s1, s2)
	fmt.Println(output)
}

// 最小叶子节点
// 题目描述
// 二叉树也可以用数组来存储，
// 给定一个数组，树的根节点的值储存在下标1，
// 对于储存在下标n的节点，他的左子节点和右子节点分别储存在下标2*n和2*n+1，
// 并且我们用-1代表一个节点为空，
// 给定一个数组存储的二叉树，
// 试求从根节点到最小的叶子节点的路径，
// 路径由节点的值组成。

// 输入描述
// 输入一行为数组的内容，
// 数组的每个元素都是正整数，元素间用空格分割，
// 注意第一个元素即为根节点的值，
// 即数组的第n元素对应下标n，
// 下标0在树的表示中没有使用，所以我们省略了，
// 输入的树最多为7层。

// 输出描述
// 输出从根节点到最小叶子节点的路径上各个节点的值，
// 由空格分割，
// 用例保证最小叶子节点只有一个。

// 示例一
// 输入
// 3 5 7 -1 -1 2 4
// 输出
// 3 7 2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr []string) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]
		leftValStr := arr[i]
		i++
		if leftValStr != "-1" {
			leftVal, _ := strconv.Atoi(leftValStr)
			leftNode := &TreeNode{Val: leftVal}
			node.Left = leftNode
			queue = append(queue, leftNode)
		}
		if i < len(arr) {
			rightValStr := arr[i]
			i++
			if rightValStr != "-1" {
				rightVal, _ := strconv.Atoi(rightValStr)
				rightNode := &TreeNode{Val: rightVal}
				node.Right = rightNode
				queue = append(queue, rightNode)
			}
		}
	}
	return root
}

func findMinLeafPath(root *TreeNode) []int {
	minLeafVal := math.MaxInt
	var minLeafPath []int
	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		newPath := append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			if node.Val < minLeafVal {
				minLeafVal = node.Val
				minLeafPath = newPath
			}
			return
		}
		dfs(node.Left, newPath)
		dfs(node.Right, newPath)
	}
	dfs(root, []int{})
	return minLeafPath
}

func findMinLeafPathTest() {
	var input string
	fmt.Scanln(&input)
	arr := strings.Fields(input)
	root := buildTree(arr)
	path := findMinLeafPath(root)
	for i, val := range path {
		if i > 0 {
			fmt.Printf(" %d", val)
		} else {
			fmt.Printf("%d", val)
		}
	}
}

// 快递货车
// 题目描述
// 一辆运送快递的货车，
// 运送的快递放在大小不等的长方体快递盒中，
// 为了能够装载更多的快递同时不能让货车超载，
// 需要计算最多能装多少个快递。
// 注：快递的体积不受限制。
// 快递数最多1000个，货车载重最大50000。

// 输入描述
// 第一行输入每个快递的重量
// 用英文逗号隔开
// 如 5,10,2,11
// 第二行输入货车的载重量
// 如 20

// 输出描述
// 输出最多能装多少个快递
// 如 3

// 示例一
// 输入
// 5,10,2,11
// 20

func maxPackages(weightsStr string, capacity int) int {
	// weights := make([]int, 0)
	// weightStrs := strings.Split(weightsStr, ",")
	// for _, weightStr := range weightStrs {
	// 	weight, _ := strconv.Atoi(weightStr)
	// 	weights = append(weights, weight)
	// }
	// 处理数据，将数据转为【】int
	weights := make([]int, 0)
	weightStrs := strings.Split(weightsStr, ",")
	for _, v := range weightStrs {
		num, _ := strconv.Atoi(v)
		weights = append(weights, num)
	}
	// 排序
	sort.Ints(weights)
	count := 0
	sum := 0
	// 遍历，计算能够处理多少个
	// for _, weight := range weights {
	// 	sum += weight
	// 	if sum <= capacity {
	// 		count++
	// 	} else {
	// 		break
	// 	}
	// }
	for _, v := range weights {
		sum += v
		if sum <= capacity {
			count++
		} else {
			break
		}
	}

	return count
}

func maxPackagesTest() {
	var weightsStr string
	fmt.Scanln(&weightsStr)
	var capacity int
	fmt.Scanln(&capacity)
	result := maxPackages(weightsStr, capacity)
	fmt.Println(result)
}

// 航天器
// 给航天器一侧加装长方形和正方形的太阳能板(图中的斜线区域);
// 需要先安装两个支柱(图中的黑色竖条);
// 再在支柱的中间部分固定太阳能板;
// 但航天器不同位置的支柱长度不同;
// 太阳能板的安装面积受限于最短一侧的那支支柱的长度;

// 现提供一组整型数组的支柱高度数据;
// 假设每个支柱间的距离相等为一个单位长度;
// 计算如何选择两根支柱可以使太阳能板的面积最大;

// 输入描述
// 10,9,8,7,6,5,4,3,2,1
// 注释，支柱至少有两根，最多10000根，能支持的高度范围1 ~ 10^9的整数

// 柱子的高度是无序的
// 例子中的递减是巧合

// 输出描述
// 可以支持的最大太阳板面积:(10m高支柱和5m高支柱之间)
// 25

// 示例一
// 输入
// 10,9,8,7,6,5,4,3,2,1
// 输出
// 25

func maxSolarPanelArea(heightsStr string) []int {
	heightStrs := strings.Split(heightsStr, ",")
	heights := make([]int, len(heightStrs))
	for i, heightStr := range heightStrs {
		heights[i], _ = strconv.Atoi(heightStr)
	}
	res := []int{}
	indexleft, indexright := 0, 0
	maxArea := 0
	high := 0
	for i, j := 0, len(heights); i < j; {
		if heights[i] < heights[j] {
			high = heights[i]
			i++
		} else {
			high = heights[j]
			j--
		}
		if maxArea < high*(j-i) {
			maxArea, indexleft, indexright = high*(j-i), i, j

		}

	}
	res = append(res, maxArea, indexleft, indexright)
	return res
}

func maxSolarPanelAreaTest() {
	var heightsStr string
	fmt.Scanln(&heightsStr)
	result := maxSolarPanelArea(heightsStr)
	fmt.Printf("可以支持的最大太阳板面积:(%dm高支柱和%dm高支柱之间)\n")
	fmt.Println(result)
}

// 单词接龙
// 题目描述
// 单词接龙的规则是:
// 可用于接龙的单词,首字母必须要与前一个单词的尾字母相同；
// 当存在多个首字母相同的单词时，取长度最长的单词；
// 如果长度也相等，则取字典序最小的单词；
// 已经参与接龙的单词不能重复使用；
// 现给定一组全部由小写字母组成的单词数组，
// 并指定其中一个单词为起始单词，进行单词接龙，
// 请输出最长的单词串。
// 单词串是单词拼接而成的，中间没有空格。

// 单词个数 1 < N < 20
// 单个单词的长度 1 ~ 30

// 输入描述
// 输入第一行为一个非负整数
// 表示起始单词在数组中的索引k
// 0 <= k < N
// 输入的第二行为非负整数N
// 接下来的N行分别表示单词数组中的单词

// 输出描述
// 输出一个字符串表示最终拼接的单词串

// 示例一
// 输入
// 0
// 6
// word
// dd
// da
// dc
// dword
// d
// 输出
// worddwordda

func wordChain(start int, words []string) string {

	used := make(map[string]bool)
	result := words[start]
	used[words[start]] = true
	temp := []string{}
	// 处理字符
	for k, v := range words {
		if k != start {
			temp = append(temp, v)
		}
	}
	//   排序字符。
	sort.Slice(temp, func(i, j int) bool {
		//    特殊条件
		if len(temp[i]) == len(temp[j]) {
			return temp[i] < temp[j]
		}
		return len(temp[i]) > len(temp[j])
	})

	next := false
	// 循环处理
	// candidates := []string{}
	for i := 0; i < len(words)-1; i++ {
		lastChar := result[len(result)-1]
		next = false
		for _, word := range temp {

			if !used[word] && word[0] == lastChar {
				// tempword=result
				result += word
				next = true
				break
			}
		}
		if !next {
			return result
		}

	}
	return result
}

func wordChainTest() {
	var start, n int
	fmt.Scan(&start)
	fmt.Scan(&n)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}
	output := wordChain(start, words)
	fmt.Println(output)
}

// 第k长子串
// 题目描述
// 给定一个字符串
// 只包含大写字母
// 求在包含同一字母的子串中
// 长度第K长的子串
// 相同字母只取最长的子串

// 输入描述
// 第一行 一个子串 1 < len <= 100
// 只包含大写字母
// 第二行为k的值

// 输出描述
// 输出连续出现次数第k多的字母的次数
// 如果子串中只包含同一字母的子串数小于k
// 则输出-1

func kthLongestSameLetterSubstring(s string, k int) int {
	if len(s) == 0 {
		return -1
	}
	// 用于记录每个字母最长子串的长度
	letterLengths := make(map[byte]int)
	curLen := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curLen++
		} else {
			if curLen > letterLengths[s[i-1]] {
				letterLengths[s[i-1]] = curLen
			}
			curLen = 1
		}
	}

	// 处理最后一组连续相同字母的情况
	if curLen > letterLengths[s[len(s)-1]] {
		letterLengths[s[len(s)-1]] = curLen
	}
	// 处理子串
	lengths := make([]int, 0, len(letterLengths))
	for _, length := range letterLengths {
		lengths = append(lengths, length)
	}
	// 排序
	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i] > lengths[j]
	})
	if len(lengths) < k {
		return -1
	}
	return lengths[k-1]
}

func kthLongestSameLetterSubstringTest() {
	var s string
	var k int
	fmt.Scanln(&s)
	fmt.Scanln(&k)
	result := kthLongestSameLetterSubstring(s, k)
	fmt.Println(result)
}

// 喊7，是一个传统的聚会游戏，
// N个人围成一圈，按顺时针从1 - 7编号，
// 编号为1的人从1开始喊数，
// 下一个人喊得数字是上一个人喊得数字+1，
// 但是当将要喊出数字7的倍数或者含有7的话，
// 不能喊出，而是要喊过。

// 假定N个人都没有失误。
// 当喊道数字k时，
// 可以统计每个人喊“过"的次数。

// 现给定一个长度N的数组，
// 存储打乱的每个人喊”过"的次数，
// 请把它还原成正确顺序。

// 即数组的第i个元素存储编号i的人喊“过“的次数

// 输入描述
// 输入为1行
// 空格分割的喊过的次数，
// 注意k并不提供，
// k不超过200
// 数字个数为N

// 输出描述
// 输出为1行
// 顺序正确的喊过的次数，空格分割

// 示例一
// 输入
// 0 1 0
// ¶输出
// 1 0 0
// 解法1
func check7(num int) bool {
	if num <= 0 {
		return false
	}
	if num%7 == 0 {
		return true
	}
	s := strconv.Itoa(num)
	if strings.Contains(s, "7") {
		return true
	}
	return false
}

func find7arrays(count, people int) []int {
	n := 7
	arrays := []int{}
	for count == 0 {
		if check7(n) {
			arrays = append(arrays, n)
			count--
		}
		n++
	}
	res := make([]int, people)
	for i := 0; i < len(arrays); i++ {
		index := arrays[i]%people - 1
		res[index]++
	}
	return res
}

// 解法2
func restoreOrder(input string) string {
	numsStr := strings.Fields(input)
	nums := make([]int, len(numsStr))
	for i, numStr := range numsStr {
		nums[i], _ = strconv.Atoi(numStr)
	}
	n := len(nums)
	result := make([]int, n)
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			realIndex := (i*n + j) % n
			if nums[realIndex] == 0 {
				result[index] = nums[realIndex]
				index++
				nums[realIndex] = -1
				break
			}
		}
	}
	var output string
	for i, num := range result {
		if i > 0 {
			output += " "
		}
		output += fmt.Sprintf("%d", num)
	}
	return output
}

func restoreOrderTest() {
	var input string
	fmt.Scanln(&input)
	output := restoreOrder(input)
	fmt.Println(output)
}

func restoreOrderByRemainder(nums []int) []int {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	correctOrder := make([]int, n)
	// count := 0
	for i := 0; i < sum; i++ {
		remainder := i % n
		correctOrder[remainder]++
	}
	return correctOrder
}

func restoreOrderByRemainderTest() {
	// 示例输入
	input := []int{0, 1, 0}
	result := restoreOrderByRemainder(input)
	fmt.Println(result)
}

// 题目描述
// 删除字符串中出现次数最少的字符
// 如果多个字符出现次数一样则都删除

// 输入描述
// 输入只包含小写字母

// 输出描述
// 输出删除后剩余的字符
// 若删除后字符串长度为0，则输出empty

// 示例一
// 输入
// abcdd
// 输出
// dd

func deleteLeastOccur(s string) string {
	if s == "" {
		return "empty"
	}
	// 用于统计每个字符出现的次数
	charCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}
	// 找出最小值
	counts := make([]int, 0, len(charCount))
	for _, count := range charCount {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	minCount := counts[0]
	// 遍历，删除最小值对应的字符
	var result string
	for i := 0; i < len(s); i++ {
		if charCount[s[i]] > minCount {
			result += string(s[i])
		}
	}
	if result == "" {
		return "empty"
	}
	return result

}

func deleteLeastOccurTest() {
	var s string
	fmt.Scanln(&s)
	output := deleteLeastOccur(s)
	fmt.Println(output)
}

// 最小步骤数
// 题目描述
// 一个正整数数组，设为nums
// 最大为100个成员
// 求从第一个成员开始正好走到数组最后一个成员所使用的最小步骤数

// 要求：

// 第一步，必须从第一元素起，且1 <= 第一步步长 < len / 2 (len为数组长度)
// 从第二步开始只能以所在成员的数字走相应的步数，不能多不能少，如果目标不可达返回-1，只输出最小的步骤数量
// 只能向数组的尾部走不能向回走
// 输入描述
// 一个正整数数组，元素用空格分割
// 数组长度 < 100

// 输出描述
// 正整数，最小步数
// 不存在输出-1

// 示例一
// 输入
// 7 5 9 4 2 6 8 3 5 4 3 9
// 输出
// 2

func minSteps(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	// 尝试不同的第一步步长
	for firstStep := 1; firstStep < length/2; firstStep++ {
		visited := make(map[int]bool)
		steps := 1
		curIndex := firstStep
		visited[curIndex] = true

		for {
			// 判断是否越界
			if curIndex >= length {
				return -1
			}
			// 到达最后一个元素，返回当前步数
			if curIndex == length-1 {
				return steps
			}
			// 按照当前位置元素的值移动
			nextIndex := curIndex + nums[curIndex]
			steps++
			// 判断是否越界或者陷入循环（已经访问过该位置）
			if nextIndex >= length || visited[nextIndex] {
				break
			}
			visited[nextIndex] = true
			curIndex = nextIndex
		}
	}
	return -1
}

func minStepsTest() {
	var input string
	fmt.Scanln(&input)
	numsStr := strings.Split(input, " ")

	// numsStr := input.split(" ")
	nums := make([]int, len(numsStr))
	for i, numStr := range numsStr {
		nums[i], _ = strconv.Atoi(numStr)
	}
	result := minSteps(nums)
	fmt.Println(result)
}

// 题目描述
// 在通信系统中有一个常见的问题是对用户进行不同策略的调度，会得到不同系统消耗的性能。
// 假设由N个待串行用户，每个用户可以使用A/B/C三种不同的调度策略。
// 不同的策略会消耗不同的系统资源，请你根据如下规则进行用户调度，并返回总的消耗资源数。
// 规则是：
// 相邻的用户不能使用相同的调度策略，例如：
// 第一个用户使用A策略，则第二个用户只能使用B和C策略。
// 对单的用户而言，不同的调度策略对系统资源的消耗可以规划后抽象为数值，例如：
// 某用户分别使用A B C策略的系统消耗，分别为15 8 17，
// 每个用户依次选择当前所能选择的对系统资源消耗最少的策略,局部最优，
// 如果有多个满足要求的策略，选最后一个。

// 输入描述
// 第一行表示用户个数N
// 接下来表示每一行表示一个用户分别使用三个策略的资源消耗
// resA resB resC

// 输出描述
// 最优策略组合下的总的系统消耗资源数

// 示例一
// 输入
// 3
// 15 8 17
// 12 20 9
// 11 7 5
// 输出
// 24

func minResourceConsumption(n int, costs [][]int) int {
	// 用于记录每个用户选择的策略
	chosenStrategies := make([]int, n)
	// 初始化第一个用户选择资源消耗最小的策略
	minCost := costs[0][0]
	chosenStrategies[0] = 0
	for j := 1; j < 3; j++ {
		if costs[0][j] < minCost {
			minCost = costs[0][j]
			chosenStrategies[0] = j
		}
	}
	for i := 1; i < n; i++ {
		availableStrategies := []int{}
		// 根据前一个用户选择的策略确定当前用户可用的策略
		if chosenStrategies[i-1] == 0 {
			availableStrategies = append(availableStrategies, 1, 2)
		} else if chosenStrategies[i-1] == 1 {
			availableStrategies = append(availableStrategies, 0, 2)
		} else {
			availableStrategies = append(availableStrategies, 0, 1)
		}
		minCost = costs[i][availableStrategies[0]]
		chosenStrategies[i] = availableStrategies[0]
		for _, strategy := range availableStrategies[1:] {
			if costs[i][strategy] < minCost {
				minCost = costs[i][strategy]
				chosenStrategies[i] = strategy
			}
		}
	}
	totalCost := 0
	for i := 0; i < n; i++ {
		totalCost += costs[i][chosenStrategies[i]]
	}
	return totalCost
}

func minResourceConsumptionTest() {
	var n int
	fmt.Scan(&n)
	costs := make([][]int, n)
	for i := 0; i < n; i++ {
		costs[i] = make([]int, 3)
		fmt.Scan(&costs[i][0], &costs[i][1], &costs[i][2])
	}
	result := minResourceConsumption(n, costs)
	fmt.Println(result)
}



