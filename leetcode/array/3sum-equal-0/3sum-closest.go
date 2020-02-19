
/* 1.题目描述
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

链接：https://leetcode-cn.com/problems/3sum-closest
*/

// 2.代码实现
func threeSumClosest(nums []int, target int) int {
    n := len(nums)
    if(n < 3){
        return 0
    }
    sort.Ints(nums)
    var sum int
    res := nums[0] + nums[1] + nums[2]
    
    for i:= 0; i < n-2; i++ {
        start, last := i+1, n - 1
        for start < last {
            sum = nums[i] + nums[start] + nums[last]
            if(absInt(sum, target) < absInt(res, target)){
                res = sum
            }
            if(sum > target){
                last--
            }else if(sum < target){
                start++
            }else {
                return sum
            }
        }
    }
    return res
    
}

func absInt(i, j int) int {
    if(i > j){
        return i - j
    }else {
        return j - i
    }
}

