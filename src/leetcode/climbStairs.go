package leetcode
/*
LeetCode No 70
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
*/
func climbStairs(n int) int {
    var tab = []int{1,1}
    for i := 2; i <= n; i++ {
        tab = append(tab,tab[i-1]+tab[i-2])
    }
    return tab[len(tab)-1]
}