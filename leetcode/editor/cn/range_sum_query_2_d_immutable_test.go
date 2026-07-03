/*
 * @lc app=leetcode.cn id=304 lang=golang
 * @lcpr version=30403
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

package leetcode_solutions

import "testing"

// @lc code=start
type NumMatrix struct {
	PreSum2D [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// m 行数
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}

	// n 列数
	n := len(matrix[0])
	if n == 0 {
		return NumMatrix{}
	}

	// 初始化前序和数组的行、列
	preSum2D := make([][]int, m+1)
	// for i := 0; i <= m; i++ {
	// 	preSum2D[i] = make([]int, n+1)
	// }

	for i := range preSum2D {
		preSum2D[i] = make([]int, n+1)
	}

	// 循环每行
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSum2D[i][j] = preSum2D[i-1][j] + preSum2D[i][j-1] + matrix[i-1][j-1] - preSum2D[i-1][j-1]
		}
	}

	return NumMatrix{PreSum2D: preSum2D}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.PreSum2D[row2+1][col2+1] - this.PreSum2D[row1][col2+1] - this.PreSum2D[row2+1][col1] + this.PreSum2D[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

func TestRangeSumQuery2dImmutable(t *testing.T) {
	// your test code here
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	obj := Constructor(matrix)
	param_1 := obj.SumRegion(2, 1, 4, 3)
	println(param_1)
}

/*
// @lcpr case=start
// ["NumMatrix","sumRegion","sumRegion","sumRegion"]\n[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]\n
// @lcpr case=end

*/
