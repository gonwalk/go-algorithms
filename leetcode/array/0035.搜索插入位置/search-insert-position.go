
// 1.题目描述
// 题目来源：链接：https://leetcode-cn.com/problems/search-insert-position
/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

*/


// 2.算法实现

// 2.1 方法一：暴力破解法，时间复杂度O(n)
func searchInsert1(nums []int, target int) int {
    n := len(nums)
    if(n == 0) {
        return 0
    }
    var i int
    for i < n {
        if(nums[i] == target) {             // 数组中存在目标数，则直接返回
            return i
        } else if(nums[i] < target) {       // 当前位置元素小于目标值，则索引下移
            i++
        }else{                              // 数组中当前位置元素大于目标值，则返回当前索引，因为数组是逐个索引遍历的
            return i
        }
    }
    return n
}


// 2.2 方法二：二分查找法：时间复杂度O(logn)
/*
如果该题目暴力解决的话需要 O(n) 的时间复杂度，但是如果二分的话则可以降低到 O(logn)的时间复杂度。
整体思路和普通的二分查找几乎没有区别，先设定左侧下标 left 和右侧下标 right，再计算中间下标 mid。
每次根据 nums[mid] 和 target 之间的大小进行判断，相等则直接返回下标，nums[mid] < target 则 left 右移，nums[mid] > target 则 right 左移
查找结束如果没有相等值则返回 left，该值为插入位置。
时间复杂度：O(logn)
*/
func searchInsert(nums []int, target int) int {
    left, mid, right := 0, 0, len(nums) - 1
    for left <= right {
        mid = (left+right)/2
        if(nums[mid] == target) {
            return mid
        }else if(nums[mid] < target) {
            left++
        }else {
            right--
        }
    }
    return left
}