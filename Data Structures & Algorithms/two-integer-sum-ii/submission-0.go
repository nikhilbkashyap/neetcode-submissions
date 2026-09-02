func twoSum(numbers []int, target int) []int {
	st,top:=0,len(numbers)-1
	var num [] int
	for st<top{
		sum:=numbers[st]+numbers[top]
		if sum==target{
			num=append(num,st+1,top+1)
			return num
		}else if sum<target{
			st++
		}else{
			top--
		}
	}
	return nil
}