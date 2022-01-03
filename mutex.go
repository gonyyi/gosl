// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// WARNING:
//     Code below will be much slower than using standard libraries.
//     This is just a test to see if it can be built without
//     using any libraries (not even built-in), but only using
//     builtin functions of the language.

// *************************************************************************
// Mutex
// *************************************************************************

// Mutex with a channel is not fast, but does not require any import of library.
// To save the size, use empty struct for the channel.
// Mutex can be initialized by either:
// - `mu := make(Mutex, 1)`
// - `var mu Mutex; mu = mu.Init()`
type Mutex chan struct{}

func NewMutex() Mutex {
	return make(Mutex, 1)
}

// Init will initialize
// No need to use this is Mutex is created by NewMutex
func (m Mutex) Init() Mutex {
	return make(Mutex, 1)
}

// Lock will lock the mutex status
func (m Mutex) Lock() {
	m <- struct{}{}
}

// Unlock the mutex
func (m Mutex) Unlock() {
	<-m
}

// LockFor will take a function and start lock before running the func, and unlock right after.
// Usage: Mutex.LockFor( func(){ c+=1 } )
func (m Mutex) LockFor(f func()) {
	m.Lock()
	if f != nil {
		f()
	}
	m.Unlock()
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

// *************************************************************************
// Mutex Integer
// *************************************************************************

func NewMuInt() MuInt {
	return MuInt{
		mu: NewMutex(),
		i:  0,
	}
}

// MuInt is a simple mutex counter for runner.
type MuInt struct {
	mu Mutex
	i  int
}

// Init will initialize
func (c MuInt) Init() MuInt {
	c.mu = c.mu.Init()
	c.i = 0
	return c
}

// Get will get MutexInt value
func (c *MuInt) Get() (i int) {
	c.mu.LockFor(func() {
		i = c.i
	})
	return i
}

// Set will update MutexInt value with given i
func (c *MuInt) Set(i int) {
	c.mu.LockFor(func() {
		c.i = i
	})
}

// Add will update MutexInt value
func (c *MuInt) Add(i int) {
	c.mu.LockFor(func() {
		c.i += i
	})
}

// Wait will wait until the value became given integer i.
func (c *MuInt) Wait(i int) {
	for {
		if c.Get() == 0 {
			break
		}
	}
}
