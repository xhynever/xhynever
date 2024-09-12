package main

import (
	"flag"
	"fmt"
)

const Pi float64 = 3.14159265358979323846
const zero = 0.0

// func add(args ...{}int) int {
// 	sum := 0
// 	for _, arg := range args {
// 		sum += arg
// 	}
// 	return sum
// }

// func (s *Slice) Remove(value interface{}) error {
// 	for i, v := range *s {
// 		if isEqual(value, v) {
// 			*s = append((*s)[:i], (*s)[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return ERR_ELEM_NT_EXIST
// }

type ArrayStack struct {
	data []interface{}
}

func NewStack() *ArrayStack {
	return &ArrayStack{}
}

type Stack interface {
	Push(e interface{})
	Pop() (interface{}, error)
	Top() (interface{}, error)
	// IsEmpty() bool

}

func (s *ArrayStack) Push(e interface{}) {
	s.data = append(s.data, e)
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, nil
	}
	e := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return e, nil
}

func (s *ArrayStack) Top() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, nil
	}
	return s.data[len(s.data)-1], nil
}
func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

var (
	appPath string
	// configDir    = flag.String("configs", osx.GetEnv("CATEGRAF_CONFIGS", "conf"), "Specify configuration directory.(env:CATEGRAF_CONFIGS)")
	debugMode    = flag.Bool("debug", false, "Is debug mode?")
	debugLevel   = flag.Int("debug-level", 0, "debug level")
	testMode     = flag.Bool("test", false, "Is test mode? print metrics to stdout")
	interval     = flag.Int64("interval", 0, "Global interval(unit:Second)")
	showVersion  = flag.Bool("version", false, "Show version.")
	inputFilters = flag.String("inputs", "", "e.g. cpu:mem:system")
	install      = flag.Bool("install", false, "Install categraf service")
	remove       = flag.Bool("remove", false, "Remove categraf service")
	start        = flag.Bool("start", false, "Start categraf service")
	stop         = flag.Bool("stop", false, "Stop categraf service")
	status       = flag.Bool("status", false, "Show categraf service status")
	update       = flag.Bool("update", false, "Update categraf binary")
	updateFile   = flag.String("update_url", "", "new version for categraf to download")
)

func main() {

	var a []int
	for _, b := range a {
		fmt.Println("hello", b)
	}
	fmt.Println("nihao ")
	// s := NewStack()
	// fmt.Println(s.IsEmpty())
	// fmt.Println(s.IsEmpty())
	// fmt.Println(s.IsEmpty())
	// s.Push("a")
	// s.Push("b")
	// s.Push(3)

	// a, _ := s.Pop()
	// b, _ := s.Pop()
	// c, _ := s.Pop()
	// fmt.Println(a)
	// // fmt.Println("")
	// fmt.Println(b)
	// fmt.Println(c)
	// s := make([]int,0)
	// a := []int64{1, 2, 3}
	// var flag bool
	// if flag == true {

	// }
	// x := []int{1, 2, 3, 4, 5, 6,}
	// b := add(a...)

	// fmt.Println(b)
	// fmt.Printf("%T\n", zero)
	// fmt.Println(reflect.TypeOf(zero))
	// fmt.Println(reflect.TypeOf(zero))

	// 共享内存切片

	// a := []byte("AAA/BBB")
	// index := bytes.IndexByte(a, '/')
	// c := a[index+1:]
	// b := a[:index]
	// b = append(b, "CCC11"...)
	// fmt.Println(string(a), string(b), string(c))

	// fmt.Println(Pi)
	// 匿名函数执行后变量效果

	// var c int
	// c = 2
	// func(d int) {
	// 	var b = 1
	// 	fmt.Println(b)
	// 	c = d

	// }(3)
	// fmt.Println(c)

	// 写一个http服务
	// http.HandleFunc("/", HelloHandler)
	// http.ListenAndServe(":8000", nil)
	// reflect.ValueOf("Hello")

}

// func HelloHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Hello World")

// }

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
//  }

//  func main () {
// 	http.HandleFunc("/", HelloHandler)
// 	http.ListenAndServe(":8000", nil)

//  }
