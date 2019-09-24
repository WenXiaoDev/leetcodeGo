package leetcode
/*
KMP算法的关键思想：
	pattern: ABCDABD
	          ABCDABD
	假如在D处发生不匹配，意味着D之前的都是匹配的；按照暴力破解的思想，会把pattern向后移动一个位置再从头开始匹配；这么做的本质其实是
	把pattern的前缀ABCDA 与它的后缀BCDAB进行匹配，如果不匹配，则依此继续迭代；那么如果我们能提前得出pattern的前缀/后缀 对中相等的那些，
	我们就可以不用去比较那些一定不匹配的前缀/后缀对；
	因此，KMP算法的关键步骤是提前得出pattern的匹配部分的公共前缀/后缀对；
	例如，如果在D处不匹配，那么ABCDAB是匹配的，因此最长的公共前缀/后缀对是AB（前）和AB（后）
	因为已经知道这个信息，所以可以直接用AB（前）之后的C与当前不匹配处进行比较；
*/
func getNext(p string) []int {
	next := make([]int,len(p))
	next[0] = -1
	i,j := 0,-1
	for i < len(p)-1 {
		if j == -1 || p[i] == p[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func Kmp(s,p string) bool {
	return true
}

func OptKmp(s,p string) bool {
	return true
}