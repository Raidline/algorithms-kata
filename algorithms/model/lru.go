package model

// Generics this time, as it is the last of the course, so why not?
type ILRU[K any, V any] interface {
	Get(K) (V, error) // error if there is no value
	Update(K, V)
}
