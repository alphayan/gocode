package datastructure

import (
	"fmt"
)

type Tree struct {
	data     interface{}
	children *Tree
}

type BinaryTree struct {
	data        interface{}
	Left, Right *BinaryTree
}

func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{data: v}
}
func (b *BinaryTree) print() {
	fmt.Print(b.data, "  ")
}
func (b *BinaryTree) Get() interface{} { return b.data }
func (b *BinaryTree) Set(v interface{}) {
	if b == nil {
		fmt.Println("err:nil")
		return
	}
	b.data = v
}
func (b *BinaryTree) Pre() {
	if b == nil {
		return
	}
	b.print()
	b.Left.Pre()
	b.Right.Pre()
}
func (b *BinaryTree) Middle() {
	if b == nil {
		return
	}
	b.Left.Middle()
	b.print()
	b.Right.Middle()
}
func (b *BinaryTree) End() {
	if b == nil {
		return
	}
	b.Left.End()
	b.Right.End()
	b.print()
}
