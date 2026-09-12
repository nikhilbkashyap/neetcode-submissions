func groupAnagrams(strs []string) [][]string {
	n:= make(map[string][]string)
	for i:=0;i<len(strs);i++{
		b:= []byte(strs[i])
		sort.Slice(b,func(i,j int)bool{
			return b[i]<b[j]
		})
		key:=string(b)
		n[key]=append(n[key],strs[i])
	}
	var m [][]string
	for _,v:= range n{
		m= append(m,v)
	}
	return m
}
