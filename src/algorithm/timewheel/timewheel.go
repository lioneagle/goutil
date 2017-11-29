package algorithm

import (
	_ "container/list"
	//"sync"
	//"time"
)

type TimeWheelCallBack func(val interface{}) bool

type TimeWheelTaskData struct {
	wheel    int32
	slot     int32
	interval int64
	data     interface{}
	callBack TimeWheelCallBack
}

type slot struct {
	head int32
	size int32
}

type wheel struct {
	currentSlot int
	max         int64
	slots       []slot
}

type TimeWheel struct {
	size      int32
	delta     int64
	last      int64
	wheels    []wheel
	allocator *TimeWheelAllocator

	//mutex  sync.Mutex
}

func NewTimeWheel(timeWheelNum int, slotNum []int, delta int64, totalDataNum int32) *TimeWheel {
	if len(slotNum) != timeWheelNum {
		return nil
	}

	if delta <= 0 {
		return nil
	}

	tw := &TimeWheel{delta: delta}
	tw.wheels = make([]wheel, timeWheelNum)

	max := delta

	for i := 0; i < timeWheelNum; i++ {
		max *= int64(slotNum[i])
		v := &tw.wheels[i]
		v.max = max

		for j := 0; j < slotNum[i]; j++ {
			slot := slot{head: -1}
			v.slots = append(v.slots, slot)
		}
	}

	tw.allocator = NewTimeWheelAllocator(totalDataNum)
	return tw
}

func (this *TimeWheel) Size() int32 {
	return this.size
}

func (this *TimeWheel) Add(interval int64, data interface{}, callBack TimeWheelCallBack) int32 {
	chunk := this.allocator.AllocEx()
	if chunk == nil {
		return -1
	}

	wheel, slot := this.calcPos(interval)

	chunk.data.wheel = int32(wheel)
	chunk.data.slot = int32(slot)
	chunk.data.interval = interval
	chunk.data.data = data
	chunk.data.callBack = callBack

	// push_back to slot
	if this.wheels[wheel].slots[slot].head == -1 {
		chunk.next = chunk.id
		chunk.prev = chunk.id
		this.wheels[wheel].slots[slot].head = chunk.id
	} else {
		head := &this.allocator.Chunks[this.wheels[wheel].slots[slot].head]
		tail := &this.allocator.Chunks[head.prev]
		chunk.next = this.wheels[wheel].slots[slot].head
		chunk.prev = tail.id
		tail.next = chunk.id
		head.prev = chunk.id
	}
	this.wheels[wheel].slots[slot].size++

	this.size++

	return chunk.id
}

func (this *TimeWheel) calcPos(interval int64) (wheel, slot int32) {
	wheelNum := len(this.wheels)
	for i := 0; i < wheelNum; i++ {
		v := &this.wheels[i]
		//max := this.wheels[i].max
		size := int64(len(v.slots))
		//size := int64(len(this.wheels[i].slots))
		//if interval < max {
		if interval < v.max {
			return int32(i), int32(interval % size)
		}
		interval /= size
	}
	return -1, -1
}

func (this *TimeWheel) Remove(id int32) bool {
	chunk := &this.allocator.Chunks[id]
	if chunk.used == 0 {
		return false
	}

	wheel := chunk.data.wheel
	slot := chunk.data.slot

	this.size--

	// remove from slot
	if this.wheels[wheel].slots[slot].size == 1 {
		this.wheels[wheel].slots[slot].head = -1
	} else {
		this.allocator.Chunks[chunk.prev].next = chunk.next
		this.allocator.Chunks[chunk.next].prev = chunk.prev
		if this.wheels[wheel].slots[slot].head == chunk.id {
			this.wheels[wheel].slots[slot].head = chunk.next
		}
	}
	this.wheels[wheel].slots[slot].size--

	this.allocator.Free(id)

	return true
}

func (this *TimeWheel) Step(current int) {
	return
}
