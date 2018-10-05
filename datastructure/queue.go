package datastructure

type queue struct {
	size int
	data []interface{}
}

func NewQueue(i int) *queue {
	return &queue{i, make([]interface{}, 0, i)}
}

//压出元素
func (s *queue) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	i := s.data[0]
	s.data = s.data[1:s.Len()]
	return i
}

//压入元素
func (s *queue) Push(v interface{}) bool {
	if s.IsFull() {
		return false
	}
	if v == nil {
		return false
	}
	s.data = append(s.data, v)
	return true
}

func (s *queue) Size() int {
	return s.size
}

//容量=数据长度
func (s *queue) IsFull() bool {
	return s.size == len(s.data)
}

//数据长度=0
func (s *queue) IsEmpty() bool {
	return s.Len() == 0
}
func (s *queue) Len() int {
	return len(s.data)
}
