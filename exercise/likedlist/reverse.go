package main

import(
	"fmt"
)

type LikedList struct{
    Val int
    Next *LikedList
}

func main(){
	node1 := &LikedList{1,nil}
	node2 := &LikedList{2,nil}
	node3 := &LikedList{3,nil}
	node4 := &LikedList{4,nil}
	node5 := &LikedList{5,nil}
	node6 := &LikedList{6,nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil
	
	head1 := node1
	fmt.Print("翻转前链表内容：")
	for head1 != nil {
		fmt.Printf("%d ->", head1.Val)
		head1 = head1.Next
	}
	fmt.Println()

	node1 = reverseLikedList(node1)
	head2 := node1
	fmt.Print("翻转后链表内容：")
	for head2 != nil {
		fmt.Printf("%d ->", head2.Val)
		head2 = head2.Next
	}
}

// 头插法，倒置单链表
func reverseLikedList(head *LikedList)*LikedList{
    if(head == nil || head.Next == nil){
        return head
    }
	var pre, next *LikedList 
	// head是当前节点，pre是当前节点的前一个几点，next是当前节点的后一个节点
    for head != nil {		// 当前节点不为空
        next = head.Next    // 备份head.Next（指向当前节点的下一个节点），防止移动节点时找不到后置节点
		head.Next = pre     // 更新head.Next，后一个节点指向前一个（翻转）
		pre = head          // 前一个节点后移
		head = next			// 当前节点后移
	}
    return pre				// 注意：当pre指向原链表的最后一个节点时，head已经指向最后一个节点的Next即空节点，所以这里应返回pre
}