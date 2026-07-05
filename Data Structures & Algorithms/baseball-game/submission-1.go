func calPoints(operations []string) int {
	record := []int{}
	for _, val := range operations{
		switch val{
			case "+":
				record = append(record, record[len(record)-1]+record[len(record)-2])
			case "C":
				record = record[:len(record)-1]
			case "D":
				record = append(record, 2*record[len(record)-1])
			default:
				num,_ := strconv.Atoi(val)
				record = append(record, num)
		}
	}
	ans := 0
	for _,val:= range record{
		ans+= val
	}
	return ans
}
