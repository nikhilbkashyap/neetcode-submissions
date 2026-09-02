func search(nums []int, target int) int {
	st,top:= 0,((len(nums))-1)
	for st<=top{
		var mid int =st+(top-st)/2
		if nums[mid]==target{
			return mid
		}else if target>nums[mid]{
			st=mid+1
		}else{
			top=mid-1
		}
	}
			return -1
}
