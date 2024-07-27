package Sequnces

func containsDuplicate(nums []int) bool {
	exist := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := exist[num]; ok {
			return true
		}
		exist[num] = struct{}{}
	}
	return false
}
