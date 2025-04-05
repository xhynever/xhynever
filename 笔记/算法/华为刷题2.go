package book

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// 山脉的个数

func countMountains(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	count := 0
	i := 0
	for i < len(arr) {
		// 寻找山脉的上升沿（从左到右递增）
		for i < len(arr)-1 && arr[i] >= arr[i+1] {
			i++
		}
		if i == len(arr)-1 {
			break
		}
		peakIndex := i + 1
		// 寻找山脉的山顶（继续上升直到达到峰值）
		for peakIndex < len(arr)-1 && arr[peakIndex] < arr[peakIndex+1] {
			peakIndex++
		}
		// 寻找山脉的下降沿（从峰值开始下降）
		for peakIndex < len(arr)-1 && arr[peakIndex] >= arr[peakIndex+1] {
			peakIndex++
		}
		if peakIndex < len(arr) && arr[peakIndex] < arr[peakIndex-1] {
			count++
		}
		i = peakIndex
	}
	return count
}

// 题目描述： 构成指定长度字符串的个数 (本题分值100)
// 给定\textsf M(0<M \le 30)个字符(a-z),从中取出任意字符 (每个字符只能用一次) 拼接成长度为IN(0<N≤5)的字符串，要求相同的字符不能相邻， 计算出给定的字符列表能拼接出多少种满足条件的字符串，
// 输入非法或者无法拼接出满足条件的字符串则返回0。
// 输入描述
// 给定的字符列表和结果字符串长度， 中间使用空格("")拼接
// 输出描述
// 满足条件的字符串个数
// 用例1
// 输入
// 1   aab 2
// 输出

// 硬背递归
func countStrings(s string, n int) int {
	if n == 0 {
		return 1
	}
	if s == "" || n > len(s) {
		return 0
	}
	var helper func(curStr string, leftChars string, leftLength int) int
	helper = func(curStr string, leftChars string, leftLength int) int {
		if leftLength == 0 {
			return 1
		}
		count := 0
		for i := 0; i < len(leftChars); i++ {
			if curStr == "" || leftChars[i] != curStr[len(curStr)-1] {
				newLeftChars := leftChars[:i] + leftChars[i+1:]
				count += helper(curStr+string(leftChars[i]), newLeftChars, leftLength-1)
			}
		}
		return count
	}
	return helper("", s, n)
}

// 题目描述： 用连续自然数之和来表达整数 (本题分值100)
// 一个整数可以由连续的自然数之和来表示。
// 给定一个整数， 计算该整数有几种连续自然数之和的表达式， 且打印出每种表达式
// 输入描述
// 一个目标整数⁻T(1<=T<=1000)
// 输出描述
// 该整数的所有表达式和表达式的个数。如果有多种表达式， 输出要求为：
// 自然数个数最少的表达式优先输出
// 每个表达式中按自然数递增的顺序输出， 具体的格式参见样例。
// 在每个测试数据结束时， 输出一行"Result：X"，其中X是最终的表达式个数。

func findContinuousSumExpressions(num int) {
	count := 0
	for n := 1; n <= num/2+1; n++ {
		for start := 1; start < num; start++ {
			sum := 0
			for i := start; i < start+n; i++ {
				sum += i
			}
			if sum == num {
				fmt.Printf("%d=", num)
				for j := start; j < start+n; j++ {
					if j != start {
						fmt.Printf("+%d", j)
					} else {
						fmt.Printf("%d", j)
					}
				}
				fmt.Println()
				count++
			}
			if sum > num {
				break
			}
		}
	}
	fmt.Printf("Result：%d\n", count)
}

func findContinuousSumExpressionsTest() {
	var num int
	fmt.Scan(&num)
	findContinuousSumExpressions(num)
}

// 题目描述： 密码输入检测 (本题分值100)
// 给定用户密码输入流 input， 输入流中字符'<"表示退格，可以清除前一个输入的字符， 请你编写程序， 输出最终得到的密码字符， 并判断密码是否满足如下的密码安全要求。
// 密码安全要求如下：
// 1.密码长度>=8;
// 2.密码至少需要包含1个大写字母；
// 3.密码至少需要包含1个小写字母；
// 4.密码至少需要包含1个数字；
// 5.密码至少需要包含1个字母和数字以外的非空白特殊字符
// 注意空串退格后仍然为空串， 且用户输入的字符串不包含‘<’字符和空白字符。

func passwordCheck(input string) (string, bool) {
	stack := []rune{}
	for _, char := range input {
		if char == '<' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, char)
		}
	}
	password := string(stack)
	if len(password) < 8 {
		return password, false
	}
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else {
			hasSpecial = true
		}
	}
	return password, hasUpper && hasLower && hasDigit && hasSpecial
}

func passwordCheckTest() {
	var input string
	fmt.Scanln(&input)
	password, valid := passwordCheck(input)
	fmt.Println(password)
	if valid {
		fmt.Println("密码符合安全要求")
	} else {
		fmt.Println("密码不符合安全要求")
	}
}

// 题目描述： 查找众数及中位数 (本题分值100)
// 众数是指一组数据中出现次数量多的那个数， 众数可以是多个。
// 中位数是指把一组数据从小到大排列， 最中间的那个数， 如果这组数据的个数是奇数， 那最中间那个就是中位数， 如果这组数据的个数为偶数， 那就把中间的两个数之和除以2， 所得的结果就是中位数。
// 查找整型数组中元素的众数并组成一个新的数组， 求新数组的中位数。
// 输入描述
// 输入一个一维整型数组， 数组大小取值范围(0<N<1000,数组中每个元素取值范围0<E<1000
// 输出描述
// 输出众数组成的新数组的中位数

func findModeAndMedian(arr []int) float64 {
	numCount := make(map[int]int)
	// 统计每个数出现的次数
	for _, num := range arr {
		numCount[num]++
	}
	maxCount := 0
	modes := []int{}
	// 找出众数
	for num, count := range numCount {
		if count > maxCount {
			maxCount = count
			modes = []int{num}
		} else if count == maxCount {
			modes = append(modes, num)
		}
	}
	// 对众数组成的新数组进行排序
	sort.Ints(modes)
	length := len(modes)
	if length%2 == 1 {
		return float64(modes[length/2])
	}
	return float64(modes[length/2-1]+modes[length/2]) / 2
}

func findModeAndMedianTest() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	result := findModeAndMedian(arr)
	fmt.Println(result)
}

// 题目描述
// 开头和结尾都是元音字母 (aeiouAEIOU) 的字符串为元音字符串， 其中混杂的非元音字母数量为其瑕疵度。比如：
// 1. “a”、“aa”是元音字符串, 其瑕疵度都为0
// 2. “aiur”不是元音字符串 (结尾不是元音字符)
// 3. “abira”是元音字符串, 其瑕疵度为2
// 给定一个字符串， 请找出指定瑕疵度的最长元音字符子串， 并输出其长度， 如果找不到满足条件的元音字符子串， 输出0。
// 子串： 字符串中任意个连续的字符组成的子序列称为该字符串的子串。
// 输入描述
// 首行输入是一个整数， 表示预期的瑕疵度 flaw， 取值范围[0，65535]。
// 接下来一行是一个仅由字符a-z和A-Z组成的字符串， 字符串长度(0，65535]。
// 输出描述
// 输出为一个整数， 代表满足条件的元音字符子串的长度。

func isVowel(char byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, vowel := range vowels {
		if char == vowel {
			return true
		}
	}
	return false
}

func maxLengthVowelSubstring(flaw int, s string) int {
	maxLength := 0
	for left := 0; left < len(s); left++ {
		for right := left; right < len(s); right++ {
			curFlaw := 0
			curLength := right - left + 1
			isVowelSubstring := true
			for i := left; i <= right; i++ {
				if !isVowel(s[i]) {
					curFlaw++
					if curFlaw > flaw {
						isVowelSubstring = false
						break
					}
				}
			}
			if isVowelSubstring && curLength > maxLength {
				maxLength = curLength
			}
		}
	}
	return maxLength
}

func main() {
	var flaw int
	fmt.Scan(&flaw)
	var s string
	fmt.Scan(&s)
	result := maxLengthVowelSubstring(flaw, s)
	fmt.Println(result)
}

// 题目描述
// 公司组织了一次考试，现在考试结果出来了， 想看一下有没人存在作弊行为，但是员工太多了，需要先对员工进行一次过滤，再进一步确定是否存在作弊行为。
// 过滤的规则为：找到分差最小的员工ID对(p1，p2)列表，要求p1<p2
// 员工个数取值范国：(0<n<100000
// 员工ID为整数，取值范围：0<=n<=10000
// 考试成绩为整数，取值范围：0<=score<=300
// 输入描述
// 员工的ID及考试分数
// 输出描述
// 分差最小的员工ID对(p1，p2)列表，要求|p1<p2。。每一行代表一个集合，每个集合内的员工ID按顺序排列，多行结果也以员工对中p1值大小升序排列(如果p1相同则p2升序)。

type Employee struct {
	ID    int
	Score int
}

func filterEmployees(employees []Employee) [][]int {
	if len(employees) < 2 {
		return nil
	}
	minDiff := 301 // 初始化为最大可能分差 + 1
	pairs := [][]int{}
	for i := 0; i < len(employees)-1; i++ {
		for j := i + 1; j < len(employees); j++ {
			diff := abs(employees[i].Score - employees[j].Score)
			if diff < minDiff {
				minDiff = diff
				pairs = [][]int{}
			}
			if diff == minDiff {
				pair := []int{employees[i].ID, employees[j].ID}
				if employees[i].ID < employees[j].ID {
					pairs = append(pairs, pair)
				}
			}
		}
	}
	// 按照要求对结果进行排序
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] < pairs[j][0] {
			return true
		} else if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return false
	})
	return pairs

}

// func abs(num int) int {
// 	if num < 0 {
// 		return -num
// 	}
// 	return num
// }

func filterEmployeesTest() {
	var n int
	fmt.Scan(&n)
	employees := make([]Employee, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&employees[i].ID, &employees[i].Score)
	}
	sort.Slice(employees, func(i, j int) bool {
		// if
		return employees[i].Score < employees[j].Score
	})

	result := filterEmployees(employees)
	for _, pair := range result {
		fmt.Println(pair[0], pair[1])
	}
}

// 题目描述
// 给你一个字符串s， 字符串s首尾相连成一个环形， 请你在环中找出'o'字符出现了偶数次最长子字符串的长度。给你一个字符串s， 字符串s首尾相连成一个环形， 请你在环中找出'o'字符出现了偶数次最长子字符串的长度。
// 输入描述
// 输入是一串小写字母组成的字符串
// 备注
// 1<=s.length<=5×10ʰ5
// s只包含小写英文字母
func Longstring(s string) int {
	count := 0
	n := len(s)
	for _, v := range s {
		if rune('o') == v {
			count++
		}
	}
	if count%2 == 1 {
		return n - 1
	} else {
		return n
	}
}
func LongstringTest() {
	var s string
	fmt.Scan(&s)
	fmt.Println(Longstring(s))
}

// 题目描述
// 在一个大型体育场内举办了一场大型活动， 由于疫情防控的需要， 要求每位观众的必须间隔至少一个空位才允许落座。
// 现在给出一排观众座位分布图， 座位中存在已落座的观众， 请计算出， 在不移动现有观众座位的情况下， 最多还能坐下多少名观众。
// 输入描述
// 一个数组， 用来标识某一排座位中， 每个座位是否已经坐人。0表示该座位没有坐人， 1表示该座位已经坐人。
// · 1≤数组长度≤10000
// 输出描述
// 整数， 在不移动现有观众座位的情况下， 最多还能坐下多少名观众。

func MaxSeat(s string) int {
	count := 0

	if len(s) == 1 || len(s) == 0 {
		return len(s)
	}

	for i := 0; i < len(s); {
		if s[i] == '1' {
			i += 2
			break
		}
		if (i == 0 && s[1] == '0') || (i < len(s) && s[i-1] == '0' && s[i+1] == '0') || i == len(s)-1 {
			count++
			i += 2

		}

	}

	// if || (i == len(s)-1 && s[i] == 0) {
	// 	count++
	// }
	return count
}

func MaxSeatTest() {
	var s string
	fmt.Scan(&s)
	fmt.Println(MaxSeat(s))

}

// 题目描述
// 寿司店周年庆， 正在举办优惠活动回馈新老客户。
// 寿司转盘上总共有n盘寿司， prices[i]是第i盘寿司的价格，
// 如果客户选择了第i盘寿司， 寿司店免费赠送客户距离第i盘寿司最近的下一盘寿司j， 前提是\mathsf prices[j]< \mathsf prices[i]，如果没有满足条件的j，则不赠送寿司。
// 每个价格的寿司都可无限供应。
// 输入描述
// 输入的每一个数字代表每盘寿司的价格， 每盘寿司的价格之间使用空格分隔， 例如：
// 3 15 6 14
// 表示：
// · 第0盘寿司价格 prices[0]为3

// 暴力循环
func EventPrice(s []string) {
	type goods struct {
		id    int
		price int
	}
	var Goods []goods
	for i := 0; i < len(s); i++ {
		var good goods
		good.id = i
		good.price, _ = strconv.Atoi(s[i])
		Goods = append(Goods, good)
	}
	sort.Slice(Goods, func(i, j int) bool {
		return Goods[i].price < Goods[j].price
	})

}
func EventPriceTest() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// s := scanner.Text()

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s1 := strings.Fields(s)

	EventPrice(s1)
}

// 题目描述
// 孙悟空爱吃蟠桃， 有一天趁着蟠桃园守卫不在来偷吃。已知蟠桃园有N棵桃树， 每颗树上都有桃子， 守卫将在H小时后回来。
// 孙悟空可以决定他吃蟠桃的速度K (个/小时) ，每个小时选一颗桃树， 并从树上吃掉K个， 如果树上的桃子少于 K个， 则全部吃掉， 并且这一小时剩余的时间里不再吃桃。
// 孙悟空喜欢慢慢吃， 但又想在守卫回来前吃完桃子。
// 请返回孙悟空可以在H小时内吃掉所有桃子的最小速度K (K为整数) 。如果以任何速度都吃不完所有桃子， 则返回0。
// 输入描述
// 第一行输入为 N个数字， N表示桃树的数量， 这N个数字表示每颗桃树上蟠桃的数量。
// 第二行输入为一个数字， 表示守卫离开的时间H。
// 其中数字通过空格分割， N、H为正整数， 每颗树上都有蟠桃， 且(0<N<10000,0<H<10000。

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, max(piles)
	for left < right {
		mid := left + (right-left)/2
		totalHours := 0
		for _, pile := range piles {
			totalHours += (pile + mid - 1) / mid
		}
		if totalHours > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func max(nums []int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func minEatingSpeedTest() {
	var input string
	fmt.Scanln(&input)
	pilesStr := strings.Split(input, " ")
	piles := make([]int, len(pilesStr))
	for i, pileStr := range pilesStr {
		piles[i], _ = strconv.Atoi(pileStr)
	}
	var h int
	fmt.Scan(&h)
	result := minEatingSpeed(piles, h)
	fmt.Println(result)
}

// 英雄联盟是一款十分火热的对战类游戏。每一场对战有10位玩家参与， 分为两组， 每组5人。每位玩家都有一个战斗力， 代表着这位玩家的厉害程度。为了对战尽可能精彩， 我们需要把玩家们分为实力尽量相等的两组。一组的实力可以表示为这一组5位玩家的战斗力和。现在， 给你10位玩家的战斗力， 请你把他们分为实力尽量相等的两组。请你输出这两组的实力差。
// 2023年题:
// 部门准备举办一场王者荣耀表演赛， 有10名游戏爱好者参与， 分5为两队， 每队5人。每位参与者都有一个评分， 代表着他的游戏水平。为了表演赛尽可能精彩， 我们需要把10名参赛者分为实力尽量相近的两队。一队的实力可以表示为这一队5名队员的评分总和。
// 现在给你10名参与者的游戏水平评分， 请你根据上述要求分队最后输出这两组的实力差绝对值。
// 例:10名参赛者的评分分别为51834671092, 分组为(135810)(24679), 两组实力差最小, 差值为1。有多种分法, 但实力差的绝对值最小为1。
// 输入描述
// 10个整数， 表示10名参与者的游戏水平评分。范围在[1，10000]之间
// 输出描述
// 1个整数， 表示分组后两组实力差绝对值的最小值.
// 用例1
// 输入：
// 1  1 2 3 4 5 6 7 8 9 10

// 暴力循环
func minDiff(nums []int) {

	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 分别选
	for i := 0; i < len(nums); i++ {

	}

}

// 均衡串定义： 字符串只包含两种字符， 且两种字符的个数相同。
// 给定一个均衡字符串， 请给出可分割成新的均衡子串的最大个数。
// 约定字符串中只包含大写的X和Y两种字符。
// 输入描述
// 均衡串: XXYYXY
// 字符串的长度[2，100001]。给定的字符串均为均衡串
// 输出描述
// 可分割为两个子串：
// XXYY
// XY

func maxSubstrings(s string) int {
	count := 0
	diff := 0
	for _, char := range s {
		if string(char) == "X" {
			diff++
		} else {
			diff--
		}
		if diff == 0 {
			count++
		}
	}
	return count
}

func maxSubstringsTest() {
	s := "XXYYXY"
	result := maxSubstrings(s)
	fmt.Println(result)
}

func work(s []string) {
	m := make(map[string]int)
	// count:=0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

}
