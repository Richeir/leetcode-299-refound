/*
 * @lc app=leetcode.cn id=876 lang=golang
 * @lcpr version=30403
 *
 * [876] 链表的中间结点
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
func middleNode(head *ListNode) *ListNode {
	nodeCount := 1
	dummy := &ListNode{-1, head}
	for head.Next != nil {
		nodeCount++
		head = head.Next
	}

	head = dummy
	for i := 0; i < (nodeCount/2)+1; i++ {
		head = head.Next
	}

	return head
}

// @lc code=end

func TestMiddleOfTheLinkedList1(t *testing.T) {
	l1 := CreateHead([]int{1, 2, 3, 4, 5})
	result := middleNode(l1)
	PrintList(result)
}

func TestMiddleOfTheLinkedList2(t *testing.T) {
	l1 := CreateHead([]int{1, 2, 3, 4, 5, 6})
	result := middleNode(l1)
	PrintList(result)
}

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

*/
