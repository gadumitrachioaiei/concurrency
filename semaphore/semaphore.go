// Package semaphore uses a channel to make a bounded semaphore.
package semaphore

type Semaphore chan int

// New returns a new Semaphore of specified size, which signifies maximum simultaneous resources available.
func New(size int) Semaphore {
	return make(Semaphore, size)
}

// NewWithResources returns a new Semaphore of specified size, which signifies maximum simultaneous resources available.
// The semaphore will already have size resources allocated.
func NewWithResources(size int) Semaphore {
	s := New(size)
	s.addResources()
	return s
}

// Down uses a resource.
//
// It will block if there are no resources to use, until balanced by a call to Up.
func (s Semaphore) Down() {
	<-s
}

// Up ads a new resource to be used.
//
// It will block if we already have more than size resources, until released by a call to Down.
func (s Semaphore) Up() {
	s <- 0
}

// addResources initializes this semaphore with size resources.
func (s Semaphore) addResources() {
	for i := 0; i < cap(s); i++ {
		s.Up()
	}
}
