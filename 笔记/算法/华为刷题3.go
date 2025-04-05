package book

import (
	"fmt"
	"sort"
	"strings"
)

// 题目描述
// 小明来到学校当老师， 需要将学生按考试总分或单科分数进行排名， 你能帮帮他吗?
// 输入描述
// 第1行输入两个整数， 学生人数n和科目数量m。
// ⋅0<n<100
// ⋅0<m<10
// 第2行输入m 个科目名称， 彼此之间用空格隔开。
// · 科目名称只由英文字母构成， 单个长度不超过10个字符。
// · 科目的出现顺序和后续输入的学生成绩——对应。
// · 不会出现重复的科目名称。
// 第3行开始的n行， 每行包含一个学生的姓名和该生m个科目的成绩 (空格隔开)
// · 学生不会重名。
// · 学生姓名只由英文字母构成， 长度不超过10个字符。
// · 成绩是0~100的整数， 依次对应第2行种输入的科目。
// 第n+22行， 输入用作排名的科目名称。若科目不存在， 则按总分进行排序。
// 输出描述
// 输出一行， 按成绩排序后的学生名字， 空格隔开。成绩相同的按照学生姓名字典顺序排序。

type Student struct {
	name   string
	scores []int
	total  int
}

func rankStudents(n int, m int, subjects []string, students []Student, rankingSubject string) []string {
	// 计算每个学生的总分
	for i := range students {
		sum := 0
		for _, score := range students[i].scores {
			sum += score
		}
		students[i].total = sum
	}

	// 根据指定科目或总分进行排序
	sort.Slice(students, func(i, j int) bool {
		if rankingSubject != "" {
			subjectIndex := -1
			for idx, sub := range subjects {
				if sub == rankingSubject {
					subjectIndex = idx
					break
				}
			}
			if subjectIndex == -1 {
				return false
			}
			if students[i].scores[subjectIndex] > students[j].scores[subjectIndex] {
				return true
			} else if students[i].scores[subjectIndex] == students[j].scores[subjectIndex] {
				return students[i].name < students[j].name
			}
			return false
		}
		if students[i].total > students[j].total {
			return true
		} else if students[i].total == students[j].total {
			return students[i].name < students[j].name
		}
		return false
	})

	result := make([]string, n)
	for i := range students {
		result[i] = students[i].name
	}
	return result
}

// 需要处理输入
// func rankStudentsTest() {
// 	var n, m int
// 	fmt.Scan(&n, &m)
// 	fmt.Scanln()
// 	subjectsStr := strings.Split(strings.TrimSpace(), " ")
// 	subjects := make([]string, m)
// 	for i := range subjectsStr {
// 		subjects[i] = subjectsStr[i]
// 	}
// 	students := make([]Student, n)
// 	for i := 0; i < n; i++ {
// 		// studentInfo := strings.Split(strings.TrimSpace(fmt.Scanln()), " ")
// 		students[i] = Student{
// 			name:   studentInfo[0],
// 			scores: make([]int, m),
// 		}
// 		for j := 1; j <= m; j++ {
// 			students[i].scores[j-1], _ = strconv.Atoi(studentInfo[j])
// 		}
// 	}
// 	rankingSubject := strings.TrimSpace(fmt.Scanln())
// 	result := rankStudents(n, m, subjects, students, rankingSubject)
// 	for i := 0; i < len(result); i++ {
// 		if i != 0 {
// 			fmt.Printf(" ")
// 		}
// 		fmt.Printf("%s", result[i])
// 	}
// 	fmt.Println()
// }

// 题目描述
// 某学校举行运动会， 学生们按编号(1、2、3…n)进行标识， 现需要按照身高由低到高排列，对身高相同的人， 按体重由轻到重排列；
// 对于身高体重都相同的人， 维持原有的编号顺序关系。请输出排列后的学生编号。
// 输入描述
// 两个序列， 每个序列由n个正整数组成 ((0<n<=100)。
// 第一个序列中的数值代表身高， 第二个序列中的数值代表体重。
// 输出描述
// 排列结果， 每个数值都是原始序列中的学生编号， 编号从1开始

type StudentTwo struct {
	index  int
	height int
	weight int
}

func rankStudents2Test() {
	var n int
	fmt.Scan(&n)
	heights := make([]int, n)
	weights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	Studenttemp := make([]StudentTwo, n)
	for i := 0; i < n; i++ {
		Studenttemp[i] = StudentTwo{
			index:  i + 1,
			height: heights[i],
			weight: weights[i],
		}
	}

	sort.Slice(Studenttemp, func(i, j int) bool {
		if Studenttemp[i].height < Studenttemp[j].height {
			return true
		} else if Studenttemp[i].height == Studenttemp[j].height {
			return Studenttemp[i].weight < Studenttemp[j].weight
		}
		return false
	})

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = Studenttemp[i].index
	}
	for i := 0; i < len(result); i++ {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", result[i])
	}
	fmt.Println()
}

// 题目描述
// 给定一个字符串s， 最多只能进行一次变换， 返回变换后能得到的最小字符串 (按照字典序进行比较) 。
// 变换规则： 交换字符串中任意两个不同位置的字符。
// 输入描述
// 一串小写字母组成的字符串s
// 输出描述
// 按照要求进行变换得到的最小字符串。
func minString(s string) string {
	minS := s
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			newS := s[:i] + string(s[j]) + s[i+1:j] + string(s[i]) + s[j+1:]
			if newS < minS {
				minS = newS
			}
		}
	}
	return minS
}

func minStringTest() {
	var s string
	fmt.Scan(&s)
	result := minString(s)
	fmt.Println(result)
}

// 题目描述
// 为了充分发挥GPU[算力]，需要尽可能多的将任务交给GPU执行， 现在有一个任务数组， 数组元素表示在这1秒内新增的任务个数且每秒都有新增任务。
// 假设GPU最多一次执行n个任务， 一次执行耗时1秒， 在保证GPU不空闲情况下， 最少需要多长时间执行完成。
// 输入描述
// · 第一个参数为GPU一次最多执行的任务个数， 取值范围[1，10000]
// · 第二个参数为任务数组长度， 取值范围[1，10000]
// · 第三个参数为任务数组， 数字范围[1，10000]
// 输出描述
// 执行完所有任务最少需要多少秒。

func minTime(n int, m int, tasks []int) int {
	totalTime := 0
	remainingTasks := 0
	for _, task := range tasks {
		remainingTasks += task
		for remainingTasks > 0 {
			remainingTasks -= n
			totalTime += 1
		}
	}
	return totalTime
}

func minTimeTest() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	tasks := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&tasks[i])
	}
	result := minTime(n, m, tasks)
	fmt.Println(result)
}

// 题目描述
// 小朋友出操， 按学号从小到大排成一列；小明来迟了， 请你给小明出个主意， 让他尽快找到他应该排的位置。
// 算法复杂度要求不高于nLog(n);；学号为整数类型， 队列规模<=10000;
// 输入描述
// 1、第一行：输入已排成队列的小朋友的学号(正整数)，以”，”隔开
// 例如:93,95,97,100,102,123,155
// 2、第二行:小明学号, 如110;
// 输出描述
// 输出一个数字， 代表队列位置(从1开始)例如：
// 6

func findPosition(queueStr string, x int) int {
	nums := make([]int, 0)
	parts := strings.Split(queueStr, ",")
	for _, part := range parts {
		num := 0
		fmt.Sscanf(part, "%d", &num)
		nums = append(nums, num)

	}
	nums = append(nums, x)
	sort.Ints(nums)
	for i, v := range nums {
		if v == x {
			return i + 1
		}
	}
	return -1
}
func findPositionTest() {
	var queueStr string
	var x int
	fmt.Scanln(&queueStr)
	fmt.Scanln(&x)
	pos := findPosition(queueStr, x)
	fmt.Println(pos)
}

// 题目描述
// 现有N个任务需要处理， 同一时间只能处理一个任务， 处理每个任务所需要的时间固定为1。
// 每个任务都有最晚处理时间限制和积分值， 在最晚处理时间点之前处理完成任务才可获得对应的积分奖励。
// 可用于处理任务的时间有限， 请问在有限的时间内， 可获得的最多积分。
// 输入描述
// 第一行为一个数N， 表示有 N 个任务
// ⋅1≤N≤100
// 第二行为一个数T， 表示可用于处理任务的时间
// ⋅1≤T≤100
// 接下来N行， 每行两个空格分隔的整数 (SLA和V) ，SLA表示任务的最晚处理时间， V表示任务对应的积分。

// 可能有问题

func maxPoints(n int, t int, tasks [][]int) int {
	// dp[i][j]表示前i个任务，使用j时间能获得的最大积分
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= t; j++ {
			// 不选第i个任务
			dp[i][j] = dp[i-1][j]
			sla, v := tasks[i-1][0], tasks[i-1][1]
			if j >= 1 && sla >= j {
				// 选第i个任务，取较大值
				dp[i][j] = max2(dp[i][j], dp[i-1][j-1]+v)
			}
		}
	}
	return dp[n][t]
}

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func maxPointsTest() {
	var n, t int
	fmt.Scanln(&n)
	fmt.Scanln(&t)
	tasks := make([][]int, n)
	for i := range tasks {
		tasks[i] = make([]int, 2)
		fmt.Scanf("%d %d", &tasks[i][0], &tasks[i][1])
	}
	result := maxPoints(n, t, tasks)
	fmt.Println(result)
}

// 题目描述
// 橱窗里有一排宝石， 不同的宝石对应不同的价格， 宝石的价格标记为 gems[i]
// 0≤i<n
// n = gems. length
// 宝石可同时出售0个或多个， 如果同时出售多个， 则要求出售的宝石编号连续；
// 例如客户最大购买宝石个数为m， 购买的宝石编号必须为： gems[i]，\mathsf gems[i+1],ldots,mathsfgems[i+m-1]0≤i<n
// m≤n
// 假设你当前拥有总面值为 value的钱， 请问最多能购买到多少个宝石， 如无法购买宝石， 则返回0。
// 输入描述
// 第一行输入n， 参数类型为 int， 取值范围： [0，10^6]，表示橱窗中宝石的总数量。
// 之后n行分别表示从第0个到第n-1个宝石的价格， 即 gems[0]到\mathsf gems[n-1]的价格, 类型为 int, 取值范围: (0,1000]。
// 之后一行输入v, 类型为 int, 取值范围: [0,10^9],表示你拥有的钱。
// 输出描述
// 输出 int类型的返回值， 表示最大可购买的宝石数量。

func maxGems(n int, gems []int, v int) int {
	maxLen := 0
	for left := 0; left < n; left++ {
		sum := 0
		for right := left; right < n; right++ {
			sum += gems[right]
			if sum <= v {
				maxLen = max2(maxLen, right-left+1)
			} else {
				break
			}
		}
	}
	return maxLen
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxGemsTest() {
	var n int
	fmt.Scanln(&n)
	gems := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&gems[i])
	}
	var v int
	fmt.Scanln(&v)
	result := maxGems(n, gems, v)
	fmt.Println(result)
}

// 题目描述
// RSA加密算法在网络安全世界中无处不在， 它利用了极大整数因数分解的困难度， 数据越大， 安全系数越高， 给定一个32位正整数， 请对其进行因数分解， 找出是哪两个素数的乘积。
// 输入描述
// 一个正整数 num, (0<num<=2147483647
// 输出描述
// 如果成功找到， 以单个空格分割， 从小到大输出两个素数， 分解失败， 请输出·-1,-1

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorize(num int) (int, int) {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 && isPrime(i) && isPrime(num/i) {
			return i, num / i
		}
	}
	return -1, -1
}

func factorizeTest() {
	var num int
	fmt.Scanln(&num)
	p, q := factorize(num)
	fmt.Printf("%d %d\n", p, q)
}

// 题目描述
// 现有两组服务器A和B， 每组有多个算力不同的CPU， 其中A[i]是A组第i个CPU的运算能力， B[i]是B组第i个CPU的运算能力。一组服务器的总算力是各CPU的算力之和。
// 为了让两组服务器的算力相等， 允许从每组各选出一个CPU进行一次交换，
// 求两组服务器中， 用于交换的CPU的算力， 并且要求从A组服务器中选出的CPU， 算力尽可能小。
// 输入描述
// 第一行输入为L1和L2， 以空格分隔， L1表示A组服务器中的CPU数量， L2表示B组服务器中的CPU数量。
// 第二行输入为A组服务器中各个CPU的算力值， 以空格分隔。
// 第三行输入为B组服务器中各个CPU的算力值， 以空格分隔。
// 1≤L1≤10000
// 1≤L2≤10000
// 1 \le \mathsf A[ \mathsf i] \le 1000000
// 1≤B[i]≤100000
// 输出描述
// 对于每组测试数据， 输出两个整数， 以空格分隔， 依次表示A组选出的CPU算力， B组选出的CPU算力。
// 要求从A组选出的CPU的算力尽可能小。

func findExchangeCPUs(l1, l2 int, a, b []int) (int, int) {
	sumA, sumB := 0, 0
	for _, v := range a {
		sumA += v
	}
	for _, v := range b {
		sumB += v
	}
	diff := sumA - sumB
	sort.Ints(a)
	sort.Ints(b)
	tempdiff := 0
	tempa, tempb := 0, 0
	for _, ai := range a {
		for _, bi := range b {
			newDiff := (sumA - ai + bi) - (sumB - bi + ai)
			if newDiff == 0 {
				return ai, bi
			} else {
				temp := newDiff - diff
				if tempdiff < temp {
					tempdiff, tempa, tempb = temp, ai, bi

				}
			}

		}

	}
	return tempa, tempb
}

func findExchangeCPUsTest() {
	var l1, l2 int
	fmt.Scanln(&l1, &l2)
	a := make([]int, l1)
	b := make([]int, l2)
	for i := 0; i < l1; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < l2; i++ {
		fmt.Scanf("%d", &b[i])
	}
	cpuA, cpuB := findExchangeCPUs(l1, l2, a, b)
	fmt.Printf("%d %d\n", cpuA, cpuB)
}

// 题目描述
// 小明在玩一个游戏， 游戏规则如下： 在游戏开始前， 小明站在坐标轴原点处 (坐标值为0) .
// 给定一组指令和一个幸运数， 每个指令都是一个整数， 小明按照指令前进指定步数或者后退指定步数。前进代表朝坐标轴的正方向走， 后退代表朝坐标轴的负方向走。
// 幸运数为一个整数， 如果某个指令正好和幸运数相等， 则小明行进步数+1。
// 例如：
// 幸运数为3, 指令为[2,3,0,-5]
// 指令为2， 表示前进2步；
// 指令为3， 正好和幸运数相等， 前进3+1=4步；
// 指令为0， 表示原地不动， 既不前进， 也不后退。
// 指令为-5， 表示后退5步。
// 请你计算小明在整个游戏过程中， 小明所处的最大坐标值。
// 输入描述
// 第一行输入1个数字， 代表指令的总个数n(1≤n≤100)
// 第二行输入1个数字， 代表幸运数rm(-100≤m≤100)

func maxPosition(n int, m int, instructions []int) int {
	maxPos := 0
	currentPos := 0
	for _, v := range instructions {
		if v == m {
			currentPos += v + 1
		} else {
			currentPos += v
		}
		if currentPos > maxPos {
			maxPos = currentPos
		}
	}
	return maxPos
}

func maxPositionTest() {
	var n, m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)
	instructions := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &instructions[i])
	}
	result := maxPosition(n, m, instructions)
	fmt.Println(result)
}

// 题目描述
// 有位客人来自异国， 在该国使用m进制计数。该客人有个幸运数字n(n<m),， 每次购物时， 其总是喜欢计算本次支付的花费(折算为异国的价格后)中存在多少幸运数字。问： 当其购买一个在我国价值k的产品时， 其中包含多少幸运数字?
// 输入描述
// 第一行输入为 k,n,m。
// 其中：
// • k表示该客人购买的物品价值 (以十进制计算的价格)
// • n表示该客人的幸运数字
// • m表示该客人所在国度的采用的进制
// 输出描述

func countLuckyNumber(k, n, m int) int {
	count := 0
	numInM := ""
	// 将十进制数k转换为m进制数
	for k > 0 {
		numInM = fmt.Sprintf("%d", k%m) + numInM
		k /= m
	}
	// 遍历m进制数的每一位
	for _, char := range numInM {
		digit := int(char - '0')
		if digit == n {
			count++
		}
	}
	return count
}

func countLuckyNumberTest() {
	var k, n, m int
	fmt.Scanln(&k)
	fmt.Scanln(&n)
	fmt.Scanln(&m)
	result := countLuckyNumber(k, n, m)
	fmt.Println(result)
}

// 题目描述
// 园区某部门举办了 Family Day， 邀请员工及其家属参加；
// 将公司园区视为一个矩形， 起始园区设置在左上角， 终点园区设置在右下角；
// 家属参观园区时， 只能向右和向下园区前进， 求从起始园区到终点园区会有多少条不同的参观路径。
// 起点  0  0  0
// 0  1  0
// 0  0  0  终点
// 输入描述
// 第一行为园区的长和宽；
// 后面每一行表示该园区是否可以参观， 0表示可以参观， 1表示不能参观
// 输出描述
// 输出为不同的路径数量

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstaclesTest() {
	var m, n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	obstacleGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		obstacleGrid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &obstacleGrid[i][j])
		}
	}
	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(result)
}

// 题目描述
// 主管期望你来实现英文输入法单词联想功能。
// 需求如下：
// · 依据用户输入的单词前缀， 从已输入的英文语句中联想出用户想输入的单词， 按字典序输出联想到的单词序列，
// · 如果联想不到， 请输出用户输入的单词前缀。
// 注意：
// 1.英文单词联想时， 区分大小写
// 2.缩略形式如"don't", 判定为两个单词, "don"和"t"
// 3.输出的单词序列， 不能有重复单词， 且只能是英文单词， 不能有标点符号
// 输入描述
// 输入为两行。
// 首行输入一段由英文单词 word和标点符号组成的语句 str；
// 接下来一行为一个英文单词前缀 pre。
// ●0<str.length<=10000
// ⋅0<pre<=20
// 输出描述
// 输出符合要求的单词序列或单词前缀， 存在多个时， 单词之间以单个空格分割

func wordSuggestion(sentence, prefix string) []string {
	words := strings.Fields(sentence)
	result := []string{}
	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			result = append(result, word)
		}
	}
	sort.Strings(result)
	if len(result) == 0 {
		return []string{prefix}
	}
	uniqueWords := make([]string, 0, len(result))
	seen := make(map[string]bool)
	for _, w := range result {
		if !seen[w] {
			uniqueWords = append(uniqueWords, w)
			seen[w] = true
		}
	}
	return uniqueWords
}

func wordSuggestionTest() {
	var sentence, prefix string
	fmt.Scanln(&sentence)
	fmt.Scanln(&prefix)
	suggestions := wordSuggestion(sentence, prefix)
	for i, s := range suggestions {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(s)
	}
	fmt.Println()
}

// 题目描述
// 输入一个由N个大小写字母组成的字符串
// 按照ASCII码值从小到大进行排序
// 查找字符串中第K个最小ASCII码值的字母((k>=1)
// 输出该字母所在字符串中的位置索引(字符串的第一个位置索引为0)
// k如果大于字符串长度则输出最大ASCII码值的字母所在字符串的位置索引
// 如果有重复字母则输出字母的最小位置索引
// 输入描述
// 第一行输入一个由大小写字母组成的字符串
// 第二行输入k ， k必须大于0 ， k可以大于输入字符串的长度
// 输出描述
// 输出字符串中第k个最小ASCII码值的字母所在字符串的位置索引
// k如果大于字符串长度则输出最大ASCII码值的字母所在字符串的位置索引
// 如果第k个最小ASCII码值的字母存在重复则输出该字母的最小位置索引

func findIndex(s string, k int) int {
	sortedS := strings.Split(s, "")
	sort.Strings(sortedS)
	if k > len(sortedS) {
		maxAsciiChar := ""
		maxAsciiIndex := 0
		for i, char := range strings.Split(s, "") {
			if char > maxAsciiChar {
				maxAsciiChar = char
				maxAsciiIndex = i
			}
		}
		return maxAsciiIndex
	}
	targetChar := sortedS[k-1]
	for i, char := range strings.Split(s, "") {
		if char == targetChar {
			return i
		}
	}
	return -1
}

func findIndexTest() {
	var s string
	var k int
	fmt.Scanln(&s)
	fmt.Scanln(&k)
	index := findIndex(s, k)
	fmt.Println(index)
}


