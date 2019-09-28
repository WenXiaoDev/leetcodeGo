package leetcode
/*
LeetCode No 44
the realization of regular expression, '?' and '*'

*/
// 1. first method: 回溯法 超时了
func processP(p string) string {
    if p == "" {
        return p
    }
    result := []uint8{p[0]}
    i := 1
    for i < len(p) {
        if p[i] != '*' || p[i-1] != '*' {
            result = append(result,p[i])
        }
        i++
    }
    return string(result)
}
func match(s string, p string) bool {
    sizeS,sizeP := len(s),len(p)
	if sizeP == 0 {
		return sizeS == 0
	} else if sizeS == 0 {
        return p == "*"
	}
	if p[0] == '?' || s[0] == p[0] {
		return isMatch(s[1:],p[1:])
	} else if p[0] == '*' {
		for i := 0; i <= len(s); i++ {
			if isMatch(s[i:],p[1:]) {
				return true
			}
		}
	}
	return false
}
func IsMatch(s string, p string) bool {
    p = processP(p)
    return match(s,p)
}

//2 基于上述回溯法的动态规划
func isMatchDp(s,p string) bool {
	if p == "" {
        return s == ""
    }
    p = processP(p)
	dp := make([][]bool,len(p)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool,len(s)+1)
	}
	dp[len(p)][len(s)],dp[len(p)-1][len(s)] = true,p[len(p)-1] == '*'
	for i := len(s)-1; i >= 0; i-- {
		for j := len(p)-1; j >= 0; j-- {
			if p[j] == '?' || p[j] == s[i] {
				dp[j][i] = dp[j+1][i+1]
			} else if p[j] == '*' {
				dp[j][i] = dp[j][i+1]
			}
		}
	}
	return dp[0][0]
}

// 3. 双指针遍历
func isMatch3(s,p string) bool {
	start,match := -1,0
	i,j := 0,0
	lenS,lenP := len(s),len(p)
	for i < lenS {
		if j < lenP && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < lenP && p[j] == '*' {
			start = i
			j++
			match = j
		} else if start > -1 {
			start += 1
			i = start
			j = match
		} else {
			return false
		}
	}
	for j < lenP && p[j] == '*' {
		j++
	}
	return j == lenP
}