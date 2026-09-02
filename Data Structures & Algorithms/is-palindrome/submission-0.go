func isPalindrome(s string) bool {
	var temp []byte
	for i:=0;i<len(s);i++{
		c:=s[i]
		if(c>='a'&&c<='z')||(c>='0'&&c<='9'){
			temp=append(temp,c)
		}else if (c>='A'&&c<='Z'){
			temp=append(temp,c+32)
		}
	}
	st,top:=0, (len(temp)-1)
	for st<top {
		if(temp[st]==temp[top]){
			st++
			top--
		}else{
			return false
		}
	}
	return true
}
