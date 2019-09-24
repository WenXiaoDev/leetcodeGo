package leetcode
/*
LeetCode No. 25
对于链表相关的问题，很重要的一个技巧是，在链表head前面添加虚拟的节点，用来辅助算法
*/
import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func reverse(head *ListNode) *ListNode {
	curr := head
	var pre *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

func reverseKGroup(head *ListNode,k int) *ListNode {
	dummy := ListNode{Val:0,Next:head}
	pre,end := &dummy,&dummy
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next =next
		pre = start
		end = pre
	}
	return dummy.Next
}

func Test25() {
	var head *ListNode
	for i := 10; i > 0; i-- {
		temp := ListNode{Val:i,Next:head}
		head = &temp
	}
	result := reverseKGroup(head,3)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}