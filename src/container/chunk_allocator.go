package container

type Chunk struct {
	id   int32
	used int32
	prev int32
	next int32
	data interface{}
}

type ChunkAllocator struct {
	capacity int32
	size     int32
	freeHead int32
	//busyHead int32
	Chunks []Chunk
}

func NewChunkAllocator(capacity int32) *ChunkAllocator {
	allocator := &ChunkAllocator{capacity: capacity}
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
	//allocator.busyHead = -1

	return allocator
}

func (this *ChunkAllocator) GetData(id int32) (interface{}, bool) {
	if id < 0 || int(id) > len(this.Chunks) {
		return nil, false
	}

	chunk := &this.Chunks[id]
	if chunk.used == 0 {
		return nil, false
	}

	return chunk.data, true
}

func (this *ChunkAllocator) SetData(id int32, data interface{}) bool {
	if id < 0 || int(id) > len(this.Chunks) {
		return false
	}

	chunk := &this.Chunks[id]
	if chunk.used == 0 {
		return false
	}

	chunk.data = data

	return true
}

func (this *ChunkAllocator) Alloc() int32 {
	chunk := this.AllocEx()
	if chunk == nil {
		return -1
	}
	return chunk.id
}

func (this *ChunkAllocator) AllocEx() *Chunk {
	if this.size >= this.capacity {
		return nil
	}

	// pop_front from free list
	freeNum := this.capacity - this.size
	chunk := &this.Chunks[this.freeHead]

	if freeNum == 1 {
		this.freeHead = -1
	} else {
		this.Chunks[chunk.prev].next = chunk.next
		this.Chunks[chunk.next].prev = chunk.prev
		this.freeHead = chunk.next
	}

	// push_back to busy list
	/*if this.busyHead == -1 {
		chunk.next = chunk.id
		chunk.prev = chunk.id
		this.busyHead = chunk.id
	} else {
		head := &this.chunks[this.busyHead]
		tail := &this.chunks[head.prev]
		chunk.next = this.busyHead
		chunk.prev = tail.id
		tail.next = chunk.id
		head.prev = chunk.id
	}*/

	chunk.used = 1

	this.size++
	return chunk
}

func (this *ChunkAllocator) Free(id int32) {
	if id < 0 || int(id) > len(this.Chunks) {
		return
	}

	chunk := &this.Chunks[id]
	if chunk.used == 0 {
		return
	}

	// remove from busy list
	/*if this.size == 1 {
		this.busyHead = -1
	} else {
		this.chunks[chunk.prev].next = chunk.next
		this.chunks[chunk.next].prev = chunk.prev
		if this.busyHead == chunk.id {
			this.busyHead = chunk.next
		}
	}*/

	// push_back to free list
	if this.freeHead == -1 {
		chunk.next = chunk.id
		chunk.prev = chunk.id
		this.freeHead = chunk.id
	} else {
		head := &this.Chunks[this.freeHead]
		tail := &this.Chunks[head.prev]
		chunk.next = this.freeHead
		chunk.prev = tail.id
		tail.next = chunk.id
		head.prev = chunk.id
	}

	chunk.used = 0
	this.size--
}

func (this *ChunkAllocator) FreeEx(chunk *Chunk) {
	if chunk.used == 0 {
		return
	}

	// remove from busy list
	/*if this.size == 1 {
		this.busyHead = -1
	} else {
		this.chunks[chunk.prev].next = chunk.next
		this.chunks[chunk.next].prev = chunk.prev
		if this.busyHead == chunk.id {
			this.busyHead = chunk.next
		}
	}*/

	// push_back to free list
	if this.freeHead == -1 {
		chunk.next = chunk.id
		chunk.prev = chunk.id
		this.freeHead = chunk.id
	} else {
		head := &this.Chunks[this.freeHead]
		tail := &this.Chunks[head.prev]
		chunk.next = this.freeHead
		chunk.prev = tail.id
		tail.next = chunk.id
		head.prev = chunk.id
	}

	chunk.used = 0
	this.size--
}
