func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	count:= 0
	l,r := 0, len(people)-1
	for l<r{
		if people[l]+people[r] <= limit{
			count++
			l++
			r--
		} else if people[r] <= limit{
			count++
			r--
		}
		if l==r{
			count++
			l++
			r--
		}
	}
	return count
}
