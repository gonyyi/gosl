// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

func NewMutex() Mutex {
	return make(Mutex, 1) // only 1 at a time
}

// Mutex with a channel is not fast, but does not require any import of library.
// To save the size, use empty struct for the channel.
type Mutex chan struct{}

// LockFor will take a function and start lock before running the func, and unlock right after.
// Usage: Mutex.LockFor( func(){ c+=1 } )
func (m Mutex) LockFor(f func()) {
	m.Lock()
	defer m.Unlock()
	if f != nil {
		f()
	}
}

// Lock will lock the mutex status
func (m Mutex) Lock() {
	m <- struct{}{}
}

// LockIfNot will obtain mutex and return true if not locked. Otherwise, it will return false.
func (m Mutex) LockIfNot() (ok bool) {
	select {
	case m <- struct{}{}:
		return true
	default:
		return false
	}
}

// Locked will return true if mutex is locked.
func (m Mutex) Locked() bool {
	return len(m) == 1
}

// Unlock the mutex
func (m Mutex) Unlock() {
	<-m
}

// NewOnce creates a channel for the function that can run only run once.
func NewOnce() Once {
	o := make(Once, 1)
	for i := 0; i < 1; i++ {
		o <- struct{}{}
	}
	return o
}

// Once is channel that is designed to be run only once.
type Once chan struct{}

// Do will execute given function and close the channel.
// If it was ran previous, it won't do anything and will return false.
func (o Once) Do(f func()) bool {
	if _, ok := <-o; ok {
		// channel exist
		f()
		close(o)
		return true
	}
	return false
}

