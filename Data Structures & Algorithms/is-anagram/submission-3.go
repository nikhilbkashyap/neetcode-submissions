func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	x,y:= make(map[byte]int),make(map[byte]int)
	
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
