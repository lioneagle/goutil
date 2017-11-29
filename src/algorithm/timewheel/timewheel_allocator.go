package algorithm

type Chunk struct {
	id   int32
	used int32
	prev int32
	next int32
	data TimeWheelTaskData
}

type TimeWheelAllocator struct {
	capacity int32
	size     int32
	freeHead int32
	Chunks   []Chunk
}

func NewTimeWheelAllocator(capacity int32) *TimeWheelAllocator {
	allocator := &TimeWheelAllocator{capacity: capacity}
	allocator.Chunks = make([]Chunk, capacity)

	allocator.Chunks[0].id = 0
	allocator.Chunks[0].used = 0
	allocator.Chunks[0].prev = capacity - 1
	allocator.Chunks[0].next = 1

	for i := int32(1); i < capacity; i++ {
		allocator.Chunks[i].id = i
		allocator.Chunks[i].prev = i - 1
		allocator.Chunks[i].next = i + 1
	}

	allocator.Chunks[capacity-1].next = 0

	return allocator
}

func (this *TimeWheelAllocator) Alloc() int32 {
	chunk := this.AllocEx()
	if chunk == nil {
		return -1
	}
	return chunk.id
}

func (this *TimeWheelAllocator) AllocEx() *Chunk {
	if this.size >= this.capacity {
		return nil
	}

	freeNum := this.capacity - this.size
	chunk := &this.Chunks[this.freeHead]

	if freeNum == 1 {
		this.freeHead = -1
	} else {
		this.Chunks[chunk.prev].next = chunk.next
		this.Chunks[chunk.next].prev = chunk.prev
		this.freeHead = chunk.next
	}

	chunk.used = 1

	this.size++
	return chunk
}

func (this *TimeWheelAllocator) Free(id int32) {
	if id < 0 || int(id) > len(this.Chunks) {
		return
	}

	chunk := &this.Chunks[id]
	if chunk.used == 0 {
		return
	}

	// push_back to free list
	if this.freeHead == -1 {
		chunk.next = id
		chunk.prev = id
		this.freeHead = id
	} else {
		head := &this.Chunks[this.freeHead]
		tail := &this.Chunks[head.prev]
		chunk.next = this.freeHead
		chunk.prev = tail.id
		tail.next = id
		head.prev = id
	}

	chunk.used = 0
	this.size--
}
