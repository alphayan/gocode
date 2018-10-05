package datastructure

type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

type LinkList struct {
	head *Node
	tail *Node
	size int
}

type Linklister interface {
	AppendHead(n *Node) bool
	GetNode(i int) *Node
	Insert(i int, n *Node) bool
	DeleteNode(i int) bool
	AppendTail(n *Node) bool
	Clear()
	Size() int
	Reverse()
}

func NewNode(value interface{}) *Node {
	return &Node{Data: value}
}

func NewLinkList() Linklister { return &LinkList{nil, nil, 0} }

func (n *Node) Location() *Node {
	return n
}

//头插法
func (l *LinkList) AppendHead(n *Node) bool {
	if n == nil {
		return false
	}
	if l.size == 0 {
		l.tail = n
	} else {
		n.Next = l.head
		l.head.Prev = n
	}
	l.head = n
	l.size++
	return true
}

//获取节点
func (l *LinkList) GetNode(i int) *Node {
	if i > l.size {
		return nil
	}
	n := l.head
	for j := 0; j < i; j++ {
		n = n.Next
	}
	return n
}

//插入
func (l *LinkList) Insert(i int, n *Node) bool {
	if i >= l.size || i <= 0 {
		return false
	}
	//头节点
	h := l.head
	//定义一个空节点
	var b *Node
	for j := 0; j < i; j++ {
		//节点左移
		b, h = h, h.Next

	}
	//空节点的下一个指向n
	b.Next = n
	n.Prev = b
	n.Next = h
	h.Prev = n
	l.size++
	return true
}

//删除
func (l *LinkList) DeleteNode(i int) bool {
	if i >= l.size || i <= 0 {
		return false
	}
	n := l.head
	var b *Node
	for j := 0; j < i; j++ {
		b, n = n, n.Next
	}
	b.Next = n.Next
	n = nil
	l.size--
	return true
}

//尾插法
func (l *LinkList) AppendTail(n *Node) bool {
	if n == nil {
		return false
	}
	if l.size == 0 {
		l.head = n
	} else {
		l.tail.Next = n
		n.Prev = l.tail
	}
	l.tail = n
	l.size++
	return true
}
func (l *LinkList) Size() int { return l.size }

func (l *LinkList) Clear() {
	for l.head.Next != nil {
		l.head = l.head.Next
	}
	l.head = nil
	l.tail = nil
	l.size = 0
}
func (l *LinkList) Reverse() {
	if l.tail == nil {
		return
	}
	nl := NewLinkList()
	for i := 0; i < l.size; i++ {
		tem := new(Node)
		tem.Data = l.GetNode(i).Data
		nl.AppendHead(tem)
	}
	//不能直接给地址赋值
	*l = *nl.(*LinkList)
}
