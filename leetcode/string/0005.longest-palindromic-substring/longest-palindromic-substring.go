
// 5. 最长回文子串

/* 题目描述

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
在真实的面试中遇到过这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
*/

最长回文子串：1.寻找回文子串；2.该子串是回文子串中长度最长的。

其他详细的分析和解法，可以参考：
最长回文子串（Longest Palindromic Substring）——三种时间复杂度的解法https://blog.csdn.net/qq_32354501/article/details/80084325

// 解题思路1 暴力破解法
```
1）从最长的子串开始，遍历所有该原字符串的子串；

2）每找出一个字符串，就判断该字符串是否为回文；

3）子串为回文时，则找到了最长的回文子串，因此结束；反之，则继续遍历。
```


时间复杂度解释：
    遍历字符串子串：嵌套一个循环、O(n^2)；   
    判断是否为回文：再次嵌套一个循环、O(n^3);
	该方法效率很低，面试的时候使用该方式面试官会很不满意。


```go
func longestPalindrome(s string) string {
	// num和maxStr分别存放当前循环中回文的最大长度和最长回文字符串
	num := 0
	maxStr := ""
    if len(s) < 2{ // s是空字符串或者只有一个字符，直接返回s本身
        return s
    }
    if isPalindrome(s){ // s串本身就是回文，那么不需要判断，其肯定比子串都长，是最长的回文串
        return s
    }
    for i := 0; i < len(s); i++{
        for j := i+1; j < len(s); j++{
            if isPalindrome(string(s[i:j+1])){
				// 注意：切片s[i:j+1]表示的元素下标为i到j，不包括j+1
                if(len(maxStr) < len(s[i:j+1])){
                    num = len(s[i:j+1])
                    maxStr = s[i:j+1]
                }
            }
        }
	}
	// 字符串只有一个字符，那这个字符就是最长回文
    if num < 2 {
        return string(s[0:1])
    }
    return maxStr
}

func isPalindrome(s string)bool{
    l := len(s)
    if l < 2 {
        return true 
    }
	for i:=0 ; i < l / 2; i++{  // 只需要判断一般的字符串就好，
		// 字符串长度为奇数，如3/2=1(比较索引0和2，最中间的索引1不需要比较，比较(len(s)-1)/2次)， 
		// 字符串长度为偶数，如4/2=2(比较索引0和3，1和2，都需要比较，比较len(s)/2次)
        if(s[i] != s[l-1-i]){
            return false
        }
    }
    return true
}
```

/*
第一种思路:

以每个字母为回文中心,考虑回文长度为奇数和偶数的情况.

第二种思路:

以每个字母为回文串的结束标志,分别考虑可能回文为奇数和偶数的情况.

第三种思路:

令dp[j][i]从字符串j到i是否为回文串

动态回归方程 dp[j][i]是看j+1和i-1是否为回文串.
*/

/* 解题思路2 中心扩展法
1.解题思路：
 	1）将子串分为单核和双核的情况，单核即指子串长度为奇数，双核则为偶数；

    2）遍历每个除最后一个位置的字符index(字符位置)，单核：初始low = 初始high = index，low和high均不超过原字符串的下限和上限；判断low和high处的字符是否相等，相等则low++、high++（双核：初始high = 初始low+1 = index + 1）；

    3）每次low与high处的字符相等时，都将当前最长的回文子串长度与high-low+1比较。后者大时，将最长的回文子串改为low与high之间的；

    4）重复执行2）、3），直至high-low+1 等于原字符串长度或者遍历到最后一个字符，取当前截取到的回文子串，该子串即为最长的回文子串。



2.时间复杂度解释：

   遍历字符：一层循环、O(n-1)；

   找以当前字符为中心的最长回文子串：嵌套两个独立循环、O(2n*(n-1)) = O(n^2)。

原文链接：https://blog.csdn.net/qq_32354501/article/details/80084325

题目要求寻找字符串中的最长回文。 当然，我们可以使用下面的程序判断字符串s[i:j+1]是否是回文。
func isPalindromic(s *string, i, j int ) bool {
    for  i< j {
        if (*s)[i] != (*s)[j] {
            return false
        } 
        i++
        j--
    }
    return true
}
但是，这样就没有利用回文的一下特点，假定回文的长度为l，x为任意字符

当l为奇数时，回文的正中间段只会是，“x”，或“xxx”，或“xxxxx”，或...
当l为偶数时，回文的正中间段只会是，“xx”，或“xxxx”，或“xxxxxx”，或...
同一个字符串的任意两个回文substring的正中间段，不会重叠。
*/

基于中心扩展法的实现方式一：
 //中心扩展算法，可以在n*n之间完成
 func longestPalindrome2(s string) string {   //babad bab
	if len(s) <= 1 {  // 空字符串、单个字符串可以认为是回文，直接返回
		return s
	}
	start,end := 0,0  // start和end分别表示当前回文串在s串中开始和结束处的索引
	for i:=0;i<len(s);i++{
		len1 :=expandAroundCenter(s,i,i) //a [b] a bd 当字符串s长度为奇数时，回文串的中间索引只有一个，需要从这个数开始向两边扩散
		len2 := expandAroundCenter(s,i,i+1)//a [bb] c 当字符串s长度为奇数时，回文串的中间位置有两个，需要从这两个中间的索引开始向两边扩散
		len := int(math.Max(float64(len1),float64(len2))) //比较出两个的最长的长度
		if len > end - start{
			//len=3  i=1
			start =i-(len-1)/2    // i:在这里就是一个回文数的中间数，（len-1）:是处理len可能是偶数的情况，(len-1)/2表示回文的半径，start既是回文开始处的索引
			end = i+len/          // end回文结束处的索引
		}
	}
	return string(s[start:end+1])
}
////babad bab
func expandAroundCenter(s string,left,right int)int{  // left和right表示字符串s最中间的两个字符对应的索引下标，s长度为偶数，索引right=left+1,；s长度为奇数，则left=right
	L,R := left,right
	for L>=0 && R <len(s) && s[L] == s[R]{
		L--      // 左侧索引向前扩展
		R++      // 右侧索引向后扩展
	}
	return R-L+1 // 回文串结束，计算出该回文串的长度
}

基于中心扩展法的实现方式二：
func longestPalindrome(s string) string {
	if len(s) < 2 { // 肯定是回文，直接返回
		return s
	}

	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1

	// 在 for 循环中
	// b 代表回文的**首**字符索引号，
	// e 代表回文的**尾**字符索引号，
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < len(s); { // 以s[i]为`正中间段`首字符开始寻找最长回文。
		if len(s)-i <= maxLen/2 {
			// 因为i是回文`正中间段`首字符的索引号
			// 假设此时能找到的最长回文的长度为l, 则，l <= (len(s)-i)*2 - 1
			// 如果，len(s)-i <= maxLen/2 成立
			// 则，l <= maxLen - 1
			// 则，l < maxLen
			// 所以，无需再找下去了。
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// 循环结束后，s[b:e+1]是一串相同的字符串
		}

		// 下一个回文的`正中间段`的首字符只会是s[e+1]
		// 为下一次循环做准备
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 循环结束后，s[b:e+1]是这次能找到的最长回文。
		}

		newLen := e + 1 - b
		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}

### 思路3 Manacher算法
最长回文子串（Longest Palindromic Substring）——三种时间复杂度的解法：https://blog.csdn.net/qq_32354501/article/details/80084325

1.解题思路
	1）将原字符串S的每个字符间都插入一个永远不会在S中出现的字符（本例中用“#”表示），在S的首尾也插入该字符，使得到的新字符串S_new长度为2*S.length()+1，保证Len的长度为奇数(下例中空格不表示字符，仅美观作用)；

        例：S：       a  a  b  a  b  b  a

        S_new:        #  a  #  a  #  b  #  a  #  b  #  b  #  a  #

    2）根据S_new求出以每个字符为中心的最长回文子串的最右端字符距离该字符的距离，存入Len数组中，即S_new[i]—S_new[r]为S_new[i]的最长回文子串的右段（S_new[2i-r]—S_new[r]为以S_new[i]为中心的最长回文子串），Len[i] = r - i + 1;

        S_new:        #  a  #  a  #  b  #  a  #  b  #  b  #  a  #

        Len:            1  2  3  2  1   4  1  4  1  2  5   2  1  2  1

        Len数组性质：Len[i] - 1即为以Len[i]为中心的最长回文子串在S中的长度。在S_new中，以S_new[i]为中心的最长回文子串长度为2Len[i] - 1，由于在S_new中是在每个字符两侧都有新字符“#”，观察可知“#”的数量一定是比原字符多1的，即有Len[i]个，因此真实的回文子串长度为Len[i] - 1，最长回文子串长度为Math.max(Len) - 1。

    3）Len数组求解（线性复杂度（O(n)））：

       a.遍历S_new数组，i为当前遍历到的位置，即求解以S_new[i]为中心的最长回文子串的Len[i];

       b.设置两个参数：sub_midd = Len.indexOf(Math.max(Len)表示在i之前所得到的Len数组中的最大值所在位置、sub_side = sub_midd + Len[sub_midd] - 1表示以sub_midd为中心的最长回文子串的最右端在S_new中的位置。起始sub_midd和sub_side设为0，从S_new中的第一个字母开始计算，每次计算后都需要更新sub_midd和sub_side；

       c.当i < sub_side时，取i关于sub_midd的对称点j（j = 2sub_midd - i，由于i <= sub_side，因此2sub_midd - sub_side <= j <= sub_midd);当Len[j] < sub_side - i时，即以S_new[j]为中心的最长回文子串是在以S_new[sub_midd]为中心的最长回文子串的内部，再由于i、j关于sub_midd对称，可知Len[i] = Len[j];
		当Len[j] >= sub.side - i时说明以S_new[i]为中心的回文串可能延伸到sub_side之外，而大于sub_side的部分还没有进行匹配，所以要从sub_side+1位置开始进行匹配，直到匹配失败以后，从而更新sub_side和对应的sub_midd以及Len[i]；

	   d.当i > sub_side时，则说明以S_new[i]为中心的最长回文子串还没开始匹配寻找，因此需要一个一个进行匹配寻找，结束后更新sub_side和对应的sub_midd以及Len[i]。

2.时间复杂度解释：

算法只有遇到还没匹配的位置时才进行匹配，已经匹配过的位置不再进行匹配，因此大大的减少了重复匹配的步骤，对于S_new中的每个字符只进行一次匹配。所以该算法的时间复杂度为O(2n+1)—>O(n)（n为原字符串的长度），所以其时间复杂度依旧是线性的。


原文链接：https://blog.csdn.net/qq_32354501/article/details/80084325

public String longestPalindrome(String s) {
        List<Character> s_new = new ArrayList<>();
        for(int i = 0;i < s.length();i++){
            s_new.add('#');
            s_new.add(s.charAt(i));
        }
        s_new.add('#');
        List<Integer> Len = new ArrayList<>();
        String sub = "";//最长回文子串
        int sub_midd = 0;//表示在i之前所得到的Len数组中的最大值所在位置
        int sub_side = 0;//表示以sub_midd为中心的最长回文子串的最右端在S_new中的位置
        Len.add(1);
        for(int i = 1;i < s_new.size();i++){
            if(i < sub_side) {//i < sub_side时，在Len[j]和sub_side - i中取最小值，省去了j的判断
                int j = 2 * sub_midd - i;
                if(j >= 2 * sub_midd - sub_side &&  Len.get(j) <= sub_side - i){
                    Len.add(Len.get(j));
                }
                else
                    Len.add(sub_side - i + 1);
            }
            else//i >= sub_side时，从头开始匹配
                Len.add(1);
            while( (i - Len.get(i) >= 0 && i + Len.get(i) < s_new.size()) && (s_new.get(i - Len.get(i)) == s_new.get(i + Len.get(i))))
                Len.set(i,Len.get(i) + 1);//s_new[i]两端开始扩展匹配，直到匹配失败时停止
            if(Len.get(i) >= Len.get(sub_midd)){//匹配的新回文子串长度大于原有的长度
                sub_side = Len.get(i) + i - 1;
                sub_midd = i;
            }
        }
        sub = s.substring((2*sub_midd - sub_side)/2,sub_side /2);//在s中找到最长回文子串的位置
        return sub;
    }