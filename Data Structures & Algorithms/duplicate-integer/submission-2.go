func hasDuplicate(nums []int) bool {
    x:= len(nums)
	a:= make(map[int]int)
	for i:=0;i<x;i++{
		if a[nums[i]]>0{
			return true
		}
		a[nums[i]]++
	}
	return false
}
