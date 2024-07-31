package xl_leetcode

import (
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) int {
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
	sort.Ints(xs)
	res, cur := 1, xs[0]
	for _, s := range xs {
		if cur+w < s {
			cur = s
			res++
		}
	}
	return res
}
