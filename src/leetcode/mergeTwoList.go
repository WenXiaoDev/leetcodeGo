package leetcode

/*
 * LeetCode No. 21
 * merge two sorted linked list.
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val:0,Next:nil}
    pre := dummy
    p1,p2 := l1,l2
    for p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            pre.Next = p1
            pre = p1
            p1 = p1.Next
        } else {
            pre.Next = p2
            pre = p2
            p2 = p2.Next
        }
    }
    if p1 == nil {
        pre.Next = p2
    } else if p2 == nil {
        pre.Next = p1
    }
    return dummy.Next
}