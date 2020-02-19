package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str1 := "aabcccde"
	str2 := "aaabcccaa"

	fmt.Println(fmt.Sprintf("%c", str1[0]))

	encodingStr1 := RLEncode(str1)
	encodingStr2 := RLEncode(str2)

	fmt.Println(encodingStr1)
	fmt.Println(encodingStr2)

	str3 := "6a3c"
	str4 := "11a1b3c2a"

	fmt.Println(fmt.Sprintf("%c", str3[0]))
	fmt.Println(isDigit(fmt.Sprintf("%c", str3[0])))

	decodingStr3 := RLEDecode(str3)
	fmt.Println(decodingStr3)
	decodingStr4 := RLEDecode(str4)
	fmt.Println(decodingStr4)
}

func RLEncode(str string)string{
	var encodingStr []string  // 记录所有的连续字符出现的次数及对应的该连续字符
	var count int = 0   // 记录当前连续字符出现的次数
	lastStr := str[0]   // 记录上一次出现的那个连续的字符
	for i := range str{
		if str[i] == lastStr{
			count++
		}else{
			encodingStr = append(encodingStr, fmt.Sprintf("%d", count)) // 将上一个连续字符出现的次数转为字符串记录在string切片中
			encodingStr = append(encodingStr, fmt.Sprintf("%c", lastStr))

			lastStr = str[i]   // 将与上一个字符不同的新出现的字符记录在最新的连续字符变量中
			count = 1          // 将新出现字符的次数置为1
		}
	}
	encodingStr = append(encodingStr, fmt.Sprintf("%d", count))   //将for循环结束后，最后出现的连续字符记录下来
	encodingStr = append(encodingStr, fmt.Sprintf("%c", lastStr))
	newStr := strings.Join(encodingStr, "")
	return newStr
}

func RLEDecode(str string)string{
	var decodeStr []string   // 存放单个的解析出来的字符（将相应的数字转为对应个数的字符串）
	num := 0 				 // 统计当前字符前面的数字所占位数（如“123a6i”中的字符a的num=3表示"123"占3位）
	i := 0

	if i < len(str){
		// if strings.ContainsAny("0123456789", fmt.Sprintf("%c", str[i])){
		if isDigit(fmt.Sprintf("%c", str[i])){
			i++
			num++
			fmt.Sprintf("当前字符%c", str[i])
		}else{ // 不是数字型字符，就开始把当前字符添加到字符串切片中
			numStr := str[(i-num):i]  //数字对应的字符串
			fmt.Println("字符前缀对应的数字：", numStr)
			num , _ := strconv.Atoi(numStr)
			for j:=0;j<num;j++{
				decodeStr = append(decodeStr, fmt.Sprintf("%c", str[i]))
			}
			i++
			num = 0
		}
	}
	return strings.Join(decodeStr, "")
}

func isDigit(str string) bool{
	if len(str) <= 0 {
		return false
	}
	num, err := strconv.Atoi(str)
	if err == nil && num > 0{
		return true
	}
	return false
}