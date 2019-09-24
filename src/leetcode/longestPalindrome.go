package leetcode
/*
LeetCode No 5
最长回文子串
*/

//1. 动态规划 dynamic programming
import "fmt"
func lngstPalnd(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}
	var tab = make([][]int,size)
	start,stop := 0,0
	for i :=0; i < size; i++ {
		tab[i] = make([]int,size)
		tab[i][i] = 1
	}
	for i := 0; i < size-1; i++ {
		if s[i] == s[i+1] {
			tab[i][i+1] = 2
			start,stop = i,i+1
		} else {
			tab[i][i+1] = 1
		}
	}
	for i := 2; i < size; i++ {
		for j := 0; j < size-i; j++ {
			x,y := j,i+j
			if s[x] == s[y] && tab[x+1][y-1] == y-x-1 {
				tab[x][y] = y-x+1
				start,stop = x,y
			} else if tab[x][y-1] >= tab[x+1][y] {
				tab[x][y] = tab[x][y-1]
			} else {
				tab[x][y] = tab[x+1][y]
			}
		}
	}
	return s[start:stop+1]
}

//2. 马拉车算法 Manacher Algorithm
func Manacher(s string) string {
	if len(s) <= 1 {
		return s
	}
	temp := make([]uint8,len(s)*2+3)
	Len := make([]int,len(s)*2+3)
	// Init temp
	temp[0],temp[len(temp)-2],temp[len(temp)-1] = '@','#','$'
	for i := 0; i < len(s); i++ {
		temp[2*i+1],temp[2*i+2] = '#',s[i]
	}
	// Manacher process https://blog.csdn.net/dyx404514/article/details/42061017
	var l,r,ans,offSet,p0,p int
	for i := 1; i < len(s)*2+1; i++ {
		if i < p && p-i+1 > Len[2*p0-i] {
			Len[i] = Len[2*p0-i]
			continue
		}
		if i < p {
			offSet,Len[i] = p-i+1,p-i+1
		} else {
			offSet,Len[i] = 1,1
		}
		for temp[i+offSet] == temp[i-offSet] {
			Len[i] += 1
			offSet++
		}
		p0,p = i,Len[i]+i-1
		if Len[i]-1 > ans {
			ans = Len[i]-1
			l,r = 2*p0-p,p
		}
	}
	var result []uint8
	for i := l; i <= r; i++ {
		if temp[i] != uint8('#') {
			result = append(result,temp[i])
		}
	}
	return string(result)
}

func Test5() {
	str := []string{"caba","babad","ab","aa","aba","a"}
	for _,v := range str {
		s := Manacher(v)
		fmt.Println(s)
	}
}