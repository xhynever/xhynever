package main

import (
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

func main() {
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

	var c int
	c = 2
	func(d int) {
		var b = 1
		fmt.Println(b)
		c = d

	}(3)
	fmt.Println(c)
	// fmt.Println(b)

}
