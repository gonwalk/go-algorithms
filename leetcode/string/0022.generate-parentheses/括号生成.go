

// 1.问题描述
/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

链接：https://leetcode-cn.com/problems/generate-parentheses


*/

反思：
首先，面向小白：什么是动态规划？在此题中，动态规划的思想类似于数学归纳法，当知道所有 i<n 的情况时，我们可以通过某种算法算出 i=n 的情况。
本题最核心的思想是，考虑 i=n 时相比 n-1 组括号增加的那一组括号的位置。

思路：
当我们清楚所有 i<n 时括号的可能生成排列后，对与 i=n 的情况，我们考虑整个括号排列中最左边的括号。
它一定是一个左括号，那么它可以和它对应的右括号组成一组完整的括号 "( )"，我们认为这一组是相比 n-1 增加进来的括号。

那么，剩下 n-1 组括号有可能在哪呢？

【这里是重点，请着重理解】

剩下的括号要么在这一组新增的括号内部，要么在这一组新增括号的外部（右侧）。

既然知道了 i<n 的情况，那我们就可以对所有情况进行遍历：

"(" + 【i=p时所有括号的排列组合】 + ")" + 【i=q时所有括号的排列组合】

其中 p + q = n-1，且 p q 均为非负整数。

事实上，当上述 p 从 0 取到 n-1，q 从 n-1 取到 0 后，所有情况就遍历完了。

注：上述遍历是没有重复情况出现的，即当 (p1,q1)≠(p2,q2) 时，按上述方式取的括号组合一定不同。

代码：
具体代码如下：（时间击败百分之 95，内存击败百分之 99.65）

Python
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        if n == 0:
            return []
        total_l = []
        total_l.append([None])    # 0组括号时记为None
        total_l.append(["()"])    # 1组括号只有一种情况
        for i in range(2,n+1):    # 开始计算i组括号时的括号组合
            l = []        
            for j in range(i):    # 开始遍历 p q ，其中p+q=i-1 , j 作为索引
                now_list1 = total_l[j]    # p = j 时的括号组合情况
                now_list2 = total_l[i-1-j]    # q = (i-1) - j 时的括号组合情况
                for k1 in now_list1:  
                    for k2 in now_list2:
                        if k1 == None:
                            k1 = ""
                        if k2 == None:
                            k2 = ""
                        el = "(" + k1 + ")" + k2
                        l.append(el)    # 把所有可能的情况添加到 l 中
            total_l.append(l)    # l这个list就是i组括号的所有情况，添加到total_l中，继续求解i=i+1的情况
        return total_l[n]


链接：https://leetcode-cn.com/problems/generate-parentheses/solution/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/



方法二：回溯法
思路和算法

只有在我们知道序列仍然保持有效时才添加 '(' or ')'，而不是像 方法一 那样每次添加。我们可以通过跟踪到目前为止放置的左括号和右括号的数目来做到这一点，

如果我们还剩一个位置，我们可以开始放一个左括号。 如果它不超过左括号的数量，我们可以放一个右括号。

JavaPython
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> ans = new ArrayList();
        backtrack(ans, "", 0, 0, n);
        return ans;
    }

    public void backtrack(List<String> ans, String cur, int open, int close, int max){
        if (cur.length() == max * 2) {
            ans.add(cur);
            return;
        }

        if (open < max)
            backtrack(ans, cur+"(", open+1, close, max);
        if (close < open)
            backtrack(ans, cur+")", open, close+1, max);
    }
}

复杂度分析

我们的复杂度分析依赖于理解 generateParenthesis(n) 中有多少个元素。这个分析超出了本文的范畴，
但事实证明这是第 n 个卡塔兰数(2n n)/(n+1)}，这是由 4^n/(n*sqrt(n)) 
渐近界定的。

时间复杂度：4^n/(sqrt(n)) ，在回溯过程中，每个有效序列最多需要 n 步。

空间复杂度：4^n/(sqrt(n)) ，如上所述，并使用 O(n)O(n) 的空间来存储序列。

方法三：闭合数
思路

为了枚举某些内容，我们通常希望将其表示为更容易计算的不相交子集的总和。

考虑有效括号序列 S 的 闭包数：至少存在 index >= 0，使得 S[0], S[1], ..., S[2*index+1]是有效的。 显然，每个括号序列都有一个唯一的闭包号。 我们可以尝试单独列举它们。

算法

对于每个闭合数 c，我们知道起始和结束括号必定位于索引 0 和 2*c + 1。然后两者间的 2*c 个元素一定是有效序列，其余元素一定是有效序列。

Java
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> ans = new ArrayList();
        if (n == 0) {
            ans.add("");
        } else {
            for (int c = 0; c < n; ++c)
                for (String left: generateParenthesis(c))
                    for (String right: generateParenthesis(n-1-c))
                        ans.add("(" + left + ")" + right);
        }
        return ans;
    }
}
复杂度分析

时间和空间复杂度：O(4^n / sqrt(n))

	
闭合法的理解： 对于任何一个括号组合（1对以上），可以表达为这么一种组合 ( left ) right 的形式。其中，若left 和 right 两部分的括号总数为 n-1对，那么新的组合 ( left ) right 则有n对，这样将n对括号求解，转化为n-1对的求解，以此类推直到零对括号可以直接给出空字符串的解。以3对为例，可拆解为 “( left=0对）right=2对”“( left=1对）right=1对”“( left=2对）right=0对”。对于上述题解给出的方案，如果缓存中间组合，可以减少搜索次数。

链接：https://leetcode-cn.com/problems/generate-parentheses/solution/gua-hao-sheng-cheng-by-leetcode/

闭合法理解:

算法: 递归关系 n对有效括号序列由c 对有效括号序列和 n-c对有效括号序列组成
n对有效括号序列 是由 "(" + $1 + ")" + $2 拼接而成, $1 表示c(0<=c<n)对有效括号序列集合中的某一个元素, $2表示剩下的n-1-c对有效括号序列构成的集合中的某一个元素. 这样说大致就能看懂程序在做什么了
这里要说一下为什么要ans.add("(" + left + ")" + right); 而不是ans.add(left+right)呢, 应该是因为有一种特殊的情况 如果left right(n=4) 全是()() 这种分隔的小括号 则当(left,right)为(1,3)或者(2,2)或者(3,1) 就会出现()()()()重复出现多次的情况, 而上面那种做法可以将()()()()的情况放到left=0,rigint=3的情况中,且只会出现一次, 非常nice


func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gene(&res,"",n,n)
	return res
}

// 递归生成括号
// left和right表示还剩余多少个左右半括号未生成，str是初始时字符串（为空）
func gene(res *[]string, str string, left, right int) {
	if left == 0 { // 此时表明left已经出栈完毕，需要把right中的全部出栈
		for i := 0; i < right; i++ {
			str += ")"
		}
		*res = append(*res, str)
		return
	}
	gene(res, str+"(", left-1,right)
	if left < right {
		gene(res, str+")", left,right-1)
	}
}