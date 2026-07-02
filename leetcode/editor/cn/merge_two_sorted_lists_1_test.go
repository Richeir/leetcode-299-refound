/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=30403
 *
 * [21] 合并两个有序链表
 */

package leetcode_solutions

import "testing"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

// @lc code=end

func TestMergeTwoSortedLists1(t *testing.T) {
	// your test code here
	l1 := CreateHead([]int{1, 2, 4})
	l2 := CreateHead([]int{1, 3, 4})

	result := mergeTwoLists1(l1, l2)
	PrintList(result)
}

func TestMergeTwoSortedLists2(t *testing.T) {
	// your test code here
	l1 := CreateHead([]int{})
	l2 := CreateHead([]int{})

	result := mergeTwoLists1(l1, l2)
	PrintList(result)
}

func TestMergeTwoSortedLists3(t *testing.T) {
	// your test code here
	l1 := CreateHead([]int{})
	l2 := CreateHead([]int{0})

	result := mergeTwoLists1(l1, l2)
	PrintList(result)
}

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/
