package peekingIterator

type Iterator struct {
	val int
	n   *Iterator
}

func (this *Iterator) hasNext() bool {
	return this.n != nil
}

func (this *Iterator) next() int {
	this = this.n
	return this.val
}

type PeekingIterator struct {
	iter    *Iterator
	hasnext bool
	peekVal int
}

func Constructor(iter *Iterator) *PeekingIterator {

	return &PeekingIterator{iter, iter.hasNext(), iter.next()}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasnext
}

func (this *PeekingIterator) next() int {
	r := this.peekVal
	this.hasnext = this.iter.hasNext()
	if this.hasnext {
		this.peekVal = this.iter.next()
	}

	return r
}

func (this *PeekingIterator) peek() int {
	return this.peekVal
}
