package main

import (
	"fmt"
	"strings"
)

func main() {
	res1 := GetResult(6)
	fmt.Println(res1)
	res2 := GetResult(8)
	fmt.Println(res2)
	res3 := GetResult(5)
	fmt.Println(res3)
	res5 := GetResult(9)
	fmt.Println(res5)
	res6 := GetResult(24)
	fmt.Println(res6)

	s1 := "A man, a plan, a canal: Panama"
	r1 := isPalindrome(s1)
	fmt.Println(r1)
}

func GetResult(n int) string {
	if n <= 0 {
		return "input number error"
	}
	var res string
	num := n       // 中间变量保存要分解的正整数
	var i int = 2  // 因子从2开始
	for num != 1 { // 如果最后分解到的因子为1就结束分解
		if num%i == 0 { // 判断当前分解后的商是否还有因子可以分解
			res = res + "*" + fmt.Sprintf("%d", i) // 把分解后的因子通过字符串连接
			num /= i                               // 计算下一个分解后的商，留作下一次继续运算
		} else { // 不能分解，带分解因子自增
			i++
		}
	}
	// 返回的时候，把最前面的"*"删去，并带上"="
	return strings.TrimLeft(res, "*") + fmt.Sprintf("=%d", n)
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	validString := str2ValidString(s)
	fmt.Printf("origin string:%s, transfer result:%s\n", s, validString)
	for i := 0; i < len(validString)/2; i++ {
		if validString[i] != validString[len(validString)-1-i] {
			return false
		}
	}
	return true
}

func str2ValidString(s string) string {
	var validCheckString = "abcdefghijklmnopqrstuvwxyz 0123456789" // 回文中的要检查的有效的字符串
	var validString string                                         // 将传入的字符串转为有效字符串，去掉非法字符，如":"
	for _, c := range s {
		lowerStr := strings.ToLower(string(c))
		if strings.Contains(validCheckString, lowerStr) {
			validString += lowerStr
		}
	}
	return validString
}
