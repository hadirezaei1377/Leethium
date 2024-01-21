package twosum

func TwoSum(nums []int, target int) []int {
	res := []int{}

	m := make(map[int]int)

	for i, num := range nums {
		if _, ok := m[num]; ok {
			res = append(res, m[num], i)
		}

		m[target-num] = i
	}

	return res
}
