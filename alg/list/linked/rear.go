package linked

import (
	"fmt"
	"go-study/alg/list/node"
)

// RearLinkedList 尾插链表定义
type RearLinkedList struct {
	Head *node.Single
	Rear *node.Single
	Len  int
}

// 遍历
func (l RearLinkedList) visit() {
	for p := l.Head.Next; p != nil; p = p.Next {
		fmt.Println(p.Data)
	}
}

// 插入新节点
func (l *RearLinkedList) insert(data node.EL) {
	newNode := node.CreateSingle(data)

	// 如果头结点next为空
	if l.Head.Next == nil {
		l.Head.Next = newNode
	}
	// 处理尾指针
	l.Rear.Next = newNode
	// 尾指针记录当前尾节点
	l.Rear = newNode

	l.Len++
}

func (l *RearLinkedList) example(n int) *RearLinkedList {
	for i := 0; i < n; i++ {
		l.insert(node.EL((i + 1) * (i + 1)))
	}
	return l
}

// 创建尾插法链表
func createRearLinkedList() *RearLinkedList {
	n := node.CreateSingle(0)
	return &RearLinkedList{
		Head: n,
		Rear: n,
		Len:  0,
	}
}
