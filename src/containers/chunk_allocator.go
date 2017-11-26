package containers

type Chunk struct {
	id   int32
	used int32
	prev int32
	next int32
	Data interface{}
}

type ChunkAllocator struct {
	capacity int32
	size     int32
	freeHead int32
	busyHead int32
	chunks   []Chunk
}

func NewChunkAllocator(capacity int32) *ChunkAllocator {
	allocator := &ChunkAllocator{capacity: capacity}
	allocator.chunks = make([]Chunk, capacity)

	allocator.chunks[0].id = 0
	allocator.chunks[0].used = 0
	allocator.chunks[0].prev = capacity - 1
	allocator.chunks[0].next = 1

	for i := int32(1); i < capacity; i++ {
		allocator.chunks[i].id = i
		allocator.chunks[i].prev = i - 1
		allocator.chunks[i].next = i + 1
	}

	allocator.chunks[capacity-1].next = 0
	allocator.busyHead = -1

	return allocator
}

func (this *ChunkAllocator) Alloc() *Chunk {
	if this.size >= this.capacity {
		return nil
	}

	// pop_front from free list
	freeNum := this.capacity - this.size
	chunk := &this.chunks[this.freeHead]

	if freeNum == 1 {
		this.freeHead = -1
	} else {
		this.chunks[chunk.prev].next = chunk.next
		this.chunks[chunk.next].prev = chunk.prev
		this.freeHead = chunk.next
	}

	// push_back to busy list
	if this.busyHead == -1 {
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
	}

	chunk.used = 1

	this.size++
	return chunk
}

func (this *ChunkAllocator) Free(chunk *Chunk) {
	if chunk.used == 0 {
		return
	}

	// remove from busy list
	if this.size == 1 {
		this.busyHead = -1
	} else {
		this.chunks[chunk.prev].next = chunk.next
		this.chunks[chunk.next].prev = chunk.prev
		if this.busyHead == chunk.id {
			this.busyHead = chunk.next
		}
	}

	// push_back to free list
	if this.freeHead == -1 {
		chunk.next = chunk.id
		chunk.prev = chunk.id
		this.freeHead = chunk.id
	} else {
		head := &this.chunks[this.freeHead]
		tail := &this.chunks[head.prev]
		chunk.next = this.freeHead
		chunk.prev = tail.id
		tail.next = chunk.id
		head.prev = chunk.id
	}

	chunk.used = 0
	this.size--
}
