func groupAnagrams(strs []string) [][]string {
	m:= make(map[string][]string)
	for i:=0;i<len(strs);i++{
		b:=[]byte(strs[i])
		sort.Slice(b,func(i,j int)bool{
			return b[i]<b[j]
		})
		key:=string(b)
		m[key]=append(m[key],strs[i])
	}
	var n[][] string
	for _,v:= range m{
		n=append(n,v)
	}
	return n
}
