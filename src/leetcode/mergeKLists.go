package leetcode
/*
 * LeetCode No 23
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 1.两两合并，共合并k-1次
 // 时间复杂度 K^2*Lmax = K*N N:列表的总长度
 func mergeTwoList(l1,l2 *ListNode) *ListNode {
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

func mergeKLists(lists []*ListNode) *ListNode {
    var target *ListNode
    if len(lists) > 0 {
        target = lists[0]
    }
    for i := 1; i < len(lists); i++ {
        target = mergeTwoList(target,lists[i])
    }
    return target
}

// 2.分治 Divide and Conqure
// 也是先两两合并，将k个list合并为 k/2 个列表，然后再合并成 k/4 个，总共经过 logK 轮合并，
// 时间复杂度为 N*logK = K*logK*Lmax； Lmax 为所有列表中最长的长度。它是N的上限，也就是
// N 在最差条件下的值
func mergeKlistsDnC(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
	lst1,lst2 := 0,1
	for true{
		size := len(lists)
		if lst2 >= size {
			break
		}
		for lst2 < size {
			temp := mergeTwoList(lists[lst1],lists[lst2])
			lists = append(lists,temp)
			lst1,lst2 = lst1+2,lst2+2
		}
	}
	return lists[lst1]
}