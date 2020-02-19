package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
	"container/list"
)

// 逆波兰表达式问题
// ```
// // 1.问题描述
// // 给定一个逆向波兰表达式，要求选择一种合适的数据结构并设计算法，计算出表达式的值。
// ```

// ```
// 逆波兰表达式满足条件：
// (1)表达式含有一个数字字符串或一串数据字符；
// (2)它满足给定格式，主要有数字和四则运算符号组成，如逆波兰表达式字符串"2,3,*,1,2,+,+"表示3*4+(1+2)
// ```

// 2.解题思路
// ```
// 处理逆向波兰表达式最合适的数据结构就是堆栈stack，使用该数据结构解决的步骤：
// (1)如果当前字符是数字，将数字压入堆栈；
// (2)如果当前字符是运算符，从堆栈中弹出两个数字，根据运算符做相应运算后，将结构压入堆栈；
// (3)当所有字符处理完毕后，堆栈中包含的数值就是表达式的值。
// ```



func main(){
	str1 := "123"
	str2 := "12.3"
	fmt.Println(unicode.IsDigit(str1))
	fmt.Println(unicode.IsDigit(str2))

	m := make(map[int]string)
	m[20200101] = "2020年1月1日"
	m[20200202] = "2020年2月2日"
	l := list.New()
	l.PushFront("first")	// 从列表前方插入字符串
	l.PushBack(2) 			// 从列表尾部插入整型2
	l.PushBack(m) 			// 从类别尾部插入map
	fmt.Println(l)
}

func reversePolishExpr(str string)int{
	strList := strings.Split(str, ",")
	numStack := []int{}   // 这里以整型数据的运算为例
	for i := range strList{
		if isOperator(fmt.Sprintf("%c", strList[i])){
			numFront := numStack[i]
			numTail := numStack[i-1]
			num1 , _ := strconv.Atoi(numFront)
			num2 , _ := strconv.Atoi(numTail)
			res := 0
			switch fmt.Sprintf("%c", strList[i]){
			case "+":
				res = num1 + num2
			case "-":
				res = num1 - num2
			case "*":
				res = num1 * num2 
			case "/":
				res = num1 / num2

			}
		}
		if unicode.IsDigit(fmt.Sprintf("%c", strList[i])){
			numStack = append(numStack, fmt.Split("%c", strList[i]))
		}

	}
}

func isOperator(str string)bool{ // 判断当前字符是否是运算符
	if len(str) == 0 || len(str) > 1{
		return false
	}
	if str == "+" || str == "-" || str == "*" || str == "/"{
		return true
	}
	return false
}


