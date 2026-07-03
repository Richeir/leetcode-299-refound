package leetcode_solutions

import (
	"fmt"
	"testing"
)

type Difference struct {
	Diff []int
}

func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{Diff: diff}
}

func (d *Difference) Increment(i, j, val int) {
	d.Diff[i] += val
	if j+1 < len(d.Diff) {
		d.Diff[j+1] -= val
	}
}

func (d *Difference) ConvertToResult() []int {
	res := make([]int, len(d.Diff))
	res[0] = d.Diff[0]
	for i := 1; i < len(d.Diff); i++ {
		res[i] = res[i-1] + d.Diff[i]
	}

	return res
}

func assert(condition bool) {
	if !condition {
		panic("condition failed")
	}
}

func getModifiedArry(length int, updates [][]int) []int {
	nums := make([]int, length)
	df := NewDifference(nums)
	for _, update := range updates {
		i, j, val := update[0], update[1], update[2]
		df.Increment(i, j, val)
	}

	return df.ConvertToResult()
}

func TestRangeAddition(t *testing.T) {
	// your test code here
	length := 5
	updates := [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}
	result := getModifiedArry(length, updates)
	fmt.Println(result)
}
