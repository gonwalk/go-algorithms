

// 题目描述
/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

链接：https://leetcode-cn.com/problems/combination-sum
*/

// 解题思路1 递归法
/*
思路：将target的值依次减去数组元素并递归调用直到target的值小于或等于数组各元素，由于数组元素可重复使用，但重复组合不计在内，所以递归调用时只包含当前元素及之后的所有元素即可。

链接：https://leetcode-cn.com/problems/combination-sum/solution/go-di-gui-jie-fa-by-mrandmrsbenben/
*/
func combinationSum(candidates []int, target int) [][]int {
	comb := [][]int{}
	for i := range candidates {
		if candidates[i] == target {
			comb = append(comb, []int{candidates[i]})
		} else if candidates[i] < target {
			rtn := combinationSum(candidates[i:], target-candidates[i])
			for j := range rtn {
				if len(rtn[j]) == 0 {
					continue
				}
				comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
			}
		}
	}
	return comb
}



// 链接：https://leetcode-cn.com/problems/combination-sum/solution/di-gui-fa-jie-jue-zu-he-zong-he-de-wen-ti-by-chung/
func GetCombine(candidates, thisRes []int, lastChoise, left int, outResult *[][]int)  {
	if left == 0 {
		*outResult = append(*outResult, thisRes)
		return
	}

	for _, value := range candidates {
		if value < lastChoise { // 选取的数递增，用于去重
			continue
		}
		if left < value { // 当前数小于剩余数时，不用再检查
			continue
		}

		newRes := make([]int, len(thisRes)+1)
		copy(newRes, thisRes[:len(thisRes):len(thisRes)])
		newRes[len(thisRes)] = value

		GetCombine(candidates, newRes, value, left - value, outResult)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	outResult := make([][]int, 0)
	GetCombine(candidates, nil, 0, target, &outResult)
	return outResult
}

// 解题思路2 回溯法
/*
首先将数组按升序排序, 因为可以重复包含相同的元素, 所以在递归的时候指向当前的索引值不需要+1, 只有当当前索引的值累加后小于target时, 索引才向后移动判断其他元素。

例如：

target = 7, candidates = [2, 3, 5]
从candidates的第一个元素开始。target -= 2, 同时将元素2放入临时数组nums中。
当前nums = [2], target = 5, 继续调用函数
得到nums = [2, 2], target = 3, 继续
nums = [2,2,2], target = 1, 继续
nums = [2,2,2,2], target = -1。此时得到数字之和不满足条件, 回溯到在这之前的状态, 判断下一个元素3的值。
nums = [2,2,3], target = 0。 满足数组和等于target的条件, 写入返回值res中。因为数组是升序排序的, 
后续元素的操作必定大于target, 所以可以跳过。

回溯到nums = [2]的状态中。开始下一个元素的判断。
nums = [2, 3], target = 2.
nums = [2, 3, 3], target = -1. 跳过。
nums = [2, 5], target = 0. 正确。

nums = [3], target = 4.
nums = [3, 3], target = 1.
nums = [3,3,3], target = -2.跳过

nums = [5], target = 2
nums = [5, 5], target = -3. 跳过。

遍历结束。

重复以上操作, 穷举所有可能, 最后返回结果。

链接：https://leetcode-cn.com/problems/combination-sum/solution/golanghui-su-by-zhouql/
*/
func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(candidates)
    dofunc(candidates, []int{}, &res, 0, target)
    return res
}

func dofunc(candidates, nums []int, res *[][]int, ind, target int) bool {
    if target == 0 {
        t := make([]int, len(nums))
        copy(t, nums)
        *res = append(*res, t)
        return true
    } else if target < 0 {
        return false
    }
    for i := ind; i < len(candidates); i++ {
        nums = append(nums, candidates[i])
        if !dofunc(candidates, nums, res, i, target-candidates[i]) {
            break
        }
        nums = nums[:len(nums)-1]
    }
    return true
}






// 解题思路3 动态规划
/*
这个问题属于完全背包问题，由于需要找到所有解，所以在计算子问题的时候需要额外记录所有可行的结果。代码如下，使用dp来记录子问题是否可解，用tab来记录所有的解。

链接：https://leetcode-cn.com/problems/combination-sum/solution/golang-dpsuan-fa-chao-guo-9275-by-jhzhu89/
*/
func combinationSum(candidates []int, target int) [][]int {
	dp := make([]bool, target+1)
	dp[0] = true
	tab := make([][][]int, target+1, target+1)
	tab[0] = [][]int{[]int{}}

	for _, c := range candidates {
		for k := c; k <= target; k++ {
			if dp[k-c] {
				dp[k] = true
				for _, l := range tab[k-c] {
					tab[k] = append(tab[k], append([]int{c}, l...))
				}
			}
		}
	}

	if dp[target] {
		return tab[target]
	}
	return nil
}
