package timeBasedKeyValueStore

type Item struct {
	t int
	v string
}

type TimeMap struct {
	m map[string][]Item
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	m := make(map[string][]Item)
	return TimeMap{m}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = []Item{}
	}

	this.m[key] = append(this.m[key], Item{t: timestamp, v: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if items, ok := this.m[key]; ok {
		l, r := 0, len(items)-1
		for l <= r {
			m := (l + r) / 2

			if items[m].t > timestamp {
				r = m - 1
			} else if items[m].t < timestamp {
				l = m + 1
			} else {
				return items[m].v
			}
		}

		if l == 0 {
			return ""
		} else {
			return items[l-1].v
		}
	}

	return ""
}