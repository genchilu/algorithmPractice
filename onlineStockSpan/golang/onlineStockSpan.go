package onlineStockSpan

type StockSpanner struct {
	prices *[]int
	counts *[]int
}

func Constructor() StockSpanner {
	prices := []int{}
	counts := []int{}
	return StockSpanner{&prices, &counts}
}

func (this *StockSpanner) Next(price int) int {
	c := 1
	for len((*this.prices)) > 0 && price >= (*this.prices)[len((*this.prices))-1] {
		c += (*this.counts)[len((*this.counts))-1]
		(*this.counts) = (*this.counts)[0 : len((*this.counts))-1]
		(*this.prices) = (*this.prices)[0 : len((*this.prices))-1]
	}

	(*this.prices) = append((*this.prices), price)
	(*this.counts) = append((*this.counts), c)

	return c
}
