package main

// 3. 无重复字符的最长子串
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

/* 题目描述
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

/* 解题思路1
此题为经典的窗口滑动中的非定长算法。
解题思路为抽象符合条件的字符串在窗口内，依次遍历字符串，当前位置字符串如果包含在窗口中，窗口中的开始位置更新为重复位置的下一位 最后得到最长非重复字符串的值 代码如下：

作者：colas
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/gosliding-window-algorithm-zhi-fei-ding-chang-by-c/
来源：力扣（LeetCode）
*/
func lengthOfLongestSubstring_1(s string) int {
    // 定义游标尺寸大小,游标的左边位置
    window,start := 0,0

    // 循环字符串
    for key := 0; key < len(s); key++ {
        // 查看当前字符串string(s[key])在滑动窗口所在游标（string(s[start:key])）内的索引，不存在isExist = -1
        isExist := strings.Index(string(s[start:key]), string(s[key]));
        
        // 如果不存在游标内部,滑动窗口长度重新计算并赋值当前字符索引 - 窗口其实字符索引 + 1
        if (isExist == -1) {
            if (key - start + 1 > window) {
                window = key - start + 1
            }
        } else { //存在，游标开始位置更换为重复字符串位置的下一个位置
            start = start + 1 + isExist
        }
    }
    return window
}




/* 解题思路2
利用location保存字符上次出现的序列号，可以避免了查询工作。location和Two Sum中的m是一样的作用。
利用s[left:i+1]来表示s[:i+1]中的包含s[i]的最长子字符串。 location[s[i]]是字符s[i]在s[:i+1]中倒数第二次出现的序列号。 当left < location[s[i]]的时候，说明字符s[i]出现了两次。需要设置 left = location[s[i]] + 1, 保证字符s[i]只出现一次。

总结

// m 负责保存map[整数]整数的序列号
	m := make(map[int]int, len(nums))
*/
func lengthOfLongestSubstring_2(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 说明s[i]已经在s[left:i+1]中重复了
		// 并且s[i]上次出现的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}


