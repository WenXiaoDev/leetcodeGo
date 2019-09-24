package leetcode
/**
 * Leetcode No 99
 * Solution:
 * 		First, iterate the tree in middle order,and sequencely save the itered node in a pointer array: arrMid.
 * 		Second, use two indexes i and j to iterate the array from left and right respectively.
 * 		If arrMid[i] >= arrMid[i+1], and arrMid[j] <= arrMid[j-1], swap their Value.
 *
 * Definition for a binary tree node:
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func middleOrderFill(root *TreeNode,arr *[]*TreeNode) {
    if root.Left != nil {
        middleOrderFill(root.Left,arr)
    }
    (*arr) = append((*arr),root)
    if root.Right != nil {
        middleOrderFill(root.Right,arr)
    }
}

func recoverTree(root *TreeNode)  {
    if root == nil {
        return
	}
    var midArr []*TreeNode
	middleOrderFill(root,&midArr)
    i,j := 0,len(midArr)-1
    for i < len(midArr)-1 {
        if midArr[i].Val >= midArr[i+1].Val {
            break
        }
        i++
    }
    for j > 0 {
        if midArr[j].Val <= midArr[j-1].Val {
            break
        }
        j--
    }
    midArr[i].Val,midArr[j].Val = midArr[j].Val,midArr[i].Val
}