package main


func wiggleMaxLength(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	//begin := 0
	//up := 1
	//down := 2
	res := 1
	state := 0

	for i := 1; i < length; i++ {
		switch state {
		case 0:
			if nums[i-1] < nums[i] {
				state = 1
				res++
			} else if (nums[i-1] > nums[i]) {
				state = 2
				res++
			}
		case 1:
			if nums[i-1] > nums[i] {
				state = 2
				res++
			}
		case 2:
			if nums[i-1] < nums[i] {
				state = 1
				res++
			}
		}
	}

	return res
}
