func isAnagram(s string, t string) bool {
	x,y:= make(map[byte]int),make(map[byte]int)
	if len(s)!=len(t){
		return false
	}
	for i:=0;i<len(s);i++{
			x[s[i]]++
			y[t[i]]++
	}
	for i,v:= range x{
		if y[i]!= v{
			return false
		}		
	}		
	return true
}
