package arrayList

func removeElement(nums []int, val int) int {
	count := 0
	var i, j int
	for j < len(nums) {
		if nums[i] == val && nums[j] != val {
			count ++
			nums[i] = nums[j]
			nums[j] = val
			i ++
			j ++
		}else if nums[i] == val && nums[j] == val {
			j ++
		} else {
			i ++
			j ++
		}
	}
	return len(nums) - (j - i)
}

func removeElement2(nums []int, val int) int {
	length:=len(nums)
	res:=0
	for i:=0;i<length;i++{
		if nums[i]!=val {
			nums[res]=nums[i]
			res++
		}
	}
	return res
}
