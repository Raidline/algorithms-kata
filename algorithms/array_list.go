package algorithms

type ArrayList struct {
	size   int
	arr    []int
	length int
}

func New(size int) *ArrayList {
	return &ArrayList{
		size:   size,
		arr:    make([]int, size),
		length: 0,
	}
}

func (a *ArrayList) InsertAt(idx int, value int) {
	if idx >= a.size || idx < 0 {
		println("No no no, bad idx")

		return
	}

	a.arr[idx] = value

	if a.length < idx {
		a.length++
	}
}

func (a *ArrayList) Append(value int) {
	a.arr[a.length] = value
	a.length++
}

func (a *ArrayList) Get(idx int) int {
	return a.arr[idx]
}

func (a *ArrayList) RemoveAt(idx int) int {
	if idx < 0 || idx > a.length {
		return -1
	}

	value := a.arr[idx]

	i := idx
	for ; i < a.length-1; i++ {
		a.arr[i] = a.arr[i+1]
	}

	a.arr[i] = -1

	a.length--
	return value
}

func (a *ArrayList) Remove(value int) int {
	for i := 0; i < a.size; i++ {
		if a.arr[i] == value {
			val := a.arr[i]
			a.arr[i] = -1
			for j := i; i < a.length-1; i++ {
				a.arr[j] = a.arr[j+1]
			}
			a.length--
			return val
		}
	}

	return -1
}

func (a *ArrayList) Length() int {
	return a.length
}

func (a *ArrayList) Prepend(value int) {
	if a.length == a.size {
		//throw some error
		return
	}

	for i := a.length; i > 0; i-- {
		a.arr[i] = a.arr[i-1]
	}

	a.arr[0] = value
	a.length++
}
