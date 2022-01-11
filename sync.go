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


// pool.go (gosl pool)
// This is to do what sync.Pool does, however, without importing any libraries at all (including standard library).
// Note that pooling using a channel is about 3 times slower than, sync.Pool. But, if this is correctly used,
// it will have better memory usage. If performance is more important, use sync.Pool instead.

// Pool is a struct with a channel and initialization function (New).
type Pool struct {
	pool chan interface{}
	New  func() interface{}
}

// Init will set the pool size. If this wasn't set or invalid value was used, then default value will be used.
// (default value: 256)
func (p Pool) Init(size int) Pool {
	p.pool = make(chan interface{}, size)
	return p
}

// Get will pull an item (pointer) from the pool if exists,
// otherwise, it will create a new item.
func (p *Pool) Get() interface{} {
	select {
	case b := <-p.pool: // Reuse
		return b
	default:
		// Item not exists --> Create new
		if p.New == nil {
			return nil
		}
		return p.New()
	}
}

// Put will take an item (pointer) and put it back to the pool.
// If the pool is full, it will not put it the item back (discard).
func (p *Pool) Put(b interface{}) {
	select {
	case p.pool <- b: // PUT BUFFER BACK
	default: // DISCARD BUF, POOL IS FULL
	}
}
