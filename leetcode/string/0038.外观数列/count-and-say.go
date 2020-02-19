
// 1.题目描述



// 2.算法实现
// 2.1 算法思路
/*
解题思路


![每一项分为计数和数字两部分](https://pic.leetcode-cn.com/e4d0e2acba7004bc831af05e0f62a10e63f317c8b1c7754fc5bf3fd8bbbe38f5-image.png)

根据题目，可以理解为每个N值都会分解为以1和2为元素的字符串序列，其每一项的值都可以划分为count+num



全部思路如下
1.递归（如果迭代请绕道）
2.找到最底层即n=1返回本身："1"
3.有了这个"1" 开始递归下面一层，即设定一个count=1（因为第一层已经有了一个1）
4.循环，条件①就是当：pre[i] != pre[i-1] 此时拼接pre[i]为 count + pre[i-1] (pre为上一层字符串)


链接：https://leetcode-cn.com/problems/count-and-say/solution/python-di-gui-zi-fu-chuan-mo-wei-pin-jie-zi-fu-chu/
*/

// 2.2 代码实现
func countAndSay(n int) string {
    if(n <= 1) {			// n=0或1直接返回"1"
        return "1"
    }
    
    var res,tmp string
    var count = 1				// 计算每个元素（1或2）出现的次数
    pre := countAndSay(n-1)		// 递归
    for i := range pre {
        if(i == 0) {			// pre首字符时，count置为1
            count = 1
        }else if(pre[i] == pre[i-1]) {	// 当前pre的第i个字符与前一个相同，计数器+1
            count++
        }else if(pre[i] != pre[i-1]) {	// 当前pre的第i个字符与前一个不同，如121，则计数根据计数器count和前一个字符计算中间结果
            tmp = fmt.Sprintf("%d", count) + string(pre[i-1])
            res += tmp
            count = 1
		}
		
        if(i == len(pre) - 1) {			// pre的最后一个字符
            tmp = fmt.Sprintf("%d", count) + string(pre[i])
            res += tmp
        }
    }
    return res
}