package algorithms

type LRU[K any, V any] struct {
	capacity int
}

func NewLRU[K any, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		capacity: capacity,
	}
}

func (l *LRU[K, V]) Get(key K) (V, error) {
	var result V
	return result, nil
}

func (l *LRU[K, V]) Update(key K, value V) {

}
