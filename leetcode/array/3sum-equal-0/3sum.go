

// 1.题目描述
/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

链接：https://leetcode-cn.com/problems/3sum
*/

// 2.解题思路
/*
排序 + 双指针
本题的难点在于如何去除重复解。

算法流程：
1.特判，对于数组长度 n，如果数组为 null 或者数组长度小于 3，返回[]。
2.对数组进行排序。
3.遍历排序后数组：
    - 若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，
    直接返回结果；
    - 对于重复元素：跳过当前循环，避免出现重复解；
    - 令左指针 L=i+1，右指针 R=n-1，当 L<R时，执行循环：
        - 当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否
        和下一位置重复，去除重复解。并同时将 L,R移到下一位置，寻找新的解；
        - 若和大于 0，说明 nums[R]太大，R左移；
        - 若和小于 0，说明 nums[L]太小，L右移。

复杂度分析
时间复杂度：O(n^2)，其中数组排序 O(NlogN)，遍历数组 O(n)，
双指针遍历O(n)，总体O(NlogN) + O(n)*O(n) = O(n^2)

空间复杂度：O(1)


链接：https://leetcode-cn.com/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/

*/

// 3.代码实现


func threeSum(nums []int) [][]int {
    var res [][]int         // 存放结果集合
    n := len(nums)          // 数组序列长度
    if(n < 3) {
        return res
    }
    sort.Ints(nums)         // 对数组进行从小到大排序
    for i:= 0; i < n; i++{
        if(nums[i] > 0) {   // 当前值太大，因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回空列表
            break
        }
        if(i > 0 && nums[i] == nums[i - 1]) {   // 当前元素与上一个重复，跳过进行下一次循环
            continue
        }
        L := i + 1
        R := n - 1          // 双指针分别指向i之后的子序列[i:n]的两端
        for L < R {
            sum := nums[i] + nums[L] + nums[R] 
            if(sum == 0) {
                var tmp []int
                res = append(res, append(tmp, nums[i], nums[L], nums[R]))  // 加入结果集后，继续下面的判断，因为有可能有重复元素
                for(L < R && nums[L] == nums[L+1]) {  // 左端元素有重复，L右移一位
                    L++
                }
                for(L < R && nums[R] == nums[R-1])  {   // 右端有重复元素，R左移一位
                    R--
                }
                L++                                                                                                                                                                                                 + // 注意这两个容易忽略
                R--
            }else if(sum < 0) { // 当前三数之和太小，L右移一位
                L++
            }else if(sum > 0) { // 当前三数之和太大，R左移一位
                R--
            }
        }
    }
    return res
}



    

// 10进制转任意进制
func decimalToAny(num, n int) string {
    new_num_str := ""
    var remainder int
    var remainder_string string
    for num != 0 {
     remainder = num % n
     if 76 > remainder && remainder > 9 {
      remainder_string = tenToAny[remainder]
     } else {
      remainder_string = strconv.Itoa(remainder)
     }
     new_num_str = remainder_string + new_num_str
     num = num / n
    }
    return new_num_str
   }
    
   // map根据value找key
   func findkey(in string) int {
    result := -1
    for k, v := range tenToAny {
     if in == v {
      result = k
     }
    }
    return result
   }
    
   // 任意进制转10进制
   func anyToDecimal(num string, n int) int {
    var new_num float64
    new_num = 0.0
    nNum := len(strings.Split(num, "")) - 1
    for _, value := range strings.Split(num, "") {
     tmp := float64(findkey(value))
     if tmp != -1 {
      new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
      nNum = nNum - 1
     } else {
      break
     }
    }
    return int(new_num)
   }