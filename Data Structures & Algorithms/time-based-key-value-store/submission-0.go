type TimeMap struct {
	m map[string][]pair
}

type pair struct{
	timestamp int
	value string
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _,exists := this.m[key]; !exists{ return ""}
	l,r := 0,len(this.m[key])-1
	pairs := this.m[key]
	for l<r{
		mid:= l + (r-l+1)/2
		if pairs[mid].timestamp == timestamp {
			return pairs[mid].value
		}
		if pairs[mid].timestamp>timestamp{
			r = mid-1
		}else{
			l=mid
		}
	}
	if pairs[l].timestamp > timestamp { return ""}
	return pairs[l].value
}
