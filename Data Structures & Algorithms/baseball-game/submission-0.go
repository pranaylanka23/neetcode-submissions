func calPoints(operations []string) int {
	record := []int{}
	for _,val := range operations{
		switch val{
			case "C":
				record = record[:len(record)-1]
			case "D":
				record = append(record, record[len(record)-1]*2)
			case "+":
				sum := record[len(record)-1]+record[len(record)-2]
				record = append(record,sum)
			default:
				num,_ := strconv.Atoi(val)
				record = append(record,num)
		}
	}
	ans := 0
	for _,num := range record{
		ans += num
	}
	return ans
}
