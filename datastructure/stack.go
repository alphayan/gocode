package datastructure

type stack struct {
	size int           //容量
	data []interface{} //用切片data作容器，定义为interface{}类型的切片以接收任意类型
}

func NewStack(i int) *stack {
	return &stack{i, make([]interface{}, 0, i)}
}

//压出元素
func (s *stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	i := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return i
}

//压入元素
func (s *stack) Push(v interface{}) bool {
	if s.IsFull() {
		return false
	}
	if v == nil {
		return false
	}
	s.data = append(s.data, v)
	return true
}

//返回栈顶
func (s *stack) Peek() interface{} {
	return s.data[s.Len()-1]
}

func (s *stack) Size() int {
	return s.size
}

//容量=数据长度
func (s *stack) IsFull() bool {
	return s.size == len(s.data)
}

//数据长度=0
func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}
func (s *stack) Len() int {
	return len(s.data)
}
