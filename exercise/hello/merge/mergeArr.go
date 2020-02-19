package main

import(
	"fmt"
	"sort"
	"bufio"
)

func main(){
	arrayA := []int{1, 3, 5,7, 9}
	arrayB := []int{2, 4, 6, 8, 10}
	m := merge(arrayA, arrayB)
	fmt.Println(m)

	var numList []int
	inputNum := getInput()
	for {
		numList = append(numList, inputNum)
		if inputNum == 0{
			break
		}
		inputStr = getInput
	}
	for _, num := range numList{
		n := divNum(num)
		fmt.Println(n)
	}
}

func merge(arrayA, arrayB []int) []int {
	lenA := len(arrayA) 
	lenB := len(arrayB)
	arrayMerge := []int{}
	
	for i := 0; i < lenA; i++ {
		arrayMerge = append(arrayMerge, arrayA[i])
	}
	
	for j := 0; j < lenB; j++ {
		arrayMerge = append(arrayMerge, arrayB[j])
   	}
	sort.Ints(arrayMerge)
	return arrayMerge  
}


fun divNum(n int)(div int){
	count := 1 
	if n = 1 {
		count = 1
		return 1
	}
	var i := 2
	if(i < 1500){
		if count < n {
			if i % 2 == 0 || i % 3 == 0 || i % 7 == 0 {
				count++
				i++ 
			}
		}else{
			div = i
			return div
		}
	}
	return div
}


func getInput() int {
	//使用os.Stdin开启输入流
	//函数原型 func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)
	//in.ReadLine函数具有三个返回值 []byte bool error
	//分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return num
}
