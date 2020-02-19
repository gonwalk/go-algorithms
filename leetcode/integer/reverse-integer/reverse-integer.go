package main

import(
	"fmt"
	"strings"
)

// ## 1.题目描述
/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

链接：https://leetcode-cn.com/problems/reverse-integer
*/


// ## 2.实现方式一：求整取余倒置法
// 通过对10取余求得当前循环中的低位，配合对10求整及上一步求得的余数（并乘以10即低位每次向高位移动一位）可得倒置后的结果。
// 如658倒置后为856=(658%10)*10*10 + (658/10%10)*10 + 658/10/10
func reverse(x int) int {
    res := 0 
    for x != 0 {
        res = res * 10 + x % 10
        x = x / 10
        if(res < -(1<<31) || res > (1<<31 -1)){ // 判断是否溢出，其数值范围不在 [−231,  231 − 1]内则溢出，根据题目要求返回0
            return 0  // 发生溢出，倒置的结果返回0
        }
    }
    return res
}

func main(){
	n := 685
	s := " -123"
	fmt.Printf("%d倒置后值为%d\n", n, reverse(n))

	flag := 1
	s = strings.Trim(s, " ")
	if strings.HasPrefix(s, "-"){
		flag *= -1
		s = s[1:]
	}
	if strings.HasPrefix(s, "+"){
		flag *= 1
		s = s[1:]
	}
	for i := 0; i < len(s); i++{
		fmt.Println(string(flag) + string(s[i]))
	}
}

