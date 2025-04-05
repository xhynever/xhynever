package book

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// 题目描述
// 给定一个 url前缀和 url后缀，通过，分割需要将其连接为一个完整的 url
// · 如果前缀结尾和后缀开头都没有/，需要自动补上/连接符
// · 如果前缀结尾和后缀开头都为/，需要自动去重
// 约束： 不用考虑前后缀URL不合法情况
// 输入描述
// url前缀(一个长度小于100的字符串) url后缀(一个长度小于100的字符串)
// 输出描述
// 拼接后的 url

func combineURL(prefix, suffix string) string {
	if len(prefix) > 0 && prefix[len(prefix)-1] != '/' && len(suffix) > 0 && suffix[0] != '/' {
		return prefix + "/" + suffix
	}
	if len(prefix) > 0 && prefix[len(prefix)-1] == '/' && len(suffix) > 0 && suffix[0] == '/' {
		return prefix[:len(prefix)-1] + suffix
	}
	return prefix + suffix
}

func combineURLTest() {
	var prefix, suffix string
	fmt.Scanln(&prefix)
	fmt.Scanln(&suffix)
	result := combineURL(prefix, suffix)
	fmt.Println(result)
}

// 题目描述
// 特定大小的停车场， 数组 cars[]表示， 其中1表示有车， 0表示没车。车辆大小不一， 小车占一个车位 (长度1) ，货车占两个车位 (长度2) ，卡车占三个车位 (长度3) 。统计停车场最少可以停多少辆车， 返回具体的数目。
// 输入描述
// 整型字符串数组 cars[]， 其中1表示有车， 0表示没车， 数组长度小于1000。
// 输出描述
// 整型数字字符串， 表示最少停车数目。

func minCars(cars []string) int {
	carArr := make([]int, len(cars))
	for i, v := range cars {
		num, _ := strconv.Atoi(v)
		carArr[i] = num
	}
	count := 0
	i := 0
	for i < len(carArr) {
		if carArr[i] == 1 {
			// count++
			i++
			continue
		}
		if i+2 < len(carArr) && carArr[i+1] == 0 && carArr[i+2] == 0 {
			count++
			i += 3
			continue
		}
		if i+1 < len(carArr) && carArr[i+1] == 0 {
			count++
			i += 2
			continue
		}
		i++
	}
	return count
}

func minCarsTest() {
	var input string
	fmt.Scanln(&input)
	cars := strings.Split(input, ",")
	result := minCars(cars)
	fmt.Println(result)
}

// 题目描述
// 某个产品的RESTfulAPI集合部署在服务器集群的多个节点上， 近期对客户端访问日志进行了采集， 需要统计各个API的访问频次， 根据热点信息在服务器节点之间做负载均衡， 现在需要实现热点信息统计查询功能。
// RESTfulAPI是由多个层级构成， 层级之间使用/连接， 如/A/B/C/D这个地址， A属于第一级， B属于第二级， C属于第三级， D属于第四级。现在负载均衡模块需要知道给定层级上某个名字出现的频次， 未出现过用0表示， 实现这个功能。
// 输入描述
// 第一行为N， 表示访问历史日志的条数， (0<N≤100。
// 接下来N行， 每一行为一个RESTfulAPI的URL地址， 约束地址中仅包含英文字母和连接符/，最大层级为10， 每层级字符串最大长度为10。最后一行为层级L和要查询的关键字
// 输出描述
// 输出给定层级上， 关键字出现的频次， 使用完全匹配方式 (大小写敏感) 。

func countAPILevelFreq(n int, logs []string, level int, keyword string) int {
	count := 0
	for _, log := range logs {
		parts := strings.Split(log, "/")
		if level <= len(parts) {
			if parts[level-1] == keyword {
				count++
			}
		}
	}
	return count
}

func countAPILevelFreqTest() {
	var n int
	fmt.Scanln(&n)
	logs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&logs[i])
	}
	var level int
	var keyword string
	fmt.Scanln(&level)
	fmt.Scanln(&keyword)
	freq := countAPILevelFreq(n, logs, level, keyword)
	fmt.Println(freq)
}

// 题目描述
// 用数组代表每个人的能力 一个比赛活动要求参赛团队的最低能力值为N每个团队可以由一人或者两人组成且一个人只能参加一个团队计算出最多可以派出多少只符合要求的队伍。
// 输入描述
// 第一行代表总人数， 范围1-500000
// 第二行数组代表每个人的能力
// -数组大小, 范围1-500000
// -元素取值, 范围1-500000
// 第三行数值为团队要求的最低能力值， 范围1-500000
// 输出描述
// 最多可以派出的团队数量

func maxTeams(numPeople int, abilities []int, minAbility int) int {
	sort.Ints(abilities)
	left, right := 0, numPeople-1
	teamCount := 0
	for left <= right {
		if abilities[left]+abilities[right] >= minAbility {
			teamCount++
			left++
			right--
		} else {
			left++
		}
	}

	return teamCount
}

func maxTeams2(numPeople int, abilities []int, minAbility int) int {
	sort.Ints(abilities)
	left, right := 0, numPeople-1
	teamCount := 0
	for left <= right {
		if abilities[right] >= minAbility {
			teamCount++
			right--
		} else if abilities[right]+abilities[left] >= minAbility {
			teamCount++
			right--
			left++
		} else {
			left++
		}
	}

	return teamCount
}

func maxTeamsTest() {
	var numPeople int
	fmt.Scanln(&numPeople)
	abilities := make([]int, numPeople)
	for i := 0; i < numPeople; i++ {
		fmt.Scanf("%d", &abilities[i])
	}
	var minAbility int
	fmt.Scanln(&minAbility)
	result := maxTeams(numPeople, abilities, minAbility)
	fmt.Println(result)
}

// 题目描述
// 给定一个含有N个正整数的数组，求出有多少个连续区间 (包括单个正整数) ，它们的和大于等于x。
// 输入描述
// 第一行两个整数Nx(0<N<=100000,0<=x<=1000000)
// 第二行有N个正整数 (每个正整数小于等于100)。
// 输出描述
// 输出一个整数， 表示所求的个数。

func countIntervals(n, x int, nums []int) int {
	count := 0
	for left := 0; left < n; left++ {
		sum := 0
		for right := left; right < n; right++ {
			sum += nums[right]
			if sum >= x {
				count++
			}
		}
	}
	return count
}

func countIntervalsTest() {
	var n, x int
	fmt.Scanln(&n)
	fmt.Scanln(&n, &x)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	result := countIntervals(n, x, nums)
	fmt.Println(result)
}

// 题目描述
// 已知火星人使用的运算符为#、$，其与地球人的等价公式如下：
// x \# y=2^{*}x+3^{*}y+4
// xSy=3^{*}x+y+2
// 1.其中x、y是无符号整数
// 2.地球人公式按C语言规则计算
// 3.火星人公式中， $的优先级高于#，相同的运算符， 按从左到右的顺序计算现有一段火星人的字符串报文， 请你来翻译并计算结果。

func calculateMartianExpression(expr string) int {
	// 先处理 $ 运算符，因为它优先级高
	for {
		match := regexp.MustCompile(`\d+\$\d+`).FindString(expr)
		if match == "" {
			break
		}
		parts := strings.Split(match, "$")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		result := 3*num1 + num2 + 2
		expr = strings.Replace(expr, match, fmt.Sprintf("%d", result), 1)
	}
	// 再处理 # 运算符
	for {
		match := regexp.MustCompile(`\d+\#\d+`).FindString(expr)
		if match == "" {
			break
		}
		parts := strings.Split(match, "#")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		result := 2*num1 + 3*num2 + 4
		expr = strings.Replace(expr, match, fmt.Sprintf("%d", result), 1)
	}
	res, _ := strconv.Atoi(expr)
	return res
}

func calculateMartianExpressionTest() {
	var martianExpr string
	fmt.Scanln(&martianExpr)
	result := calculateMartianExpression(martianExpr)
	fmt.Println(result)
}

// 题目描述
// 在一款虚拟游戏中生活， 你必须进行投资以增强在虚拟游戏中的资产以免被淘汰出局。
// 现有一家 Bank，它提供有若干理财产品m个， 风险及投资回报不同， 你有N (元)进行投资， 能接收的总风险值为X。你要在可接受范围内选择最优的投资方式获得最大回报。
// 备注：
// · 在虚拟游戏中， 每项投资风险值相加为总风险值；
// · 在虚拟游戏中， 最多只能投资2个理财产品；
// · 在虚拟游戏中， 最小单位为整数， 不能拆分为小数；
// · 投资额*回报率=投资回报

// findMaxReturn计算在给定条件下的最大投资回报
func findMaxReturn(m, N, X int, products [][]int) int {
	maxReturn := 0
	// 最多投资2个理财产品，分别考虑投资1个和2个产品的情况
	// 投资1个产品的情况
	for _, product := range products {
		invest, rate := product[0], product[1]
		if invest <= N && invest*rate <= X {
			returnValue := invest * rate
			if returnValue > maxReturn {
				maxReturn = returnValue
			}
		}
	}
	// 投资2个产品的情况
	for i := 0; i < len(products); i++ {
		for j := i + 1; j < len(products); j++ {
			invest1, rate1 := products[i][0], products[i][1]
			invest2, rate2 := products[j][0], products[j][1]
			totalInvest := invest1 + invest2
			totalRisk := invest1*rate1 + invest2*rate2
			if totalInvest <= N && totalRisk <= X {
				returnValue := invest1*rate1 + invest2*rate2
				if returnValue > maxReturn {
					maxReturn = returnValue
				}
			}
		}
	}
	return maxReturn
}

func findMaxReturnTest() {
	var m, N, X int
	fmt.Scanln(&m)
	fmt.Scanln(&N)
	fmt.Scanln(&X)
	products := make([][]int, m)
	for i := 0; i < m; i++ {
		products[i] = make([]int, 2)
		fmt.Scanf("%d %d", &products[i][0], &products[i][1])
	}
	result := findMaxReturn(m, N, X, products)
	fmt.Println(result)
}

// 题目描述
// XX市机场停放了多架飞机， 每架飞机都有自己的航班号CA3385， CZ6678， SC6508等， 航班号的前2个大写字母(或数字) 代表航空公司的缩写， 后面4个数字代表航班信息。但是XX市机场只有一条起飞用跑道， 调度人员需要安排目前停留在机场的航班有序起飞。为保障航班的有序起飞， 调度员首先按照航空公司的缩写 (航班号前2个字母) 对所有航班进行排序， 同一航空公司的航班再按照航班号的后4个数字进行排序最终获得安排好的航班的起飞顺序。请编写一段代码根据输入的航班号信息帮助调度员输出航班的起飞顺序。航空公司缩写排序按照从特殊符号$ & *,0g, AZ排序;
// 输入描述
// 第一行输入航班信息， 多个航班号之间用逗号 (“， ”) 分隔， 输入的航班号不超过100个例如：
// CA3385,CZ6678,SC6508,DU7523,HK4456,MK0987
// 备注： 航班号为6位长度， 后4位为纯数字， 不考虑存在后4位重复的场景
// 输出描述
// CA3385,CZ6678,DU7523,HK4456,MK0987,SC6508

type Flight struct {
	Airline string
	Number  int
}

func toIntTest() {
	var input string
	fmt.Scanln(&input)
	flights := strings.Split(input, ",")
	flightList := make([]string, len(flights))
	for i, flight := range flights {
		flightList[i] = flight
	}
	sort.Slice(flightList, func(i, j int) bool {
		// if flightList[i].Airline < flightList[j].Airline {
		// 	return true
		// } else if flightList[i].Airline == flightList[j].Airline {
		// 	return flightList[i].Number < flightList[j].Number
		// }
		// return false
		tempi, tempj := flightList[i][:2], flightList[j][:2]
		if tempi == tempj {
			return flightList[i][2:] > flightList[j][2:]
		} else {
			return flightList[i][:2] > flightList[j][:2]
		}

	})

	fmt.Println(strings.Join(flightList, ","))
}

// 题目描述
// 围棋棋盘由纵横各19条线垂直相交组成， 棋盘上一共119×19=361个交点， 对弈双方一方执白棋， 一方执黑棋， 落子时只能将棋子置于交点上。
// “气”是围棋中很重要的一个概念， 某个棋子有几口气， 是指其上下左右方向四个相邻的交叉点中， 有几个交叉点没有棋子， 由此可知：
// 1.在棋盘的边缘上的棋子最多有3口气 (黑1) ，在棋盘角点的棋子最多有2口气 (黑2) ，其他情况最多有4口气 (白1)
// 2.所有同色棋子的气之和叫做该色棋子的气， 需要注意的是， 同色棋子重合的气点， 对于该颜色棋子来说， 只能计算一次气， 比如下图中， 黑棋一共4口气， 而不是5口气， 因为黑1和黑2中间红色三角标出来的气是两个黑棋共有的， 对于黑棋整体来说只能算一个气。
// 3.本题目只计算气， 对于眼也按气计算， 如果您不清楚“眼”的概念， 可忽略， 按照前面描述的规则计算即可。
// 现在， 请根据输入的黑棋和白棋得到坐标位置， 计算黑棋和白棋一共各有多少气?

const BOARD_SIZE = 19

// 判断坐标是否在棋盘内
func isValid(x, y int) bool {
	return x >= 0 && x < BOARD_SIZE && y >= 0 && y < BOARD_SIZE
}

// 计算某种颜色棋子的气数
func calculateGas(chessPieces [][]int) int {
	gasCount := 0
	visitedGas := make(map[string]bool)
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, piece := range chessPieces {
		x, y := piece[0], piece[1]
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if isValid(newX, newY) && !isInPieces(chessPieces, newX, newY) && !visitedGas[fmt.Sprintf("%d,%d", newX, newY)] {
				gasCount++
				visitedGas[fmt.Sprintf("%d,%d", newX, newY)] = true
			}
		}
	}
	return gasCount
}

// 判断坐标是否在棋子列表中（即是否有棋子在该坐标上）
func isInPieces(chessPieces [][]int, x, y int) bool {
	for _, piece := range chessPieces {
		if piece[0] == x && piece[1] == y {
			return true
		}
	}
	return false
}

type Point struct {
	X int
	Y int
}

func calculateGasTest() {
	// blackm := make(map[Point]bool)
	// 构造blackm

	// 遍历black，获得4个，并去重。

	// 遍历withem，去除blackm

	// whitem := make(map[Point]bool)

	var blackPiecesStr, whitePiecesStr string
	fmt.Println("请输入黑棋坐标位置（格式如：[(1,2),(3,4)]）：")
	fmt.Scanln(&blackPiecesStr)
	fmt.Println("请输入白棋坐标位置（格式如：[(5,6),(7,8)]）：")
	fmt.Scanln(&whitePiecesStr)

	blackPieces := parsePieces(blackPiecesStr)
	whitePieces := parsePieces(whitePiecesStr)

	blackGas := calculateGas(blackPieces)
	whiteGas := calculateGas(whitePieces)

	fmt.Printf("黑棋的气数为：%d\n", blackGas)
	fmt.Printf("白棋的气数为：%d\n", whiteGas)
}

// 解析输入的棋子坐标字符串为二维切片
func parsePieces(input string) [][]int {
	input = input[1 : len(input)-1]
	pieces := [][]int{}
	if input == "" {
		return pieces
	}
	parts := strings.Split(input, "),(")
	for _, part := range parts {
		coords := strings.Split(part, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		pieces = append(pieces, []int{x, y})
	}
	return pieces
}

// 题目描述
// 小华按照地图去寻宝， 地图上被划分成m行和n列的方格， 横纵坐标范围分别是[0,n-1]和[0,m-1],
// 在横坐标和纵坐标的数位之和不大于k的方格中存在横金 (每个方格中仅存在一克黄金)，但横坐标和纵坐标数位之和大于k 的方格存在危险不可进入。小华从入口(0.0)进入， 任何时候只能向左， 右， 上，
// 格。
// 请问小华最多能获得多少克黄金?
// 输入描述
// 坐标取值范围如下：
// ⋅0≤m≤50
// ⋅0≤n≤50
// k的取值范围如下：
// ⋅0≤k≤100
// 输入中包含3个字数， 分别是m，n，k
// 输出描述
// 输出小华最多能获得多少克黄金

// digitSum计算一个整数的数位之和
func digitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// inRange判断坐标是否在方格范围内
func inRange(x, y, m, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

// dfs进行深度优先搜索来统计可获取的黄金数量
func dfs(x, y, m, n, k int, visited [][]bool) int {
	if !inRange(x, y, m, n) || visited[y][x] || digitSum(x)+digitSum(y) > k {
		return 0
	}
	visited[y][x] = true
	gold := 1
	for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		newX, newY := x+dir[0], y+dir[1]
		gold += dfs(newX, newY, m, n, k, visited)
	}
	return gold
}

func dfsTest() {
	var m, n, k int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	fmt.Scanln(&k)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	result := dfs(0, 0, m, n, k, visited)
	fmt.Println(result)
}

// 题目描述
// 石头剪刀布游戏有3种出拳形状： 石头、剪刀、布。分别用字母A，B，C表示。
// 游戏规则：
// 出拳形状之间的胜负规则如下：A>B;B>C;C>A;">"左边一个字母， 表示相对优势形状。右边一个字母， 表示相对劣势形状。当本场次中有且仅有一种出拳形状优于其它出拳形状， 则该形状的玩家是胜利者。否则认为是平局。
// 当发生平局， 没有赢家。有多个胜利者时， 同为赢家。
// · 例如1： 三个玩家出拳分别是A，B，C， 由于出现三方优势循环(即没有任何一方优于其它出拳者)，判断为平局。
// · 例如2：两个玩家， 出拳分别是A，B， 出拳A的获胜。
// · 例如3： 三个玩家， 出拳全部是A， 判为平局。
// 输入描述
// 在一场游戏中， 每个玩家的信息为一行。玩家数量不超过 1000 。每个玩家信息有 2 个字段， 用空格隔开：
// 1.玩家ID： 一个仅由英文字母和数字组成的字符串
// 2.出拳形状： 以英文大写字母表示， A、B、C形状。 例：

// judgeGame根据玩家出拳情况判断胜负并返回胜利者ID列表
func judgeGame() []string {
	var n int
	fmt.Scanln(&n)
	punchCount := make(map[string]int)
	players := make(map[string]string)
	for i := 0; i < n; i++ {
		var playerID, punch string
		fmt.Scanf("%s %s", &playerID, &punch)
		players[playerID] = punch
		punchCount[punch]++
	}
	winnerIDs := []string{}
	if len(punchCount) == 3 {
		winnerPunch := ""
		for _, p := range []string{"A", "B", "C"} {
			if punchCount[p] == 1 {
				winnerPunch = p
				break
			}
		}
		if winnerPunch != "" {
			for playerID, punch := range players {
				if punch == winnerPunch {
					winnerIDs = append(winnerIDs, playerID)
				}
			}
		}
	} else if len(punchCount) == 2 {
		for _, p := range []string{"A", "B", "C"} {
			if punchCount[p] == 1 {
				for playerID, punch := range players {
					if punch == p {
						winnerIDs = append(winnerIDs, playerID)
					}
				}
			}
		}
	}
	return winnerIDs
}

func judgeGameTest() {
	winnerIDs := judgeGame()
	if len(winnerIDs) > 0 {
		fmt.Printf("胜利者的ID为：%s\n", strings.Join(winnerIDs, ", "))
	} else {
		fmt.Println("本场游戏为平局")
	}
}

// 题目描述
// 输入字符串s， 输出s中包含所有整数的最小和。
// 说明：
// 字符串s， 只包含a-zA-Z±;
// 合法的整数包括
// · 正整数一个或者多个0-9组成, 如02 3 002 102
// · 负整数负号 (-)开头， 数字部分由一个或者多个0-9组成， 如·-0-012-23-00023
// 输入描述
// 包含数字的字符串
// 输出描述
// 所有整数的最小和

// MinSumOfIntegers计算字符串中所有整数的最小和
func MinSumOfIntegers(s string) int {
	sum := 0
	r := regexp.MustCompile(`-?\d+`)
	integers := r.FindAllString(s, -1)
	for _, numStr := range integers {
		num, _ := strconv.Atoi(numStr)
		sum += num
	}
	return sum
}

func MinSumOfIntegersTest() {
	var s string
	fmt.Scanln(&s)
	result := MinSumOfIntegers(s)
	fmt.Println(result)
}

// 题目描述
// 给一个正整数数列 nums, 一个跳数 jump, 及幸存数量 left。
// 运算过程为： 从索引0的位置开始向后跳， 中间跳过J个数字， 命中索引为J+1的数字， 该数被敲出， 并从该点起跳， 以此类推， 直到幸存 left 个数为止， 然后返回幸存数之和。
// 约束：
// • 0是第一个起跳点
// • 起跳点和命中点之间间隔 jump个数字， 已被敲出的数字不计入在内。
// • 跳到末尾时无缝从头开始 (循环查找) ，并可以多次循环。
// • 若起始时\mathsf left>len(nums)则无需跳数处理过程。

// surviveSum计算幸存数之和
func surviveSum(nums []int, jump int, left int) int {
	if left > len(nums) {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}
	// 标记哪些数字已经被敲出
	knockedOut := make(map[int]bool)
	currentIndex := 0
	count := len(nums)
	for count > left {
		// 计算下一个要敲出的索引，处理循环情况
		targetIndex := (currentIndex + jump + 1) % len(nums)
		// 找到下一个未被敲出的有效索引
		for knockedOut[targetIndex] {
			targetIndex = (targetIndex + 1) % len(nums)
		}
		knockedOut[targetIndex] = true
		count--
		currentIndex = targetIndex
	}
	sum := 0
	for index := range nums {
		if !knockedOut[index] {
			sum += nums[index]
		}
	}
	return sum
}

func surviveSumTest() {
	nums := []int{1, 2, 3, 4, 5}
	jump := 2
	left := 2
	result := surviveSum(nums, jump, left)
	fmt.Println(result)
}

// 题目描述
// 给定一个二叉树， 每个节点上站一个人， 节点数字表示父节点到该节点传递悄悄话需要花费的时间。
// 初始时， 根节点所在位置的人有一个悄悄话想要传递给其他人， 求二叉树所有节点上的人都接收到悄悄话花费的时间。
// 输入描述
// 给定二叉树
// 0920-1-1157-1-1-1-132
// 注：-1表示空节点

func maxTime(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归计算左子树最长路径花费时间
	leftMaxTime := maxTime(root.Left)
	// 递归计算右子树最长路径花费时间
	rightMaxTime := maxTime(root.Right)
	// 选择左右子树中较长的路径，加上根节点传递到子节点的时间（根节点自身的Val）
	return max2(leftMaxTime, rightMaxTime) + root.Val
}

// func max2(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func buildTree2(nodes []int, index int) *TreeNode {
	if index >= len(nodes) || nodes[index] == -1 {
		return nil
	}
	root := &TreeNode{
		Val:   nodes[index],
		Left:  buildTree2(nodes, 2*index+1),
		Right: buildTree2(nodes, 2*index+2),
	}
	return root
}

func maxTimeTest() {
	nodes := []int{0, 9, 20, -1, -1, 15, 7, -1, -1, -1, -1, 3, 2}
	root := buildTree2(nodes, 0)
	result := maxTime(root)
	fmt.Println(result)
}

// 题目描述
// 给定一段“密文”字符串s， 其中字符都是经过“密码本”映射的， 现需要将“密文”解密并输出。
// 映射的规则(`a'∼i')分别用(`1'∼'9')表示；(ij'∼'z')分别用 (“10*”~“26*”) 表示。
// 约束： 映射始终唯一。
// 输入描述
// “密文”字符串
// 输出描述
// 明文字符串
// 备注： 翻译后的文本长度在100以内

// decrypt将密文解密为明文
func decrypt(s string) string {
	var result string
	i := 0
	for i < len(s) {
		if i+1 < len(s) && s[i] >= '1' && s[i] <= '2' && s[i+1] == '*' {
			num := int(s[i]-'0')*10 + int(s[i+1]-'*')
			result += string('j' + num - 10)
			i += 2
		} else if s[i] >= '1' && s[i] <= '9' {
			result += string('a' + int(s[i]-'1'))
			i++
		}
	}
	return result
}

func decryptTest() {
	var s string
	fmt.Scanln(&s)
	plaintext := decrypt(s)
	fmt.Println(plaintext)
}

// 题目描述
// 疫情期间， 小明隔离在家， 百无聊赖， 在纸上写数字玩。他发明了一种写法：
// 给出数字个数n和行数nm(0<n≤999,0<m≤999),， 从左上角的1开始， 按照顺时针螺旋向内写方式， 依次写出2，3…n， 最终形成一个m行矩阵。
// 小明对这个矩阵有些要求：
// · 每行数字的个数一样多
// · 列的数量尽可能少
// · 填充数字时优先填充外部
// · 数字不够时， 使用单个*号占位
// 输入描述
// 输入一行， 两个整数， 空格隔开， 依次表示n、m
// 输出描述
// 符合要求的唯一矩阵


// func 