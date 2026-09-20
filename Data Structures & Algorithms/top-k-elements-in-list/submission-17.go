func topKFrequent(nums []int, k int) []int {
	m:= make(map[int]int)
	n:= make(map[int][]int)
	for i:=0; i< len(nums);i++{
		m[nums[i]]++
	}
	temp:= []int{}
	var a []int
	for i,v:= range m{
		n[v]=append(n[v],i)
		temp=append(temp,v)
	}
	seen := make(map[int]bool)
	sort.Ints(temp)
	for i:=(len(temp)-1);i>=(len(temp)-k);i--{
			freq := temp[i]
		if seen[freq] {
			continue
		}
		seen[freq] = true
			a=append(a,n[temp[i]]...)
		}
	return a[:k]
}
