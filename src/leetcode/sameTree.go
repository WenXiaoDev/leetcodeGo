package leetcode
/*
LeetCode No 100
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q || p.Val == q.Val {
		if isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right) {
			return true
		} else {
			return false
		}
	}
	return false
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return true
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	deque := [][]*TreeNode{{p,q}}
	var l,r *TreeNode
	for len(deque) != 0 {
		l,r,deque = deque[:1][0][0],deque[:1][0][1],deque[1:]
		if !check(l,r){
			return false
		}
		if l != nil {
			deque = append(deque,[]*TreeNode{l.Left,r.Left},[]*TreeNode{l.Right,r.Right})
		}
	}
	return true
}