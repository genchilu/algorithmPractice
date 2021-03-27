package designHitCounter

type HitCounter struct {
	hit_time  []int
	hit_count []int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{make([]int, 300), make([]int, 300)}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	t := timestamp % 300
	if this.hit_time[t] != timestamp {
		this.hit_time[t] = timestamp
		this.hit_count[t] = 1
	} else {
		this.hit_count[t]++
	}
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	s := 0
	for i, _ := range this.hit_count {
		if this.hit_time[i] > timestamp-300 {
			s += this.hit_count[i]
		}
	}
	return s
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
