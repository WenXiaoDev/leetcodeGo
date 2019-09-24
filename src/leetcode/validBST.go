package leetcode
/*
leetCode No 98
validate a Binary Search Tree:
the value of nodes in left sub-tree must be smaller than the root node;
the value of nodes in right sub-tree must be bigger than the root node;
*/
type TreeNode struct {
    Val int
     Left *TreeNode
     Right *TreeNode
}

 const (
    INT64_MAX = int(^uint(0)>>1)
    INT64_MIN = ^INT64_MAX
)
// BFS
func isValidBST(root *TreeNode) bool {
    queue := []*TreeNode{root}
	bounds := [][]int{{INT64_MIN,INT64_MAX}}
	var current *TreeNode
	var low,high int
	for len(queue) != 0 {
		current,queue = queue[0],queue[1:]
		low,high,bounds = bounds[0][0],bounds[0][1],bounds[1:]
		if current == nil {
			continue
		}
		if current.Left != nil && (current.Left.Val >= current.Val || current.Left.Val <= low) {
			return false
		}
		if current.Right != nil && (current.Right.Val <= current.Val || current.Right.Val >= high) {
			return false
		}
		// queue = append(queue,current.Left,current.Right)
		// bounds = append(bounds,[]int{low,current.Val},[]int{current.Val,high})
        if current.Left != nil {
            queue = append(queue,current.Left)
            bounds = append(bounds,[]int{low,current.Val})
        }
        if current.Right != nil {
            queue = append(queue,current.Right)
            bounds = append(bounds,[]int{current.Val,high})
        }
	}
	return true
}
//DFS
func isBstDFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root}
	bounds := [][]int{{INT64_MIN,INT64_MAX}}
	var current *TreeNode
	var low,high int
	for len(stack) > 0 {
		top := len(stack)-1
		current,stack = stack[top],stack[:top]
		low,high,bounds = bounds[top][0],bounds[top][1],bounds[:top]
		if current == nil {
			continue
		}
		if current.Val <= low || current.Val >= high {
			return false
		}
		stack = append(stack,current.Right,current.Left)
		bounds = append(bounds,[]int{current.Val,high},[]int{low,current.Val})
	}
	return true
}

func middleFill(root *TreeNode,array *[]int) {
	if root.Left != nil {
		 middleFill(root.Left,array)
	}
	(*array) = append((*array),root.Val)
	if root.Right != nil {
		middleFill(root.Right,array)
	}
}

func isBSTmiddleFill(root *TreeNode) bool {
	if root == nil {
		return true
	}
	line := []int{}
	middleFill(root,&line)
	for i := 1; i < len(line); i++ {
		if line[i-1] >= line[i] {
			return false
		}
	}
	return true
}

