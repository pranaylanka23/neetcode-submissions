type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	ans := ""
	for _, str := range strs{
		ans += strconv.Itoa(len(str))+"#"+str
	}
	return ans
}

func (s *Solution) Decode(encoded string) []string {
	ans := []string{}
	for i:=0; i< len(encoded);{
		j:=i
		for encoded[j]!= '#'{
			j++
		}
		length,_ := strconv.Atoi(encoded[i:j])
		i = j+1
		ans = append(ans,encoded[i:i+length])
		i += length
	}
	return ans
}
