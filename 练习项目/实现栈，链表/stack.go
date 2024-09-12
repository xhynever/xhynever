package stack

// 测试
// s := New()
//     s.Push(1)
//     s.Push(2)
//     s.Push(3)

//     fmt.Println(s)
//     fmt.Println(s.Pop())
//     fmt.Println(s)
//     fmt.Println(s.Top())

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
	if s.IsEmpty() {
		return nil, nil
	}
	e := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return e, nil
}
func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *ArrayStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, nil
	}
	return s.data[len(s.data)-1], nil
}

// func (s *ArrayStack) IsEmpty() bool {
// 	return len(s.data) == 0
// }

type LinkedStack struct {
	topPtr *node
	size   int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}
