/*
 * @lc app=leetcode.cn id=303 lang=golang
 * @lcpr version=30403
 *
 * [303] 区域和检索 - 数组不可变
 */

package leetcode_solutions

import "testing"

// @lc code=start
type NumArray struct {
	PreSum []int
}

func Constructor303(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{PreSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.PreSum[right+1] - this.PreSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

func TestRangeSumQueryImmutable(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]\n
// @lcpr case=end

*/
