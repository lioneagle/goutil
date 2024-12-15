package scheduler

import (
	"github.com/lioneagle/goutil/src/algorithm"

	"github.com/pkg/errors"
)

type WeightNode[T any] struct {
	data   T
	weight int
}

type WeightedRoundRobin[T any] struct {
	curIndex  int
	curWeight int
	maxWeight int
	gcdWeight int
	nodes     []*WeightNode[T]
}

func NewWeightedRoundRobin[T any]() *WeightedRoundRobin[T] {
	return &WeightedRoundRobin[T]{
		curIndex: -1,
		nodes:    make([]*WeightNode[T], 0),
	}
}

func (this *WeightedRoundRobin[T]) Add(data T, weight int) error {
	if weight <= 0 {
		return errors.Errorf("weight (=%d) <= 0", weight)
	}
	this.nodes = append(this.nodes, &WeightNode[T]{data: data, weight: weight})

	if len(this.nodes) == 1 {
		this.gcdWeight = weight
	} else {
		this.gcdWeight = algorithm.Gcd(this.gcdWeight, weight)
	}

	if weight > this.maxWeight {
		this.maxWeight = weight
	}

	return nil
}

func (this *WeightedRoundRobin[T]) Schedule() (*T, error) {
	if len(this.nodes) <= 0 {
		return nil, errors.Errorf("no nodes to schedule")
	}

	for {
		this.curIndex = (this.curIndex + 1) % len(this.nodes)
		if this.curIndex == 0 {
			this.curWeight -= this.gcdWeight
			if this.curWeight <= 0 {
				this.curWeight = this.maxWeight
				if this.curWeight == 0 {
					return nil, errors.Errorf("this.curWeight is zero")
				}
			}
		}

		if this.nodes[this.curIndex].weight >= this.curWeight {
			return &this.nodes[this.curIndex].data, nil
		}
	}
}

func (this *WeightedRoundRobin[T]) Clone() *WeightedRoundRobin[T] {
	ret := &WeightedRoundRobin[T]{
		curIndex:  this.curIndex,
		curWeight: this.curWeight,
		maxWeight: this.maxWeight,
		gcdWeight: this.gcdWeight,
		nodes:     make([]*WeightNode[T], len(this.nodes)),
	}

	copy(ret.nodes, this.nodes)

	return ret
}
