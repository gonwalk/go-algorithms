



/*
# 按行访问
## 思路

按照与逐行读取 Z 字形图案相同的顺序访问字符串。

## 算法

首先访问 行 0 中的所有字符，接着访问 行 1，然后 行 2，依此类推...

对于所有整数k，

- 行0中的字符位于索引k(k=2⋅numRows−2) 处;
- 行numRows−1中的字符位于索引k(2⋅numRows−2)+numRows−1 处;
- 内部的行 i 中的字符位于索引k(2⋅numRows−2)+i 以及(k+1)(2⋅numRows−2)−i 处;

## 复杂度分析

时间复杂度：O(n)O(n)，其中 n == \text{len}(s)n==len(s)。每个索引被访问一次。
空间复杂度：O(n)O(n)。对于 C++ 实现，如果返回字符串不被视为额外空间，则复杂度为 O(1)O(1)。

作者：LeetCode
链接：https://leetcode-cn.com/problems/zigzag-conversion/solution/z-zi-xing-bian-huan-by-leetcode/
*/






func convert(s string, numRows int) string {
    
    if numRows <= 1{
        return s
    }
    l := len(s)
    cycleLen := 2 * numRows - 2
    var newStr string
    for i := 0; i < numRows; i++{       // 行号
        for j := 0; j + i < l; j += cycleLen {   // 列号
            newStr += string(s[j+i])
            if(i != 0 && i != (numRows - 1) && (j + cycleLen - i) < l){              // 非第0行和最后一行
                newStr += string(s[j + cycleLen - i])
            }
        }
    }
    return newStr
}