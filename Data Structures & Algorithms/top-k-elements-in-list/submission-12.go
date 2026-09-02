func topKFrequent(nums []int, k int) []int {
	freq:= make(map[int]int)
	for _,n:= range nums{
		freq[n]++
	}
	var temp []int
	seen := make(map[int]bool)
	for _, count := range freq {
		if !seen[count] {
			seen[count] = true
			temp = append(temp, count)
		}
	}
	sort.Ints(temp)
	var num []int
	for i:=(len(temp)-1); i>=0 && len(num)<k; i--{
		for id,v:= range freq {
			if temp[i]==v{
				num=append(num,id)
				if len(num)==k{
					break
				}
			}
		}
	}
	return num
}