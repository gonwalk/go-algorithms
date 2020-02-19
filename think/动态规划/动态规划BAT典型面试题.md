




# 1.爬楼梯

链接：https://leetcode-cn.com/problems/climbing-stairs

## 1.1 题目描述

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶


方法一：暴力法（不可取，超时，效率太低）


在暴力法中，我们将会把所有可能爬的阶数进行组合，也就是 1 和 2 。而在每一步中我们都会继续调用 climbStairsclimbStairs 这个函数模拟爬 11 阶和 22 阶的情形，并返回两个函数的返回值之和。

climbStairs(i,n)=(i + 1, n) + climbStairs(i + 2, n)
climbStairs(i,n)=(i+1,n)+climbStairs(i+2,n)

其中 ii 定义了当前阶数，而 nn 定义了目标阶数。

func climbStairs(n int) int {
    if(n == 1 || n == 2){   // 一个台阶只有一种情况：1，两个台阶有两种：1+1,2
        return n
    }
    return climbStairs(n-1) + climbStairs(n-2)  // 递归为第一步先爬1个台阶然后剩下的n-1步递归，或者一次性先爬2个台阶剩下的再去递归，这两种情况加起来
}

复杂度分析

时间复杂度：O(2^n)，树形递归的大小为 2^n。

空间复杂度：O(n)，递归树的深度可以达到 n 。

方法二：记忆化递归
算法

在上一种方法中，我们计算每一步的结果时出现了冗余。另一种思路是，我们可以把每一步的结果存储在 memomemo 数组之中，每当函数再次被调用，我们就直接从 memomemo 数组返回结果。

在 memomemo 数组的帮助下，我们得到了一个修复的递归树，其大小减少到 nn。

Java
public class Solution {
    public int climbStairs(int n) {
        int memo[] = new int[n + 1];
        return climb_Stairs(0, n, memo);
    }
    public int climb_Stairs(int i, int n, int memo[]) {
        if (i > n) {
            return 0;
        }
        if (i == n) {
            return 1;
        }
        if (memo[i] > 0) {
            return memo[i];
        }
        memo[i] = climb_Stairs(i + 1, n, memo) + climb_Stairs(i + 2, n, memo);
        return memo[i];
    }
}
复杂度分析

时间复杂度：O(n)O(n)，树形递归的大小可以达到 nn。
空间复杂度：O(n)O(n)，递归树的深度可以达到 nn。


方法三：动态规划
算法

不难发现，这个问题可以被分解为一些包含最优子结构的子问题，即它的最优解可以从其子问题的最优解来有效地构建，我们可以使用动态规划来解决这一问题。

第 i阶可以由以下两种方法得到：

在第 (i-1)阶后向上爬一阶。

在第 (i-2) 阶后向上爬 2 阶。

所以到达第 i阶的方法总数就是到第 (i-1)阶和第 (i-2) 阶的方法数之和。

令 dp[i] 表示能到达第 i 阶的方法总数：

dp[i]=dp[i-1]+dp[i-2]

示例：



func climbStairs(n int) int {
    dp := make([]int, n+1)        // 动态规划数组，存放结果集
    if(n <= 2){                   // 注意，如果n=1不从这里返回，程序走到dp[2]的时候就出现溢出异常
        return n
    }
    dp[1] = 1
    dp[2] = 2
    for i:= 3; i <= n; i++ {    // 假设要到达第N层，只有两种走法：从N-1层上来再走一步；从N-2层上来然后走2步。这两种走法之和，即为到达第N层的走法总数
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

复杂度分析：

时间复杂度：O(n)，单循环到 n 。

空间复杂度：O(n)，dp数组用了 n的空间。

方法四：斐波那契数
算法

在上述方法中，我们使用 dpdp 数组，其中 dp[i]=dp[i-1]+dp[i-2]。可以很容易通过分析得出 dp[i] 其实就是第 i 个斐波那契数。

Fib(n)=Fib(n-1)+Fib(n-2)

现在我们必须找出以 11 和 22 作为第一项和第二项的斐波那契数列中的第 nn 个数，也就是说 Fib(1)=1Fib(1)=1 且 Fib(2)=2Fib(2)=2。

Java
public class Solution {
    public int climbStairs(int n) {
        if (n == 1) {
            return 1;
        }
        int first = 1;
        int second = 2;
        for (int i = 3; i <= n; i++) {
            int third = first + second;
            first = second;
            second = third;
        }
        return second;
    }
}
复杂度分析

时间复杂度：O(n)O(n)，单循环到 nn，需要计算第 nn 个斐波那契数。

空间复杂度：O(1)O(1)，使用常量级空间。


链接：https://leetcode-cn.com/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode/

