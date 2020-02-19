

// 题目描述
/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解题思路1 字符串遍历比较
// 转化为字符串，前半段跟后半段字符逐个遍历比较，执行比较字符串长度的一半

func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    // 只需循环字符串长度的一半即可，即前半段跟后半段字符逐个比较
    for i:= 0; i < len(str) / 2; i++{
        if(str[i] != str[len(str)-1-i]){
            return false 
        }
    }
    return true
}

// 解题思路2 进阶解法--- 求整取余比较法
/*
通过取整和取余操作获取整数中对应的数字进行比较。

举个例子：1221 这个数字。

通过计算 1221 / 1000， 得首位1
通过计算 1221 % 10， 可得末位 1
进行比较
再将 22 取出来继续比较

链接：https://leetcode-cn.com/problems/palindrome-number/solution/dong-hua-hui-wen-shu-de-san-chong-jie-fa-fa-jie-ch/
*/

public boolean isPalindrome(int x) {
	//边界判断：负数都不是回文数字
	if (x < 0) return false;
	int div = 1;
	//
	if (x / div >= 10) {
		div *= 10
	}
	if (x > 0) {
		int left = x / div;
		int right = x % 10;
		if (left != right) return false;
		x = (x % div) / 10;
		div /= 100;
	}
	return true;
}


// 解题思路3 进阶解法---巧妙解法
/*
直观上来看待回文数的话，就感觉像是将数字进行对折后看能否一一对应。

所以这个解法的操作就是 取出后半段数字进行翻转。

这里需要注意的一个点就是由于回文数的位数可奇可偶，所以当它的长度是偶数时，它对折过来应该是相等的；当它的长度是奇数时，那么它对折过来后，有一个的长度需要去掉一位数（除以 10 并取整）。

具体做法如下：

每次进行取余操作 （ %10），取出最低的数字：y = x % 10
将最低的数字加到取出数的末尾：revertNum = revertNum * 10 + y
每取一个最低位数字，x 都要自除以 10
判断 x 是不是小于 revertNum ，当它小于的时候，说明数字已经对半或者过半了
最后，判断奇偶数情况：如果是偶数的话，revertNum 和 x 相等；如果是奇数的话，最中间的数字就在revertNum 的最低位上，将它除以 10 以后应该和 x 相等。

链接：https://leetcode-cn.com/problems/palindrome-number/solution/dong-hua-hui-wen-shu-de-san-chong-jie-fa-fa-jie-ch/

*/

public boolean isPalindrome(int x) {
	//思考：这里大家可以思考一下，为什么末尾为 0 就可以直接返回 false
	if (x < 0 || (x % 10 == 0 && x != 0)) return false;
	int revertedNumber = 0;
	if (x > revertedNumber) {
		revertedNumber = revertedNumber * 10 + x % 10;
		x /= 10;
	}
	return x == revertedNumber || x == revertedNumber / 10;
}

