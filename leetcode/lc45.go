package main


func jump(nums []int) int {
	i, count, end := 0, 0, len(nums) -1
	var j, jV, m int
	for i < end {
		if i + nums[i] >= end {
			return count + 1
		}

		j, jV = i+1,i+ nums[i]
		for j <= jV  {
			if j + nums[j] > m {
				m, i = j+nums[j], j
			}
			j++
		}
		count++
	}
	return count
}
