package reverse_list

import (
	"fmt"
	"testing"
)

// 反转单链表

type Node struct {
	Value    int
	NextNode *Node
}

// 这种是遍历链表反转的算法
// 思路就是pre指向cur的前一个节点，具体就是刚开始时, pre指向头节点，cur指向头节点的下一个节点
// 然后每次循环都将指向箭头逆转，先用temp保存cur的NextNode，然后把pre赋值给cur的NextNode，
// 因为temp暂时保存了cur的NextNode，所以cur的NextNode被覆盖没关系，这样指向箭头就被逆转了，
// 然后就是移动pre到cur，cur移动到temp，只要cur不为nil，就可以继续循环这个过程，
// 最后再将原来链表的第一个节点的NextNode设置为nil就可以，当然，函数外原来的headNode已经指向新
// 链表的最后一个节点了，此时pre指向新链表的第一个节点，把pre返回即可。
func reverseList(headNode *Node) *Node {
	pre := headNode
	cur := headNode.NextNode

	for cur != nil {
		temp := cur.NextNode
		cur.NextNode = pre
		pre = cur
		cur = temp
	}

	headNode.NextNode = nil

	return pre
}

func createList(headNode *Node, nodeNumber int) {
	cur := headNode

	for i := 1; i < nodeNumber; i++ {
		cur.NextNode = &Node{}
		cur.NextNode.Value = i
		cur.NextNode.NextNode = nil

		cur = cur.NextNode
	}
}

func printList(cur *Node) {
	for cur != nil {
		fmt.Print(*cur, " ")
		cur = cur.NextNode
	}
	fmt.Println()
}

func TestReverseList(t *testing.T) {
	headNode := &Node{Value: 0, NextNode: nil}
	nodeNumber := 6

	createList(headNode, nodeNumber)
	printList(headNode)

	headNode = reverseList(headNode)
	printList(headNode)
}
