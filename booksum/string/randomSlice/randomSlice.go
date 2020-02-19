


package main

import (
    "errors"
    "fmt"
    "time"
    "math/rand"
)

/*
// 1 ~ 50的整数存放在[]int切片中，如何尽可能减少循环次数的情况下随机不重复地输出全部slice中的每个元素？

把一个数组随机打乱实质就是“洗牌问题”，洗牌问题不仅追求速度，还要求洗的足够开。
应用场景：播放器的随机播放，三国杀游戏，斗地主游戏等。

Fisher-Yates随机置乱算法
也称高纳德置乱算法，该算法是无偏的，所以每个排列都是等可能的。
主要思路是：每次从已知数组随机一个数，然后将数组的最后一个值赋值到前面随机到的数的位置上，
然后将长度-1，再从原数组下标-1的数组中随机。 运算次数就是数组长度
*/

func init() {
    rand.Seed(time.Now().Unix())
}

func main() {
    strs := []string{
        "1", "2", "3", "4", "5", "6", "7", "8",
    }
    a, _ := Random(strs, 8)
    fmt.Println(a)
}

func Random(strings []string, length int) (string, error) {
    if len(strings) <= 0 {
        return "", errors.New("the length of the parameter strings should not be less than 0")
    }

    if length <= 0 || len(strings) < length {
        return "", errors.New("the size of the parameter length illegal")
    }

    for i := len(strings) - 1; i > 0; i-- {
        num := rand.Intn(i + 1)
        strings[i], strings[num] = strings[num], strings[i]
    }

    str := ""
    for i := 0; i < length; i++ {
        str += strings[i]
    }
    return str, nil
}