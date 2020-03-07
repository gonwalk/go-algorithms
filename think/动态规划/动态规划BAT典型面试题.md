
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


## 1.2 算法实现

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

# 2.不同路径II

算法来源：

来源：力扣（LeetCode 63）
链接：https://leetcode-cn.com/problems/unique-paths-ii

## 2.1 题目描述

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

## 2.2 算法实现

### 2.2.1 算法思路

使用动态规划的思想，步骤如下：
```
1.声明动态规划数组dp[][]，用于计算通往当前节点的路径数。
2.如果第一个格点 obstacleGrid[0,0] 是 1，说明有障碍物，那么机器人不能做任何移动，我们返回结果 0。否则，如果 obstacleGrid[0,0] 是 0，我们初始化这个值为 1 然后继续算法。
3.遍历第一行，如果有一个格点初始值为 1 ，说明当前节点有障碍物，没有路径可以通过，设值为 0 ；否则设这个值是前一个节点的值 obstacleGrid[i,j] = obstacleGrid[i,j-1]。
4.遍历第一列，如果有一个格点初始值为 1 ，说明当前节点有障碍物，没有路径可以通过，设值为 0 ；否则设这个值是前一个节点的值 obstacleGrid[i,j] = obstacleGrid[i-1,j]。
5.现在，从 obstacleGrid[1,1] 开始遍历整个数组，如果某个格点初始不包含任何障碍物，就把值赋为上方和左侧两个格点方案数之和 obstacleGrid[i,j] = obstacleGrid[i-1,j] + obstacleGrid[i,j-1]。
6.如果这个点有障碍物，设值为 0 ，这可以保证不会对后面的路径产生贡献。
```

### 2.2.2 代码实现
```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    if(m <= 0 || n <= 0) {
        return 0
    }
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    if(obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1) {
        return 0                        // 起点或重点有障碍，不能进行任何走动，返回0种路径
    } 
    dp[0][0] = 1
    
    for i := 1; i < m; i++ {
        if(obstacleGrid[i][0] != 1) {   // 第0列中，当前位置没有障碍，就赋值为上个位置的路径数
            dp[i][0] = dp[i-1][0]
        } else {
            dp[i][0] = 0                // 第0列中，有任何一个格子有障碍，就不能往下走，当前路径数赋为0
        }
    }

    for i := 1; i < n; i++ {
        if(obstacleGrid[0][i] != 1) {   // 第0行中，当前位置没有障碍，就赋值为上个位置的路径数
            dp[0][i] = dp[0][i-1]
        } else {
            dp[0][i] = 0
        }
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if(obstacleGrid[i][j] != 1) {   // 当前位置没有障碍，就赋值为左侧和上侧位置的路径和
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            } else {                        // 当前位置有障碍，则当前位置路径数置为0
                dp[i][j] = 0
            }
        }
    }
    return dp[m-1][n-1]
}
```

