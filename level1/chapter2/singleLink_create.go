package main

import (
	"fmt"
)

type LinkNode struct {
	Val  interface{}
	Next *LinkNode
}

type SingleLink struct {
	Head *LinkNode
	Size int
}

func Constructor() SingleLink {
	head := &LinkNode{nil, nil}

	return SingleLink{Head: head, Size: 0}
}

//头插法
func (this *SingleLink) AddHead(data interface{}) {
	newnode := &LinkNode{Val: data, Next: nil}
	newnode.Next = this.Head.Next
	this.Head.Next = newnode
	this.Size++
}

//尾插法
func (this *SingleLink) AddTail(data interface{}) {
	newnode := &LinkNode{Val: data, Next: nil}
	cur := this.Head
	//找到最后一个结点
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newnode
	newnode.Next = nil
	this.Size++
}

//获取第i位的值
func (this *SingleLink) Get(index int) interface{} {
	if index < 0 || index > this.Size-1 {
		return -1
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

//删除第i位结点
func (this *SingleLink) DeleteIndex(index int) {
	if index < 0 || index > this.Size-1 {
		fmt.Println("删除失败")
		return
	}
	cur := this.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.Size--
}

func (this *SingleLink) Display() {
	cur := this.Head
	for cur.Next != nil {
		fmt.Print(cur.Next.Val)
		cur = cur.Next
	}
}

func main() {
	obj := Constructor()
	fmt.Printf("type=%T", obj)
	a := []int{1, 2, 8, 4, 5}
	for i := 0; i < len(a); i++ {
		obj.AddHead(a[i])
	}
	// fmt.Println(obj.Get(3))
	obj.Display()

	// obj.AddTail(2)
	// obj.AddIndex(0, 0)
	// obj.DeleteIndex(0)
	// time.Sleep(1)
	fmt.Println(obj.Get(3))

}
