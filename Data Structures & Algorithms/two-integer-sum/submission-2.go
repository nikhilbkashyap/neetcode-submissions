func twoSum(nums []int, target int) []int {
    for i:= range nums{
		for j:=i+1; j<len(nums);j++{
			if nums[i]+nums[j]==target && (i!=j){
			return []int{i,j}
		}
		}
	}
	return []int {0}
}
