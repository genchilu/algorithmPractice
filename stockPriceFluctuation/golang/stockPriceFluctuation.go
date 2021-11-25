package stockPriceFluctuation

import "sort"

type StockPrice struct {
	m      map[int]int
	cur    int
	prices []int
}

func Constructor() StockPrice {
	return StockPrice{make(map[int]int), -1, []int{}}
}

func (this *StockPrice) Update(timestamp int, price int) {

	if timestamp > this.cur {
		this.cur = timestamp
	}

	if oldPrice, ok := this.m[timestamp]; ok {
		idx := sort.SearchInts(this.prices, oldPrice)
		lpart := this.prices[0:idx]
		rpart := this.prices[idx+1:]
		this.prices = append(lpart, rpart...)
	}

	this.m[timestamp] = price
	idx := sort.SearchInts(this.prices, price)
	this.prices = append(this.prices, 0)
	copy(this.prices[idx+1:], this.prices[idx:])
	this.prices[idx] = price

}

func (this *StockPrice) Current() int {

	return this.m[this.cur]
}

func (this *StockPrice) Maximum() int {
	return this.prices[len(this.prices)-1]
}

func (this *StockPrice) Minimum() int {
	return this.prices[0]
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
