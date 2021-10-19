package palindrome

import (
	"strconv"
	"testing"
)

// 判断一个整数是否是回文数,
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数
// 首先分析题意，只要正序读与反序读时一样的，那就是回文数，那么负数就可以排除掉
// 首先我们能想到的就是把整数转成字符串，然后比较字符串的第一位与最后一位、第二位与倒数第二位。。。以此类推，只要有不一样的就返回false

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	t.Log("x is 1: ", isPalindrome(1))
	t.Log("x is 11: ", isPalindrome(11))
	t.Log("x is 12: ", isPalindrome(12))
	t.Log("x is 123: ", isPalindrome(123))
	t.Log("x is 121: ", isPalindrome(121))
	t.Log("x is 1221: ", isPalindrome(1221))
	t.Log("x is 1231: ", isPalindrome(1231))
}
