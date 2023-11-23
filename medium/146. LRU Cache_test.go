package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type LRUCache struct {
	list []lruData
	cap  int
}

type lruData struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{list: make([]lruData, 0, capacity), cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	for k, data := range this.list {
		if data.key == key {
			this.list = append(this.list[:k], this.list[k+1:]...)
			this.list = append([]lruData{data}, this.list...)
			return data.value
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	for i := 0; i < len(this.list); i++ {
		if this.list[i].key == key {
			this.list = append(this.list[:i], this.list[i+1:]...)
			break
		}
	}
	this.list = append([]lruData{{key: key, value: value}}, this.list...)
	for len(this.list) > this.cap {
		this.list = this.list[:len(this.list)-1]
	}
}

func TestLRUCache(t *testing.T) {
	const notfound = -1
	// LRUCache lRUCache = new LRUCache(2);
	lruCache := Constructor(2)

	// lRUCache.put(1, 1); // cache is {1=1}
	lruCache.Put(1, 1)

	// lRUCache.put(2, 2); // cache is {1=1, 2=2}
	lruCache.Put(2, 2)

	// lRUCache.get(1);    // return 1
	require.Equal(t, lruCache.Get(1), 1)

	// lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lruCache.Put(3, 3)

	// lRUCache.get(2);    // returns -1 (not found)
	require.Equal(t, lruCache.Get(2), notfound)

	// lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	lruCache.Put(4, 4)

	// lRUCache.get(1);    // return -1 (not found)
	require.Equal(t, lruCache.Get(1), notfound)

	// lRUCache.get(3);    // return 3
	require.Equal(t, lruCache.Get(3), 3)

	// lRUCache.get(4);    // return 4
	require.Equal(t, lruCache.Get(4), 4)
}

func TestLRUCache1(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(2, 6)
	lruCache.Put(1, 5)
	lruCache.Put(1, 2)
	require.Equal(t, lruCache.Get(1), 2)
	require.Equal(t, lruCache.Get(2), 6)
}
