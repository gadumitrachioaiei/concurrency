// Package boundedbuffer provides a bounded buffer.
package boundedbuffer

import (
	"sync"

	"github.com/gadumitrachioaiei/concurrency/semaphore"
)

// BoundedBuffer represents a queue as described in the package's README file.
type BoundedBuffer struct {
	size   int                 // the maximum size of the buffer
	filled semaphore.Semaphore // how many slots are actually filled
	empty  semaphore.Semaphore // how many slots are available to be filled

	mu sync.Mutex
	a  []int // the buffer data
}

// New returns a new buffer with the specified maximum capacity.
func New(size int) *BoundedBuffer {
	b := BoundedBuffer{
		size:   size,
		filled: semaphore.New(size),
		empty:  semaphore.NewWithResources(size),
	}
	return &b
}

// Produce adds an item at the end of the queue and blocks when capacity is reached.
func (b *BoundedBuffer) Produce(item int) {
	b.empty.Down()
	b.mu.Lock()
	b.a = append(b.a, item)
	b.mu.Unlock()
	b.filled.Up()
}

// Consume removes and returns the item from the beginning of the queue, and blocks when there is no more data.
func (b *BoundedBuffer) Consume() int {
	var item int
	b.filled.Down()
	b.mu.Lock()
	item = b.a[0]
	b.a = b.a[1:]
	b.mu.Unlock()
	b.empty.Up()
	return item
}
