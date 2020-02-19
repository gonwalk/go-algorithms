package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{5, nil}
	e := &ListNode{6, nil}
	a.Next = b
	b.Next = c
	c.Next = d	
	d.Next = e
	e.Next = c 
	cycleNode := detectCycle(a)
	cycleLen := cycleLen(a)
	fmt.Printf("链表中环的起始节点：%v, 环长：%d\n", cycleNode, cycleLen)
}

// 检测链表中环的起始节点，无环则返回nil
func detectCycle(head *ListNode) *ListNode {
    if(head == nil || head.Next == nil) {
        return nil
    }
    slow := head        // 开始时，快慢指针都指向头结点
    fast := head
    var meet *ListNode
    for fast != nil && fast.Next != nil {
        // 如果不存在环，快指针会先遍历完
        slow = slow.Next        // 链表没结束，快慢指针没相遇，慢指针一次往后走一步，快指针一次两步
        fast = fast.Next.Next  
        if(slow == fast) {      // 链表没结束，快慢指针相遇，有环，提前跳出
            meet = slow         // 快慢指针相遇点一定在环中，但不一定是环的起点
            break
        }
    }
    if(meet == nil) {           // 遍历链表结束，快慢指针没有相遇，则链表中不存在环
        return nil
    }

    p := head                   // 快慢指针能相遇，有环，则头结点和相遇处依次往后移动一个节点，再次相遇处则为环的起始节点
    for p != meet {
        p = p.Next
        meet = meet.Next
    }
    return p
}   

func cycleLen(head *ListNode) int {
	cycleNode := detectCycle(head)
	if(cycleNode == nil) {	// 不存在环，则环长为0
		return 0
	}

	start := cycleNode				// 存在环，将环起始位置记下
	var count int = 1				// 初始环长为1，即环的起始位置
	for start != cycleNode.Next {	// 当前节点与环的起始位置不同
		cycleNode = cycleNode.Next	// 指针后移一位
		count++						// 计数器加1
	}
	return count
}