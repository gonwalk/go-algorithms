package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Node struct {
}

func main() {

	fmt.Println("Hello, World!")

	//var i int
	str := "abcdefgh"
	for i := range str {
		fmt.Println(i)
	}

	var j int
	for ; j < 10; j++ {
		fmt.Println(j)
	}
	fmt.Println()
	str = strings.Trim(str, " ")

	m := make(map[int]string)
	m[20200101] = "2020年1月1日"
	m[20200202] = "2020年2月2日"
	l := list.New()
	l.PushFront("first") // 从列表前方插入字符串
	l.PushBack(2)        // 从列表尾部插入整型2
	l.PushBack(m)
	for i := l.Front(); i != nil; i = i.Next() { // （双端链表）列表的遍历
		fmt.Println(i.Value)
	} // 从类别尾部插入map

	path := "/a/../../b/../c//.//"
	pathList := strings.Split(path, "/")
	for i, item := range pathList {
		fmt.Printf("pathList[%d]=%s\n", i, item)
	}

	var bytes []byte
	var i byte
	for i = 0; i < 10; i++ {
		bytes = append(bytes, i)
	}
	var mp map[string]string
	for _, b := range bytes {
		fmt.Printf("map[string(%v)]=%v ", b, b)
		key := string(b)
		v, _ := mp[key]
		fmt.Printf("m[%v]=%v ", key, v)
	}

	fmt.Println()
	mm := make([]int, 0)
	for i := 0; i < 17; i++ {
		mm = append(mm, i)
	}
	fmt.Println("length=%d, cap=%d", len(mm), cap(mm))

	var s *TestStruct
	NilOrNot(s)

	aa := "1028"
	for i := 0; i < len(aa); i++ {
		fmt.Printf("当前字符：%s, 对应的ASCII：%d\n", string((aa[i])), aa[i])
	}
}

type TestStruct struct{}

func NilOrNot(v interface{}) {
	if v == nil {
		println("nil")
	} else { // 结果为：non-nil
		println("non-nil")
	}
}
