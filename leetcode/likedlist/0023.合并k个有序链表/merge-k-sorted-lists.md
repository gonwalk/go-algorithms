
# 1.题目描述

（1）题目来源

LeetCode 23. 合并K个排序链表
https://leetcode-cn.com/problems/merge-k-sorted-lists/
难度：hard


（2）题目描述

合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6


# 2.算法实现

## 2.1 实现方式一：两两合并

将k个链表简化为双有序链表合并
（再优化：使用分治法）
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
      if len(lists) == 0 {
          return nil
      }
      li := lists[0]
      for i:=1; i<len(lists); i++ {
		  li = mergeTwoLists(li, lists[i])
	  }
	  return  li
}


//简化为合并2个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return  l2
	}else  if l2 == nil {
		return l1
	}else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	}else {
		l2.Next = mergeTwoLists(l1,l2.Next)
		return l2
	}
}
```


## 2.2 实现方式二：分治
思路：使用分治递归的方式

将合并k个有序链表（链表范围[l, r]，其中l=0, r=k-1）的方式，分解为以前半段（范围[l, m]，其中l=0, m= (l+r)/2内）和后半段（范围[m, r]，其中m=(l+r)/2，r=k内）的链表两两合并过程。
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if(len(lists) == 0) {
        return nil
    }

    return merge(lists, 0, len(lists) - 1)                // 对lists中的0到len(lists)-1个链表合并
}

func merge(lists []*ListNode, l, r int) *ListNode {       // 对列表中的l到r范围内的链表合并
    if(l == r) {
        return lists[l]
    }

    m := (l+r)/2
    return mergeTwoList(merge(lists,l, m), merge(lists, m+1, r))
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {     // 链表两两合并
    if(list1 == nil) {          // list1链表为空，合并后为list2
        return list2
    }
    if(list2 == nil) {
        return list1
    }
    if(list1.Val > list2.Val){  // 以较小的链表的头结点，进行递归插入
        list2.Next = mergeTwoList(list1, list2.Next)
        return list2
    } else {
        list1.Next = mergeTwoList(list1.Next, list2)
        return list1
    }
}
```
## 2.3 实现方式三：小顶堆

构建长度为k的小顶堆（优先队列）
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNodeHeap []*ListNode

func (lp *ListNodeHeap) Push(x interface{}) {
	*lp = append(*lp, x.(*ListNode))
}

func (lp *ListNodeHeap) Pop() interface{} {
	old := *lp
	n := len(old)
	x := old[n-1]
	*lp = old[0 : n-1]
	return x
}

func (lp ListNodeHeap) Len() int {
	return len(lp)
}

func (lp ListNodeHeap) Less(i, j int) bool {
	return lp[i].Val < lp[j].Val
}

func (lp ListNodeHeap) Swap(i, j int) {
	lp[i], lp[j] = lp[j], lp[i]
}

// 小顶堆
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	lp := &ListNodeHeap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(lp, lists[i])
		}
	}

	if lp.Len() == 0 {
		return nil
	}

	head := heap.Pop(lp).(*ListNode)
	cur := head
	nextToPush := cur.Next

	var next *ListNode
	for lp.Len() != 0 {
		if nextToPush != nil {
			heap.Push(lp, nextToPush)
		}
		next = heap.Pop(lp).(*ListNode)
		cur.Next = next
		cur = cur.Next
		nextToPush = next.Next
	}

	return head
}
```