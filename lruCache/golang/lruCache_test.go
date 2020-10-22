package lruCache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestLRUCache1(t *testing.T) {
// 	cache := Constructor(2)
// 	var actual int

// 	cache.Put(1, 1)
// 	cache.Put(2, 2)

// 	actual = cache.Get(1)
// 	assert.Equal(t, 1, actual)

// 	cache.Put(3, 3)

// 	actual = cache.Get(2)
// 	assert.Equal(t, -1, actual)

// 	cache.Put(4, 4)

// 	actual = cache.Get(1)
// 	assert.Equal(t, -1, actual)

// 	actual = cache.Get(3)
// 	assert.Equal(t, 3, actual)

// 	actual = cache.Get(4)
// 	assert.Equal(t, 4, actual)
// }

func TestLRUCache2(t *testing.T) {
	cache := Constructor(2)
	var actual int

	cache.Put(2, 1)
	cache.Put(2, 2)

	actual = cache.Get(2)
	assert.Equal(t, 2, actual)

	cache.Put(1, 4)
	cache.Put(4, 1)

	actual = cache.Get(2)
	assert.Equal(t, -1, actual)
}
