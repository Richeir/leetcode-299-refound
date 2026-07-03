/*
 * @lc app=leetcode.cn id=1094 lang=golang
 * @lcpr version=30403
 *
 * [1094] 拼车
 */

package leetcode_solutions

import "testing"

type Difference1094 struct {
	Diff []int
}

func NewDifference1094(nums []int) *Difference1094 {
	assert1094(len(nums) > 0)
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference1094{Diff: diff}
}

func (d *Difference1094) Increment(i, j, val int) {
	d.Diff[i] += val
	if j+1 < len(d.Diff) {
		d.Diff[j+1] -= val
	}
}

func (d *Difference1094) ConvertToResult() []int {
	res := make([]int, len(d.Diff))
	res[0] = d.Diff[0]
	for i := 1; i < len(d.Diff); i++ {
		res[i] = res[i-1] + d.Diff[i]
	}

	return res
}

func assert1094(condition bool) {
	if !condition {
		panic("condition failed")
	}
}

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	df := NewDifference1094(nums)
	for _, trip := range trips {
		numPasengers, from, to := trip[0], trip[1], trip[2]-1
		df.Increment(from, to, numPasengers)
	}

	result := df.ConvertToResult()
	for _, v := range result {
		if v > capacity {
			return false
		}
	}
	return true
}

// @lc code=end

func TestCarPooling(t *testing.T) {
	// your test code here
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 4
	result := carPooling(trips, capacity)
	println(result)
}

func TestCarPooling2(t *testing.T) {
	// your test code here
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 5
	result := carPooling(trips, capacity)
	println(result)
}

func TestCarPooling3(t *testing.T) {
	// your test code here
	trips := [][]int{{2, 1, 5}, {3, 5, 7}}
	capacity := 3
	result := carPooling(trips, capacity)
	println(result)
}

/*
// @lcpr case=start
// [[2,1,5],[3,3,7]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[2,1,5],[3,3,7]]\n5\n
// @lcpr case=end

*/
