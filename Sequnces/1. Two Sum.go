package Sequnces

func TwoSum(nums []int, target int) []int {
	visit := make(map[int]int)
	for i, num := range nums {
		if j, exist := visit[target-num]; exist {
			return []int{j, i}
		}
		visit[num] = i
	}
	return []int{}
}
