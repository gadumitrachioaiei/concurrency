package boundedbuffer

import (
	"testing"
)

func TestNewBoundedBuffer(t *testing.T) {
	size := 10
	b := New(size)
	for i := 0; i < size; i++ {
		b.Produce(i)
	}
	go b.Produce(size)
	for i := 0; i < size+1; i++ {
		if v := b.Consume(); v != i {
			t.Fatalf("got value: %d, expected: %d", v, i)
		}
	}
}
