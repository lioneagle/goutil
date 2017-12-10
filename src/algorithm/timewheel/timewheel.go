package algorithm

import (
	_ "container/list"
	//"fmt"
	//"sync"
)

type TimeWheelCallBack func(val interface{})

type TimeWheelTaskData struct {
	wheel    int32
	slot     int32
	ticks    int64
	data     interface{}
	callBack TimeWheelCallBack
}

type slot struct {
	head int32
	size int32
}

type wheel struct {
	currentSlot int32
	max         int64
	slots       []slot
}

type TimeWheel struct {
	isBinaryBits bool
	size         int32
	tick         int64
	currentTicks int64
	maxTicks     int64
	wheels       []wheel
	wheelBits    []uint32
	wheelMasks   []int32
	allocator    *TimeWheelAllocator
	stat         TimeWheelStat

	//mutex  sync.Mutex
}

type TimeWheelStat struct {
	Add              uint64
	AddOk            uint64
	InternalAdd      uint64
	InternalAddOk    uint64
	Remove           uint64
	RemoveOk         uint64
	InternalRemove   uint64
	InternalRemoveOk uint64
	Step             uint64
	Expire           uint64
	ExpireBeforeAdd  uint64
	Post             uint64
	RemoveAll        uint64
	MoveWheels       uint64
	MoveSlot         uint64
}

func (this *TimeWheelStat) Clear() {
	this.Add = 0
	this.AddOk = 0
	this.InternalAdd = 0
	this.InternalAddOk = 0
	this.Remove = 0
	this.RemoveOk = 0
	this.InternalRemove = 0
	this.InternalRemoveOk = 0
	this.Step = 0
	this.Expire = 0
	this.ExpireBeforeAdd = 0
	this.Post = 0
	this.RemoveAll = 0
}

func NewTimeWheel(timeWheelNum int, slotNum []int, delta, currentTicks int64, totalDataNum int32) *TimeWheel {
	if len(slotNum) != timeWheelNum {
		return nil
	}

	if delta <= 0 {
		return nil
	}

	tw := &TimeWheel{tick: delta}
	tw.wheels = make([]wheel, timeWheelNum)
	tw.maxTicks = delta

	for i := 0; i < timeWheelNum; i++ {
		v := &tw.wheels[i]
		v.max = delta * int64(slotNum[i])
		tw.maxTicks *= int64(slotNum[i])

		for j := 0; j < slotNum[i]; j++ {
			slot := slot{head: -1}
			v.slots = append(v.slots, slot)
		}
	}

	tw.allocator = NewTimeWheelAllocator(totalDataNum)
	tw.currentTicks = currentTicks
	return tw
}

func NewTimeWheelBinaryBits(timeWheelNum int, slotNumBit []int, delta, currentTicks int64, totalDataNum int32) *TimeWheel {
	var slotNum []int

	for i := 0; i < len(slotNumBit); i++ {
		slotNum = append(slotNum, 1<<uint32(slotNumBit[i]))
	}

	tw := NewTimeWheel(timeWheelNum, slotNum, delta, currentTicks, totalDataNum)
	if tw == nil {
		return nil
	}

	tw.isBinaryBits = true

	for i := 0; i < len(slotNumBit); i++ {
		tw.wheelBits = append(tw.wheelBits, uint32(slotNumBit[i]))
		tw.wheelMasks = append(tw.wheelMasks, int32(1<<uint32(slotNumBit[i])-1))
	}

	return tw
}

func (this *TimeWheel) Size() int32 {
	return this.size
}

func (this *TimeWheel) Add(interval int64, data interface{}, callBack TimeWheelCallBack) int32 {
	this.stat.Add++

	if interval <= 0 {
		this.stat.Expire++
		this.stat.ExpireBeforeAdd++
		if callBack != nil {
			this.stat.Post++
			callBack(data)
		}
		return -2
	}

	if interval >= this.maxTicks {
		return -1
	}

	chunk := this.allocator.AllocEx()
	if chunk == nil {
		return -1
	}

	chunk.data.ticks = interval + this.currentTicks
	chunk.data.data = data
	chunk.data.callBack = callBack

	this.internalAdd(interval, chunk)

	this.size++
	this.stat.AddOk++

	return chunk.id
}

func (this *TimeWheel) internalAdd(interval int64, chunk *Chunk) {
	this.stat.InternalAdd++

	var wheelIndex int32
	var slotIndex int32

	if this.isBinaryBits {
		wheelIndex, slotIndex = this.calcPosBinaryBits(interval)
	} else {
		wheelIndex, slotIndex = this.calcPos(interval)
	}

	chunk.data.wheel = wheelIndex
	chunk.data.slot = slotIndex

	this.slotAppendChunk(&this.wheels[wheelIndex].slots[slotIndex], chunk)

	this.stat.InternalAddOk++
}

func (this *TimeWheel) slotAppendChunk(slot *slot, chunk *Chunk) {
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
}

func (this *TimeWheel) calcPos(interval int64) (wheel, slot int32) {
	wheelNum := len(this.wheels)
	for i := 0; i < wheelNum; i++ {
		size := int64(len(this.wheels[i].slots))
		if interval < this.wheels[i].max {
			return int32(i), int32((interval / this.tick) % size)
		}
		interval /= size
	}
	return -1, -1
}

func (this *TimeWheel) calcPosBinaryBits(interval int64) (wheel, slot int32) {
	wheelNum := len(this.wheels)
	for i := 0; i < wheelNum; i++ {
		if interval < this.wheels[i].max {
			return int32(i), int32(interval/this.tick) & this.wheelMasks[i]
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

	this.internalRemove(chunk)

	this.allocator.Free(id)

	this.size--
	this.stat.RemoveOk++

	return true
}

func (this *TimeWheel) internalRemove(chunk *Chunk) {
	this.stat.InternalRemove++

	slot := &this.wheels[chunk.data.wheel].slots[chunk.data.slot]

	if slot.size == 1 {
		slot.head = -1
	} else {
		this.allocator.Chunks[chunk.prev].next = chunk.next
		this.allocator.Chunks[chunk.next].prev = chunk.prev
		if slot.head == chunk.id {
			slot.head = chunk.next
		}
	}
	slot.size--

	this.stat.InternalRemoveOk++
}

func (this *TimeWheel) Step(current int64) {
	this.stat.Step++

	if current < this.currentTicks {
		return
	}

	ticks := (current - this.currentTicks) / this.tick

	for i := int64(0); i < ticks; i++ {
		this.StepOne()
	}

	this.currentTicks = current
}

func (this *TimeWheel) StepOne() {
	this.currentTicks += this.tick

	wheel := &this.wheels[0]
	slotIndex := (wheel.currentSlot + 1) % int32(len(wheel.slots))
	slot := &wheel.slots[slotIndex]
	wheel.currentSlot = slotIndex

	if slot.size > 0 {
		this.expireSlot(slot)
	}

	if slotIndex == 0 {
		this.moveWheels()
	}

}

func (this *TimeWheel) moveWheels() {
	this.stat.MoveWheels++
	num := len(this.wheels)
	for i := 1; i < num; i++ {
		wheel := &this.wheels[i]
		slotIndex := (wheel.currentSlot + 1) % int32(len(wheel.slots))
		slot := &wheel.slots[slotIndex]
		if slot.size == 0 {
			return
		}

		this.moveSlot(slot)

		if slotIndex > 0 {
			return
		}
	}
}

func (this *TimeWheel) moveSlot(slot *slot) {
	this.stat.MoveSlot++
	num := slot.size
	index := slot.head
	for i := int32(0); i < num; i++ {
		chunk := &this.allocator.Chunks[index]
		if chunk.used == 0 {
			panic("chunk should be used")
		}
		this.internalRemove(chunk)

		interval := chunk.data.ticks - this.currentTicks
		if interval <= 0 {
			this.stat.Expire++
			if chunk.data.callBack != nil {
				this.stat.Post++
				chunk.data.callBack(chunk.data)
			}
		} else {
			this.internalAdd(interval, chunk)
		}

		index = chunk.next
	}

}

func (this *TimeWheel) expireSlot(slot *slot) {
	num := slot.size
	index := slot.head
	for i := int32(0); i < num; i++ {
		chunk := &this.allocator.Chunks[index]
		if chunk.used == 0 {
			panic("chunk should be used")
		}
		this.stat.Expire++
		if chunk.data.callBack != nil {
			this.stat.Post++
			chunk.data.callBack(chunk.data.data)
		}
		index = chunk.next
		this.allocator.Free(chunk.id)
	}

	slot.head = -1
	slot.size = 0
}

func (this *TimeWheel) RemoveAll() {
	this.stat.RemoveAll++
	wheelNum := len(this.wheels)
	for i := 0; i < wheelNum; i++ {
		slotNum := len(this.wheels[i].slots)
		for j := 0; j < slotNum; j++ {
			this.wheels[i].slots[j].head = -1
			this.wheels[i].slots[j].size = -1
		}
	}
	this.allocator.FreeAll()
}
