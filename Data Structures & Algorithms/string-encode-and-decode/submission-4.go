type Solution struct{}

func (s *Solution) Encode(strs []string) string {
var encoded string
if len(strs) == 0{
	return ""
}
for _,i:= range strs{
	encoded+=strconv.Itoa(len(i))+"#"+i
}	
return encoded
}

func (s *Solution) Decode(encoded string) []string {
	var decoded []string
	var i int = 0
	if encoded ==""{
		return []string{}
	}
	for i < len(encoded){
		j:=i
		for encoded[j]!= '#'{
			j++
		}
		length,_:= strconv.Atoi(encoded[i:j])
		st:=j+1
		end:=st+length
		decoded=append(decoded,encoded[st:end])
		i=end
		}
	return decoded	
	}
	

