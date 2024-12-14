package container

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		data: make(map[T]struct{}),
	}
}

func (this *HashSet[T]) Len() int {
	return len(this.data)
}

func (this *HashSet[T]) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *HashSet[T]) Add(val T) {
	this.data[val] = struct{}{}
}

func (this *HashSet[T]) Del(val T) {
	delete(this.data, val)
}

func (this *HashSet[T]) Has(val T) bool {
	_, ok := this.data[val]
	return ok
}

func (this *HashSet[T]) Clone() *HashSet[T] {
	ret := &HashSet[T]{
		data: make(map[T]struct{}, len(this.data)),
	}

	for k, _ := range this.data {
		ret.data[k] = struct{}{}
	}
	return ret
}

func (this *HashSet[T]) ToSlice() []T {
	ret := make([]T, 0, len(this.data))

	for key, _ := range this.data {
		ret = append(ret, key)
	}
	return ret
}

func (this *HashSet[T]) ForEach(
	op func(val T) (halt bool, err error),
) error {
	if op == nil {
		return nil
	}

	for key, _ := range this.data {
		halt, err := op(key)
		if err != nil {
			return err
		}

		if halt {
			break
		}
	}
	return nil
}
