# 1.题目描述

来源：力扣（LeetCode 25）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
难度：hard

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



# 2.算法实现

## 2.1 算法思路

1、先反转以 head 开头的 k 个元素；
2、将第 k + 1 个元素作为 head 递归调用 reverseKGroup 函数；
3、将上述两个过程的结果连接起来。


## 2.2 代码实现


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if(head == nil || k == 1) {
        return head     // 不需要翻转
    }

    a, b := head, head      // a,b开始都指向头结点
    for i:=0; i<k; i++ {    // 区间[a,b)包含k个待翻转元素
        if(b == nil) {      // 不足k个元素，则不需要翻转，直接返回这部分的链表头
            return head
        }
        b = b.Next
    }

    newHead := reverse(a, b)        // 翻转前k个元素
    a.Next = reverseKGroup(b, k)    // 递归翻转后续链表并连接起来
    return newHead
}

// 翻转区间[a, b)的元素
func reverse(a, b *ListNode) *ListNode {
    var pre *ListNode
    cur, next := a, a
    for(cur != b) { // 未到达b节点
        next = cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre      // 返回翻转后的头结点
}
```

执行用时 :4 ms, 在所有 Go 提交中击败了93.13%的用户
内存消耗 :3.7 MB, 在所有 Go 提交中击败了47.39%的用户