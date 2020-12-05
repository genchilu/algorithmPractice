package myCalendarI

type MyCalendar struct {
	events [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{[][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if end <= start {
		return false
	}

	e := []int{start, end}
	for _, v := range this.events {
		if isConflict(e, v) {
			return false
		}
	}
	this.events = append(this.events, e)
	return true
}

func isConflict(a, b []int) bool {
	if a[1] <= b[0] || a[0] >= b[1] {
		return false
	}

	return true
}
