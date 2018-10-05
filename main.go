package main

import (
	"fmt"
	"gocode/datastructure"
)

func main() {
	//var l  = datastructure.NewLinkList()
	//for i := 0; i < 10; i++ {
	//	n := new(datastructure.Node)
	//	n.Data = i
	//	l.AppendTail(n)
	//}
	//l.Insert(1, &datastructure.Node{10, nil, nil})
	//for i := 0; i < l.Size(); i++ {
	//	fmt.Print(l.GetNode(i)," ")
	//	fmt.Printf("%p:\n",l.GetNode(i))
	//}
	//var s = datastructure.NewStack(10)
	//s.Push(1)
	//s.Push(2)
	//s.Push(3)
	//s.Push("s")
	//fmt.Println(s.Len())
	//fmt.Println(s.Peek())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Peek())
	//var q=datastructure.NewQueue(10)
	//q.Push(1)
	//q.Push(2)
	//fmt.Println(q.Pop())
	//fmt.Println(q.Pop())
	bt := datastructure.NewBinaryTree(1)
	bt.Left = datastructure.NewBinaryTree(2)
	bt.Right = datastructure.NewBinaryTree(3)
	bt.Left.Left = datastructure.NewBinaryTree(4)
	bt.Right.Left = datastructure.NewBinaryTree(5)
	fmt.Println("先序遍历")
	bt.Pre()
	fmt.Println()
	fmt.Println("中序遍历")
	bt.Middle()
	fmt.Println()
	fmt.Println("后序遍历")
	bt.End()
}
