/*
 * @lc app=leetcode.cn id=160 lang=golang
 * @lcpr version=30403
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB

	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

// @lc code=end

func TestIntersectionOfTwoLinkedLists(t *testing.T) {
	// your test code here
	lcommon := CreateHead([]int{8, 4, 5})
	l1 := CreateHead([]int{4, 1})
	l1.Next.Next = lcommon
	l2 := CreateHead([]int{5, 6, 1})
	l2.Next.Next.Next = lcommon

	result := getIntersectionNode(l1, l2)
	PrintList(result)
}

/*
// @lcpr case=start
// 8\n[4,1,8,4,5]\n[5,6,1,8,4,5]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n[1,9,1,2,4]\n[3,2,4]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 0\n[2,6,4]\n[1,5]\n3\n2\n
// @lcpr case=end

*/
