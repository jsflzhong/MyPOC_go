package _001_two_sum

// example: [2, 7, 11, 15], target = 9
// 重点: go的: map[n-m] 这个用法很巧妙, 可以自动迭代里面的key查找是否有匹配的.
func twoSum(array []int, target int) []int {
	m := make(map[int]int)
	for key1, value := range array {
		//m[9-7]是否存在,存在就返回7的下标,和一个true
		if key2, ok := m[target-value]; ok {
			return []int{key1, key2}
		}
		m[value] = key1
	}
	return nil
}
