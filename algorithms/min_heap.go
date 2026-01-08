package algorithms

type MinHeap struct {
	length int
	items  []int
}

func NewHeap() *MinHeap {
	return &MinHeap{
		length: 0,
		items:  make([]int, 0), // i could have used an arraylist to make this grow dinamically, this is just easier
	}
}

func (h *MinHeap) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	if h.length == 0 {
		return -1
	}

	out := h.items[0]
	h.length--

	if h.length == 0 {
		h.items = make([]int, 0)
		return out
	}

	h.items[0] = h.items[h.length]
	h.heapifyDown(0)

	return out
}

func (h *MinHeap) Length() int {
	return h.length
}

func (h *MinHeap) heapifyDown(idx int) {
	lIdx := left(idx)
	rIdx := right(idx)

	if idx >= h.length || lIdx >= h.length {
		return
	}

	curr := h.items[idx]
	lv := h.items[lIdx]
	lr := h.items[rIdx]

	if lv <= lr && lv < curr {
		//compare with l
		h.items[lIdx] = curr
		h.items[idx] = lv
		h.heapifyDown(lIdx)
	} else if lr < curr {
		//compare with r
		h.items[rIdx] = curr
		h.items[idx] = lr
		h.heapifyDown(rIdx)
	}
}

func (h *MinHeap) heapifyUp(idx int) {

	if idx == 0 {
		return
	}

	curr := h.items[idx]
	parentIdx := parent(idx)
	parentValue := h.items[parentIdx]

	if parentValue > curr { //min heap so the lower go up top
		h.items[parentIdx] = curr
		h.items[idx] = parentValue
		h.heapifyUp(parentIdx)
	}

}

func left(idx int) int {
	return (2 * idx) + 1
}

func right(idx int) int {
	return (2 * idx) + 2
}

func parent(idx int) int {
	return (idx - 1) / 2
}
