// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package gosl

// WARNING:
//     Code below will be slower than using standard libraries.
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

func (m Mutex) Init() Mutex {
	return make(Mutex, 1)
}

// Unlock the mutex
func (m Mutex) Unlock() {
	<-m
}

// Lock will lock the mutex status
func (m Mutex) Lock() {
	m <- struct{}{}
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
// Once -- Run only once
// *************************************************************************

// Once is channel that is designed to be run only once.
// Once can be initialized by `var o Once; o = o.Init()` or `o := make(Once,1)`
type Once chan struct{}

// Init takes int `n` and Once.Do() can run `n`th times..
func (o Once) Init(n int) Once {
	o = make(Once, n)
	return o
}

// Reset will fill counters again. If initialized with size 5,
// and 2 were ran, and Reset() was called, then it will have 5 again.
// (additional 3)
func (o Once) Reset() {
	for i := 0; i < cap(o); i++ {
		select {
		case <-o:
		default:
		}
	}
}

// Do will execute given function and close the channel.
// If it was ran previous, it won't do anything and will return false.
func (o Once) Do(f func()) bool {
	select {
	case o <- struct{}{}:
		f()
		return true
	default:
		// close(o)
		return false
	}
}

// Available will return how many times, it can run.
func (o Once) Available() int {
	return cap(o) - len(o)
}

// Close will close the channel
func (o Once) Close() error {
	close(o)
	return nil
}

// *************************************************************************
// MUTEX BOOL
// *************************************************************************

// MuBool is a simple mutex counter for runner.
type MuBool struct {
	mu Mutex
	t  bool
}

// Init will initialize
func (b MuBool) Init() MuBool {
	b.mu = b.mu.Init()
	b.t = false
	return b
}

// Get will pull value from the MutexBool
func (b *MuBool) Get() (t bool) {
	b.mu.LockFor(func() {
		t = b.t
	})
	return
}

// Set will set MutexBool
func (b *MuBool) Set(t bool) {
	b.mu.LockFor(func() {
		b.t = t
	})
}

// *************************************************************************
// INT COUNTER for RUNNER
// *************************************************************************

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
func (c *MuInt) Get() (n int) {
	c.mu.LockFor(func() {
		n = c.i
	})
	return
}

// Add will update MutexInt value
func (c *MuInt) Add(n int) {
	c.mu.LockFor(func() {
		c.i += n
	})
}

