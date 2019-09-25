package leetcode

/*
LeetCode No 61
rotate list right
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

func rotateRight(head *ListNode, k int) *ListNode {
	size := 0
	dummy := &ListNode{Val: 0, Next: head}
	pre, end := dummy, dummy
	for end.Next != nil {
		end = end.Next
		size++
	}
	if size == 0 {
		return head
	}
	k_ := size - k%size
	for i := 0; i < k_; i++ {
		pre = pre.Next
	}
	end.Next = head
	head = pre.Next
	pre.Next = nil
	return head
}
