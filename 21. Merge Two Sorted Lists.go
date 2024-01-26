// 1/26
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func appendListLast(li *ListNode, v int) *ListNode {
	if li == nil {
		return &ListNode{Val: v}
	}
	start := li
	for li.Next != nil {
		li = li.Next
	}
	li.Next = &ListNode{Val: v}
	return start
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ansList *ListNode
	for list1 != nil || list2 != nil {
		if list1 == nil {
			ansList = appendListLast(ansList, list2.Val)
			list2 = list2.Next
		} else if list2 == nil {
			ansList = appendListLast(ansList, list1.Val)
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			ansList = appendListLast(ansList, list1.Val)
			list1 = list1.Next
		} else {
			ansList = appendListLast(ansList, list2.Val)
			list2 = list2.Next
		}
	}
	return ansList
}

func main() {
	var list1, list2 *ListNode
	var ans []int

	// list1[1, 2, 4]
	// list2[1, 3, 4]
	list1 = appendListLast(list1, 1)
	list1 = appendListLast(list1, 2)
	list1 = appendListLast(list1, 4)
	list2 = appendListLast(list2, 1)
	list2 = appendListLast(list2, 3)
	list2 = appendListLast(list2, 4)
	ret := mergeTwoLists(list1, list2)
	for ret != nil {
		ans = append(ans, ret.Val)
		ret = ret.Next
	}
	fmt.Println(ans)
	println("---")

	// list1[]
	// list2[]
	list1, list2 = nil, nil
	ans = nil
	ret = mergeTwoLists(list1, list2)
	for ret != nil {
		ans = append(ans, ret.Val)
		ret = ret.Next
	}
	fmt.Println(ans)
	println("---")

	// list1[]
	// list2[0]
	list1 = nil
	list2 = appendListLast(list2, 0)
	ret = mergeTwoLists(list1, list2)
	for ret != nil {
		ans = append(ans, ret.Val)
		ret = ret.Next
	}
	fmt.Println(ans)

	return
}
