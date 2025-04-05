package book

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// findMinColors returns the minimum number of colors needed to color N numbers such that all numbers of the same color can be divided by the smallest number of the same color.
func findMinColors(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sort the slice in ascending order.
	sort.Ints(nums)

	// Map to keep track of the smallest number for each color group.
	colorGroups := make(map[int]bool)

	colorCount := 0
	for _, num := range nums {
		canBeColored := false
		for smallest := range colorGroups {
			if num%smallest == 0 {
				canBeColored = true
				break
			}
		}

		// If the current number cannot be colored with any existing color, create a new color.
		if !canBeColored {
			colorCount++
			colorGroups[num] = true // Use the current number as the smallest for this new color group.
		}
	}

	return colorCount
}

// 给定一个非空数组(列表)，起元素数据类型为整型，
// 请按照数组元素十进制最低位从小到大进行排序，
// 十进制最低位相同的元素，相对位置保持不变，
// 当数组元素为负值时，十进制最低为等同于去除符号位后对应十进制值最低位。

func sortletter(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		lasta := math.Abs(float64(a[i])) - (math.Abs(float64(a[i]))/10)*10
		lastb := math.Abs(float64(a[j])) - (math.Abs(float64(a[j]))/10)*10
		return lasta < lastb
	})
	return a
}

func removalSum(a []int, nums int) int {
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	b := []int{}
	for k, _ := range m {
		b = append(b, k)
	}
	sort.Ints(b)
	if nums*2 > len(b) {
		return -1
	}
	sum := 0
	// for i, j := 0, len(b); i < nums; {
	// 	sum += b[i] + b[j]
	// 	i++
	// 	j--
	// }
	j := len(b) - 1
	for i := 0; i < nums; i++ {
		sum += b[i] + b[j-i]
	}
	return sum
}

func findcontinuousnums(num int) [][]int {
	// index := 1
	var res [][]int
	for i := 1; i <= num; i++ {
		sum := 0
		temp := []int{}
		for j := i; sum < num; j++ {
			sum += j
			temp = append(temp, j)
			if sum == num {
				res = append(res, temp)
				break
			}
		}

	}

	return res
}
func printfindcontinuousnums(nums [][]int, index int) {
	// fmt.Println("")

	res := []string{}
	for _, v := range nums {
		str := fmt.Sprintf("%d%s", index, "=")
		temp := ""
		for _, key := range v {
			temp += fmt.Sprintf("%d%s", key, "+")
		}
		temp = temp[:len(temp)-1]
		str = str + temp
		res = append(res, str)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

// 有一个N个整数的数组，和一个长度为M的窗口。
// 窗口从数组内的第一个数开始滑动，直到窗口不能滑动为止。
// 每次滑动产生一个窗口，和窗口内所有数的和，
// 求窗口滑动产生的所有窗口和的最大值

func findMaxWindowsNums(nums []int, index int) int {
	res := 0
	// i的取值范围需要注意
	for i := 0; i <= len(nums)-index; i++ {
		sum := 0
		// j的最大限制需要注意

		for j := i; j < i+index; j++ {
			sum += nums[j]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

func MemoryPoolUse() {
	var s string
	fmt.Scanln(&s)
	result := make(map[int]int)
	pairs := strings.Split(s, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		if len(kv) != 2 {
			break
		}
		key, _ := strconv.Atoi(kv[0])
		value, _ := strconv.Atoi(kv[1])
		result[key] = value
	}

}

// MemoryBlock 定义每个内存块
type MemoryBlock struct {
	size  int // 内存块大小
	count int // 此大小内存块的数量
}

// MemoryPool 定义内存池
type MemoryPool struct {
	blocks []MemoryBlock // 存储不同大小粒度的内存块
}

// NewMemoryPool 初始化内存池
func NewMemoryPool(blocks []MemoryBlock) *MemoryPool {
	return &MemoryPool{
		blocks: blocks,
	}
}

// Allocate 尝试按顺序分配内存，并返回结果列表
func (mp *MemoryPool) Allocate(requests []int) []bool {
	results := make([]bool, len(requests))

	for i, requestSize := range requests {
		allocated := false
		for j := range mp.blocks {
			if mp.blocks[j].size >= requestSize && mp.blocks[j].count > 0 {
				mp.blocks[j].count--
				results[i] = true
				allocated = true
				break
			}
		}
		if !allocated {
			results[i] = false
		}
	}

	return results
}

func TestMemoryBlock() {
	// 假设我们有如下粒度的内存块：10字节、20字节、40字节，分别有3个、2个、1个。
	blocks := []MemoryBlock{{size: 10, count: 3}, {size: 20, count: 2}, {size: 40, count: 1}}
	pool := NewMemoryPool(blocks)

	// 用户的内存申请序列
	requests := []int{10, 20, 5, 15, 40, 10}

	// 按照规则分配内存并打印结果
	results := pool.Allocate(requests)
	fmt.Println("Allocation results:", results)

	// 输出当前内存池状态
	fmt.Print("Current memory pool state:\n")
	sort.Slice(pool.blocks, func(i, j int) bool {
		return pool.blocks[i].size < pool.blocks[j].size
	})
	for _, block := range pool.blocks {
		fmt.Printf("Block size %d bytes, remaining count: %d\n", block.size, block.count)
	}
}

// 快速排序
func quickSort(arr []int) {
	if len(arr) < 2 {
		return // 数组长度小于2时，已经是有序的了，直接返回
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivotIndex := len(arr) / 2

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Partition the array
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Move the pivot to its final place
	arr[left], arr[right] = arr[right], arr[left]

	// Recursively apply the same logic to the left and right subarrays
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

// 快速排序的辅助函数，用于保持原始接口不变
func QuickSort(arr []int) []int {
	quickSort(arr)
	return arr
}

// 背包问题

// Knapsack 实现了0/1背包问题的解决方案。
func Knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	if n == 0 || capacity <= 0 {
		return 0
	}

	// 创建一个二维dp数组，dp[i][j]表示前i个物品放入容量为j的背包可以获得的最大价值。
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// 填充dp表
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] > w {
				// 如果当前物品重量大于背包剩余容量，则不选这个物品
				dp[i][w] = dp[i-1][w]
			} else {
				// 否则选择或不选择这个物品，取两者中的最大值
				dp[i][w] = max3(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			}
		}
	}

	// 返回最后的结果，即考虑所有物品且背包容量为capacity时的最大价值
	return dp[n][capacity]
}

// 辅助函数，返回两个整数中的较大者
func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestKnapsack() {
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 5
	maxValue := Knapsack(weights, values, capacity)
	fmt.Printf("The maximum value that can be put in a knapsack of capacity %d is %d\n", capacity, maxValue)
}

func dp(wight []int, value []int, index int) int {
	n := len(wight)
	if len(wight) == 0 || index < 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for k, _ := range dp {
		dp[k] = make([]int, index+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= index; j++ {
			if wight[i] > j {
				dp[i+1][j] = dp[i][j]
			} else {

				dp[i+1][j] = max3(dp[i][j], dp[i][j-wight[i]]+value[i])
			}

		}
	}
	return dp[n][index]
}

// 一个工厂有m条流水线
// 来并行完成n个独立的作业
// 该工厂设置了一个调度系统
// 在安排作业时，总是优先执行处理时间最短的作业
// 现给定流水线个数m
// 需要完成的作业数n
// 每个作业的处理时间分别为 t1,t2...tn
// 请你编程计算处理完所有作业的耗时为多少
// 当n > m时 首先处理时间短的m个作业进入流水线
// 其他的等待
// 当某个作业完成时，
// 依次从剩余作业中取处理时间最短的
// 进入处理
func pipeline(m, n int, array []int) int {
	max := 0
	if m > n {

		for _, v := range array {
			if v > max {
				max = v
			}
		}
		return max
	} else {
		temp := make([]int, m)
		// count := n/m + 1
		for i := 0; i < m; i++ {
			temp[i] = array[i]
		}
		for i := m; i < n; i++ {
			sort.Ints(temp)
			temp[0] += array[i]
		}
		for _, v := range temp {
			if v > max {
				max = v
			}
		}
		return max
	}
}

// 并发控制1000个，超时任务

func contextWork(ctx context.Context, wg *sync.WaitGroup, id int) {
	wg.Add(1)
	defer wg.Done()
	select {
	case <-time.After(time.Second):
		fmt.Println(id)
	case <-ctx.Done():
		fmt.Println("超时")
	}
}

func testcontextwork() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var wg sync.WaitGroup
	// wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go contextWork(ctx, &wg, i)
	}

	wg.Wait()
}

// 幼儿园两个班的小朋友排队时混在了一起
// 每个小朋友都知道自己跟前面一个小朋友是不是同班
// 请你帮忙把同班的小朋友找出来
// 小朋友的编号为整数
// 与前面一个小朋友同班用Y表示
// 不同班用N表示

func findSameClass(array []bool) [][]int {
	temp := false
	m := make([]bool, len(array))
	for k, v := range array {
		if v == true {
			temp = !temp
		}
		m[k] = temp
	}
	res := make([][]int, 2)
	for k, v := range m {
		if v == true {
			res[0] = append(res[0], k+1)
		} else {
			res[1] = append(res[1], k+1)
		}
	}
	return res
}


