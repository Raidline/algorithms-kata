package algorithms

import "errors"

type LruNode[V comparable] struct {
	value V
	next  *LruNode[V]
	prev  *LruNode[V]
}

type LRU[K comparable, V comparable] struct {
	capacity   int
	len        int
	head       *LruNode[V]
	tail       *LruNode[V]
	keyToValue map[K]*LruNode[V]
	valueToMap map[*LruNode[V]]K
}

func NewLRU[K comparable, V comparable](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		capacity:   capacity,
		len:        0,
		head:       nil,
		tail:       nil,
		keyToValue: make(map[K]*LruNode[V]),
		valueToMap: make(map[*LruNode[V]]K),
	}
}

func (l *LRU[K, V]) Get(key K) (V, error) {
	var result V

	if v, b := l.keyToValue[key]; b {
		l.detach(v)
		l.prepend(v)
		return v.value, nil
	}

	return result, errors.New("Not found")
}

func (l *LRU[K, V]) Update(key K, value V) {

	if v, b := l.keyToValue[key]; b {
		//put the hotest value to the front
		l.detach(v)
		l.prepend(v)
		v.value = value
	} else {
		node := &LruNode[V]{
			value: value,
			next:  nil,
			prev:  nil,
		}

		l.len++
		l.prepend(node)
		l.trimCache()
		l.keyToValue[key] = node
		l.valueToMap[node] = key
	}
}

func (l *LRU[K, V]) detach(node *LruNode[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = l.head.next
	}

	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (l *LRU[K, V]) prepend(node *LruNode[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node

		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *LRU[K, V]) trimCache() {
	if l.len <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)

	key := l.valueToMap[tail]
	delete(l.keyToValue, key)
	delete(l.valueToMap, tail)
	l.len--
}
