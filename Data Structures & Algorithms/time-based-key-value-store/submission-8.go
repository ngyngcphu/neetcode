type Item struct {
	Timestamp int
	Val string
}

type TimeMap struct {
	data map[string][]Item
}

func Constructor() TimeMap {
	return TimeMap{data: make(map[string][]Item)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], Item{Timestamp: timestamp, Val: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.data[key]; !ok {
		return ""
	}
	item := this.data[key]
	left := 0
	right := len(item) - 1
	ans := ""
	for left <= right {
		mid := (left + right) / 2
		if item[mid].Timestamp <= timestamp {
			left = mid + 1
			ans = item[mid].Val
		} else {
			right = mid - 1
		}
	}
	return ans
}
