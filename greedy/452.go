package greedy

import (
	"fmt"
	"sort"
)

type Test452 struct{}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		a := points[i]
		b := points[j]
		return a[0] < b[0]
	})
	//fmt.Println(points)
	var ans, flag, work int
	var left, right int
	for flag = 0; flag < len(points); {
		ans++
		left = points[flag][0]
		right = points[flag][1]
		//fmt.Println(points[flag])
		for work = flag + 1; work < len(points); {
			//fmt.Println(left, right)
			if (points[work][0] <= right && points[work][0] >= left) || (points[work][1] <= right && points[work][1] >= left) {
				left = max(points[work][0], left)
				right = min(points[work][1], right)
				work++
				continue
			} else {
				break
			}
		}
		flag = work
	}
	return ans
}

/*

别人的代码优雅一点

func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
        return 0
    }
    sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
    maxRight := points[0][1]
    ans := 1
    for _, p := range points {
        if p[0] > maxRight {
            maxRight = p[1]
            ans++
        }
    }
    return ans
}

*/

func (T Test452) Run() {
	var people [][]int = [][]int{
		{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8},
	}
	fmt.Println(findMinArrowShots(people))
}
