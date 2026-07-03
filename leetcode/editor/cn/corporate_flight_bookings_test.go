/*
 * @lc app=leetcode.cn id=1109 lang=golang
 * @lcpr version=30403
 *
 * [1109] 航班预订统计
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

type Difference1109 struct {
	Diff []int
}

func NewDifference1109(nums []int) *Difference1109 {
	assert1109(len(nums) > 0)
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference1109{Diff: diff}
}

func (d *Difference1109) Increment1109(i, j, val int) {
	d.Diff[i] += val
	if j+1 < len(d.Diff) {
		d.Diff[j+1] -= val
	}
}

func (d *Difference1109) ConvertToResult1109() []int {
	res := make([]int, len(d.Diff))
	res[0] = d.Diff[0]
	for i := 1; i < len(d.Diff); i++ {
		res[i] = res[i-1] + d.Diff[i]
	}

	return res
}

func assert1109(condition bool) {
	if !condition {
		panic("condition failed")
	}
}

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	df := NewDifference1109(nums)
	for _, booking := range bookings {
		i, j, val := booking[0]-1, booking[1]-1, booking[2]
		df.Increment1109(i, j, val)
	}

	return df.ConvertToResult1109()
}

// @lc code=end

func TestCorporateFlightBookings(t *testing.T) {
	// your test code here
	n := 5
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	result := corpFlightBookings(bookings, n)
	fmt.Println(result)
}

/*
// @lcpr case=start
// [[1,2,10],[2,3,20],[2,5,25]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,10],[2,2,15]]\n2\n
// @lcpr case=end

*/
