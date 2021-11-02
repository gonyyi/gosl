// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

func NewMutex() Mutex {
	return make(Mutex, 1) // only 1 at a time
}

// Mutex with a channel is not fast, but does not require any import of library.
// To save the size, use empty struct for the channel.
type Mutex chan struct{}

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

