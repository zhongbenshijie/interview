package ring

import (
	"testing"
)

// 判断单链表是否存在环是一个经典的快慢指针问题，
// 一个每次走一步的指针，和一个每次走两步的指针，如果链表里有环的话，两个指针最终肯定会相遇
// 具体实现如下：

type Node struct {
	Value    int
	NextNode *Node
}

func isHaveRing(head *Node) bool {
	slow := head
	fast := head

	// 这里第一个功能是判断链表的长度不能小于2, 判断顺序不能调换
	// 第二个功能是fast走到最后一个节点或者最后一个节点的NextNode时，如果走不下去了，就意味着不存在环，可以跳出循环，如果存在环会一直走下去，快慢指针终究会相遇
	for fast != nil && fast.NextNode != nil {
		slow = slow.NextNode
		fast = fast.NextNode.NextNode

		if slow == fast {
			return true
		}
	}

	return false
}

func TestIsHaveRing(t *testing.T) {
	node1 := &Node{Value: 1, NextNode: nil}
	node2 := &Node{Value: 2, NextNode: nil}
	node3 := &Node{Value: 3, NextNode: nil}
	node4 := &Node{Value: 4, NextNode: nil}
	node5 := &Node{Value: 5, NextNode: nil}

	node1.NextNode = node2
	node2.NextNode = node3
	node3.NextNode = node4
	node4.NextNode = node5

	t.Log(isHaveRing(node1))

	// 构造环
	node5.NextNode = node1
	t.Log(isHaveRing(node1))
}
