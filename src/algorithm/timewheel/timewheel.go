package algorithm

import (
	_ "container/list"
	//"sync"
	//"time"
)

type TimeWheelCallBack func(val interface{})

type TimeWheelTaskData struct {
	wheel    int32
	slot     int32
	interval int64
	data     interface{}
	callBack TimeWheelCallBack
}

type slot struct {
	head     int32
	size     int32
	interval int32
}

type wheel struct {
	max   int64
	slots []slot
}

type TimeWheel struct {
	isBinaryBits bool
	size         int32
	currentWheel int32
	currentSlot  int32
	delta        int64
	current      int64
	wheels       []wheel
	wheelBits    []uint32
	wheelMasks   []int32
	allocator    *TimeWheelAllocator
	stat         TimeWheelStat

	//mutex  sync.Mutex
}

type TimeWheelStat struct {
	Add             int64
	AddOk           int64
	Remove          int64
	RemoveOk        int64
	Step            int64
	Expire          int64
	ExpireBeforeAdd int64
	Post            int64
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

func NewTimeWheelBinaryBits(timeWheelNum int, slotNumBit []int, delta int64, totalDataNum int32) *TimeWheel {
	var slotNum []int

	for i := 0; i < len(slotNumBit); i++ {
		slotNum = append(slotNum, 2<<uint32(slotNumBit[i]))
	}

	tw := NewTimeWheel(timeWheelNum, slotNum, delta, totalDataNum)
	if tw == nil {
		return nil
	}

	tw.isBinaryBits = true

	for i := 0; i < len(slotNumBit); i++ {
		tw.wheelBits = append(tw.wheelBits, uint32(slotNumBit[i]))
		tw.wheelMasks = append(tw.wheelMasks, int32(2<<uint32(slotNumBit[i])-1))
	}

	return tw
}

func (this *TimeWheel) Size() int32 {
	return this.size
}

func (this *TimeWheel) Add(interval int64, data interface{}, callBack TimeWheelCallBack) int32 {
	this.stat.Add++

	if interval < this.current {
		this.stat.Expire++
		this.stat.ExpireBeforeAdd++
		if callBack != nil {
			this.stat.Post++
			callBack(data)
			return -2
		}
	}

	chunk := this.allocator.AllocEx()
	if chunk == nil {
		return -1
	}

	var wheelIndex int32
	var slotIndex int32

	if this.isBinaryBits {
		wheelIndex, slotIndex = this.calcPosBinaryBits(interval)
	} else {
		wheelIndex, slotIndex = this.calcPos(interval)
	}

	chunk.data.wheel = int32(wheelIndex)
	chunk.data.slot = int32(slotIndex)
	chunk.data.interval = interval
	chunk.data.data = data
	chunk.data.callBack = callBack

	slot := &this.wheels[wheelIndex].slots[slotIndex]

	// push_back to slot
	if slot.head == -1 {
		chunk.next = chunk.id
		chunk.prev = chunk.id
		slot.head = chunk.id
	} else {
		head := &this.allocator.Chunks[slot.head]
		tail := &this.allocator.Chunks[head.prev]
		chunk.next = slot.head
		chunk.prev = tail.id
		tail.next = chunk.id
		head.prev = chunk.id
	}

	slot.size++
	this.size++
	this.stat.AddOk++

	return chunk.id
}

func (this *TimeWheel) calcPos(interval int64) (wheel, slot int32) {
	wheelNum := len(this.wheels)
	for i := 0; i < wheelNum; i++ {
		v := &this.wheels[i]
		size := int64(len(v.slots))
		if interval < v.max {
			return int32(i), int32(interval % size)
		}
		interval /= size
	}
	return -1, -1
}

func (this *TimeWheel) calcPosBinaryBits(interval int64) (wheel, slot int32) {
	wheelNum := len(this.wheels)
	for i := 0; i < wheelNum; i++ {
		if interval < this.wheels[i].max {
			return int32(i), int32(interval) & this.wheelMasks[i]
		}

		interval >>= this.wheelBits[i]
	}
	return -1, -1
}

func (this *TimeWheel) Remove(id int32) bool {
	this.stat.Remove++
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

	this.stat.RemoveOk++

	return true
}

func (this *TimeWheel) Step(current int64) {
	this.stat.Step++

	if current < this.current {
		return
	}

	wheelNum := len(this.wheels)
	for i := 0; i < wheelNum; i++ {
		wheelIndex = (this.currentWheel + i) % wheelNum
		slotIndex = this.currentSlot

		if interval < this.wheels[i].max {
			return int32(i), int32(interval) & this.wheelMasks[i]
		}

		interval >>= this.wheelBits[i]
	}

	return
}
