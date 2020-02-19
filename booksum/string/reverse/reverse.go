package main

import(
	"fmt"
	"strings"
)
func main(){
	// golang中对字符串的遍历（注意：对中文字符串的处理）https://studygolang.com/articles/13639
	str1 := "hello 北京"
	str2 := []rune(str1)
	for i, ch := range str2 {
		fmt.Printf("str2[%d]=%c\n", i, ch)
	}

	str3 := "Welcome to beijing, I am in beijing"
	res := WordReverseInSentenceBySplit(str3)
	fmt.Printf("句子%s中的单词倒置后的结果：%s\n", str3, res)
}

func WordReverseInSentenceBySplit(txt string)string{
	txt = strReverse(txt)
	wordList := strings.Split(txt, " ")
	for i := range wordList{
		wordList[i] = strReverse&(wordList[i])
	}
 	result := strings.Join(wordList, " ")
	return result
}

func strReverse(str *string)string{
	i := 0  // i指向单词开头
	j := len(&str) - 1 // j指向单词末尾
	tmp := ""   // 中间变量，暂存要交换的字符
	if i <=j {
		tmp = str[i] 
		str[i] = str[j]
		str[j] = tmp
		i++
		j--
	}
	return &str
}