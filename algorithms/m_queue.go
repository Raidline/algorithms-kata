package algorithms

// should be generic, but i do not want to deal with go generics
// single linked list
type node struct {
	value int
	next  *node
}

type MQueue struct {
	Length int
	head   *node
	tail   *node
}

func NewQueue() *MQueue {
	return &MQueue{
		Length: 0,
	}
}

func (q *MQueue) Enqueue(value int) {
	q.Length++

	n := createNode(value, nil)
	if q.tail == nil || q.head == nil {
		q.tail = n
		q.head = n

		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *MQueue) Deque() int {
	if q.head == nil {
		return -1 //in real applications should return nil
	}
	q.Length--

	head := q.head
	q.head = q.head.next

	head.next = nil

	return head.value
}

func (q *MQueue) Peek() int {
	return q.head.value
}

func createTail(value int) *node {
	return &node{
		value: value,
		next:  nil,
	}
}

func createNode(value int, next *node) *node {
	return &node{
		value: value,
		next:  next,
	}
}
