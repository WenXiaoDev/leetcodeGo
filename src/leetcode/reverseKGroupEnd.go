package leetcode

/*
LeetCode No ??
the transformation of the topic "reverseKGroup", namely reversing list by K group each time;
this time, need to do the reverse from the end of the list.
The first method is recursive algorithm.

type ListNode struct {
	Val int
	Next *ListNode
}
*/
import (
	"fmt"
	"time"
)

func reverseK(dum *ListNode, k int) {
	pre, cur := dum.Next, dum.Next.Next
	for i := 1; i < k; i++ {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	dum.Next.Next = cur
	dum.Next = pre
}
func reverseKGroupEnd(head *ListNode, k int) {
	dummy := &ListNode{Val: 0, Next: head}
	pre, cur := dummy, head
	for i := 1; i < k; i++ {
		if cur == nil {
			return
		}
		cur = cur.Next
	}
	var recursiveRev func(pre, lat *ListNode) *ListNode
	recursiveRev = func(pre, lat *ListNode) *ListNode {
		if lat.Next == nil {
			reverseK(pre, k)
			return pre
		} else {
			if result := recursiveRev(pre.Next, lat.Next); result == lat {
				reverseK(pre, k)
				return pre
			} else {
				return result
			}
		}
	}
	recursiveRev(pre, cur)
}

func Test003() {
	head := &ListNode{Val: 0}
	cur := head
	for i := 1; i < 10; i++ {
		temp := &ListNode{Val: i}
		cur.Next = temp
		cur = cur.Next
	}
	fmt.Println("original list:")
	for cur = head; cur != nil; cur = cur.Next {
		fmt.Print(" ", cur.Val)
	}
	fmt.Println("\nreversed list:")
	start := time.Now().UnixNano()
	reverseKGroupEnd(head, 3)
	end := time.Now().UnixNano()
	for cur = head; cur != nil; cur = cur.Next {
		fmt.Print(" ", cur.Val)
	}
	fmt.Println("duration: ", end-start, " ns")
}
