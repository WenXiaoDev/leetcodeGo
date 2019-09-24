package leetcode

/*
 * LeetCode No. 19
 * remove 
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    pre,end := head,head
    dummy := &ListNode{Val:0,Next:head}
    index := 1
    for index < n {
        end = end.Next
        index++
    }
    for end.Next != nil {
        dummy = pre
        pre,end = pre.Next,end.Next
    }
    dummy.Next = pre.Next
    if pre == head {
        return dummy.Next
    }
    return head
}