package datastructure

type Node struct {
	Data interface{}
	Next *Node
}

type LinkList struct {
	Head *Node
	Tail *Node
	size int
}

type Linklister interface {
	AppendHead(n *Node) bool
	GetNode(i int) *Node
	Insert(i int, n *Node) bool
	DeleteNode(i int) bool
	AppendTail(n *Node) bool
}

func NewLinkList() *LinkList { return &LinkList{nil, nil, 0} }

func (l *LinkList) AppendHead(n *Node) bool {
	if n == nil {
		return false
	}
	if l.size == 0 {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
	l.size++
	return true
}
func (l *LinkList) GetNode(i int) *Node {
	if i > l.size {
		return nil
	}
	n := l.Head
	for j := 0; j < i; j++ {
		n = n.Next
	}
	return n
}
func (l *LinkList) Insert(i int, n *Node) bool {
	return false
	return true
}
func (l *LinkList) DeleteNode(i int) bool {
	if i >= l.size || i < 0 {
		return false
	}
	n := l.Head
	var b *Node
	for j := 0; j <= i; j++ {
		b, n = n, n.Next
	}
	b.Next = n.Next
	n = nil
	l.size--
	return true
}
func (l *LinkList) AppendTail(n *Node) bool {
	if n == nil {
		return false
	}
	if l.size == 0 {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
	l.size++
	return true
}
func (l *LinkList) Size() int { return l.size }
