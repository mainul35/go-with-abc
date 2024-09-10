package leetcode

// First, contat the 2 strings and check if their concatanated values are equal from forward and reverse
// If yes, we can check GCD

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	gcd := gcd_str(len(str1), len(str2))
	return str1[0:gcd]
}

func gcd_str(str1Len int, str2Len int) int {
	if str2Len == 0 {
		return str1Len
	} else {
		return gcd_str(str2Len, str1Len%str2Len)
	}
}
