package onlineStockSpan

type StockSpanner struct {
	m   map[int]int
	idx int
}

func Constructor() StockSpanner {
	m := make(map[int]int)
	return StockSpanner{m, 0}
}

func (this *StockSpanner) Next(price int) int {

	this.m[price] = this.idx
	this.idx++

	max := -1
	for k, v := range this.m {
		if k > price && v > max {
			max = v
		}
	}

	return this.idx - max - 1
}
