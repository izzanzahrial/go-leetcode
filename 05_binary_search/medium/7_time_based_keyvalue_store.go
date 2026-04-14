package binarysearch

// https://leetcode.com/problems/time-based-key-value-store/
type TimeMap struct {
	data map[string][]val
}

type val struct {
	v  string
	ts int
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]val),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], val{v: value, ts: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	vals, exists := this.data[key]
	if !exists {
		return ""
	}

	left, right, val := 0, len(vals)-1, ""
	for left <= right {
		mid := left + (right-left)/2
		if vals[mid].ts == timestamp {
			return vals[mid].v
		}

		// The key is here, since it will return the predecessor timestamp
		// if the timestamp you're looking for is not exists
		// meaning if the timestamp 5 is not there, but the most recent timestamp 3 have value
		// we return the timestamp 3
		if vals[mid].ts <= timestamp {
			val = vals[mid].v
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
