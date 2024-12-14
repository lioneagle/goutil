package container

import (
	_ "fmt"
	"sort"

	"github.com/lioneagle/goutil/src/constraints"

	"github.com/pkg/errors"
)

type OderedMap[KEY constraints.Ordered, VAL any] struct {
	order []KEY
	data  map[KEY]VAL
}

func NewOrderedMap[KEY constraints.Ordered, VAL any]() *OderedMap[KEY, VAL] {
	return &OderedMap[KEY, VAL]{
		order: make([]KEY, 0),
		data:  make(map[KEY]VAL),
	}
}

func (this *OderedMap[KEY, VAL]) Len() int {
	return len(this.data)
}

func (this *OderedMap[KEY, VAL]) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *OderedMap[KEY, VAL]) Add(key KEY, val VAL) {
	_, ok := this.data[key]
	if !ok {
		this.order = append(this.order, key)
	}
	this.data[key] = val
}

func (this *OderedMap[KEY, VAL]) Del(key KEY) {
	_, ok := this.data[key]
	if !ok {
		return
	}
	delete(this.data, key)
	for i := 0; i < len(this.order); i++ {
		if this.order[i] == key {
			this.order = append(this.order[:i], this.order[i+1:]...)
			break
		}
	}
}

func (this *OderedMap[KEY, VAL]) Find(key KEY) (VAL, bool) {
	val, ok := this.data[key]
	return val, ok
}

func (this *OderedMap[KEY, VAL]) Clone() *OderedMap[KEY, VAL] {
	ret := &OderedMap[KEY, VAL]{
		order: make([]KEY, len(this.order), len(this.order)),
		data:  make(map[KEY]VAL, len(this.data)),
	}

	copy(ret.order, this.order)
	for k, v := range this.data {
		ret.data[k] = v
	}
	return ret
}

func (this *OderedMap[KEY, VAL]) ToSlice() ([]KEY, []VAL) {
	key := make([]KEY, len(this.order), len(this.order))
	val := make([]VAL, 0, len(this.data))

	copy(key, this.order)
	for i := 0; i < len(this.order); i++ {
		val = append(val, this.data[this.order[i]])
	}
	return key, val
}

func (this *OderedMap[KEY, VAL]) ForEach(
	op func(key KEY, val VAL) (halt bool, err error),
) error {
	if op == nil {
		return nil
	}

	for _, key := range this.order {
		val, ok := this.data[key]
		if !ok {
			return errors.Errorf("OrderedMap Foreach: cannot find by key \"%v\"", key)
		}
		halt, err := op(key, val)
		if err != nil {
			return err
		}

		if halt {
			break
		}
	}
	return nil
}

func (this *OderedMap[KEY, VAL]) Sort() {
	sort.Slice(this.order, func(i, j int) bool {
		return this.order[i] < this.order[j]
	})
}

func (this *OderedMap[KEY, VAL]) SortStable() {
	sort.SliceStable(this.order, func(i, j int) bool {
		return this.order[i] < this.order[j]
	})
}

func (this *OderedMap[KEY, VAL]) SortReverse() {
	sort.Slice(this.order, func(i, j int) bool {
		return this.order[i] > this.order[j]
	})
}

func (this *OderedMap[KEY, VAL]) SortReverseStable() {
	sort.SliceStable(this.order, func(i, j int) bool {
		return this.order[i] > this.order[j]
	})
}
