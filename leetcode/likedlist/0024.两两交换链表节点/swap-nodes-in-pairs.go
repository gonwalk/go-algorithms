
// 1.题目描述
/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.


链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
*/

// 2.算法实现

// 2.1 思路一：递归法
/*
从链表的头节点 head 开始递归。
每次递归都负责交换一对节点。由 firstNode 和 secondNode 表示要交换的两个节点。
下一次递归则是传递的是下一对需要交换的节点。若链表中还有节点，则继续递归。
交换了两个节点以后，返回 secondNode，因为它是交换后的新头。
在所有节点交换完成以后，我们返回交换后的头，实际上是原始链表的第二个节点。

时间复杂度：O(N)O(N)，其中 NN 指的是链表的节点数量。
空间复杂度：O(N)O(N)，递归过程使用的堆栈空间。
*/

func swapPairs(head *ListNode) *ListNode {
    if(head == nil ||head.Next == nil ) {
        return head
    }
    
    firstNode := head
    secondNode := head.Next

    firstNode.Next = swapPairs(secondNode.Next)  // 交换后，头结点的Next指向decond.Next递归后的节点，交换后1->2->3->4中的 1指向 3和4节点的头
    secondNode.Next = firstNode     // 交换后2指向1
    return secondNode
}

// 2.2 思路二：迭代法
/*
我们把链表分为两部分，即奇数节点为一部分，偶数节点为一部分，firstNode指的是交换节点中的前面的节点，
secondNode指的是要交换节点中的后面的节点。在完成它们的交换，还得用 prevNode 记录 A 的前驱节点。

算法：

1）firstNode（即 A） 和 secondNode（即 B） 分别遍历偶数节点和奇数节点，即两步看作一步。
2）交换两个节点：
 firstNode.next = secondNode.next
 secondNode.next = firstNode
3）更新 prevNode.next 指向交换后的头。
prevNode.next = secondNode
4）迭代完成后得到最终的交换结果。

时间复杂度：O(N)，其中 N 指的是链表的节点数量。
空间复杂度：O(1)。
*/


func swapPairs(head *ListNode) *ListNode {
    if(head == nil || head.Next == nil) {
        return head
    }
    var dummy = &ListNode{-1, nil}
    dummy.Next = head
    pre := dummy
    for head != nil && head.Next != nil {
        firstNode := head
        secondNode := head.Next
        
        pre.Next = secondNode
        firstNode.Next = secondNode.Next
        secondNode.Next = firstNode

        pre = firstNode
        head = firstNode.Next
    }
    return dummy.Next
}