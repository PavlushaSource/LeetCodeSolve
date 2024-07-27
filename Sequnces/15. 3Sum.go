package Sequnces

import (
	"sort"
)

func threeSum(nums []int) (ans [][]int) {
	if len(nums) < 3 {
		return
	}
	sort.Ints(nums)
	for index1 := range nums {
		if index1 > 0 && nums[index1-1] == nums[index1] {
			continue
		}
		index2 := index1 + 1
		index3 := len(nums) - 1
		for index2 < index3 {
			sum := nums[index1] + nums[index2] + nums[index3]
			if sum == 0 {
				ans = append(ans, []int{nums[index1], nums[index2], nums[index3]})
			}
			if sum >= 0 {
				index3--
				for ; index2 < index3 && nums[index3+1] == nums[index3]; index3-- {
				}
			} else {
				index2++
				for ; index2 < index3 && nums[index2-1] == nums[index2]; index2++ {
				}
			}
		}
	}
	return
}
