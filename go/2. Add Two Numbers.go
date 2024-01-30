// 1/31
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func appendListLast(li, value *ListNode) *ListNode {
	if li == nil {
		return value
	}
	ret := li
	for li.Next != nil {
		li = li.Next
	}
	li.Next = value

	return ret
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ansList *ListNode
	var carry int

	for l1 != nil || l2 != nil || carry != 0 {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		tmp := n1 + n2 + carry
		carry = tmp / 10
		ansList = appendListLast(ansList, &ListNode{Val: tmp % 10})
	}

	return ansList
}

type TestSet struct {
	li1 []int
	li2 []int
}

func NewLinkList(ts TestSet) (*ListNode, *ListNode) {
	var li1, li2 *ListNode

	for _, v := range ts.li1 {
		tmp := &ListNode{Val: v}
		li1 = appendListLast(li1, tmp)
	}
	for _, v := range ts.li2 {
		tmp := &ListNode{Val: v}
		li2 = appendListLast(li2, tmp)
	}

	return li1, li2
}

func main() {
	inputs := []TestSet{
		{[]int{2, 4, 3}, []int{5, 6, 4}},
		{[]int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
	}

	for _, v := range inputs {
		ret := addTwoNumbers(NewLinkList(v))
		fmt.Println(ret)
	}

	return
}
