// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// pool.go (gosl pool)
// This is to do what sync.Pool does, however, without importing any libraries at all (including standard library).
// Note that pooling using a channel is about 3 times slower than, sync.Pool. But, if this is correctly used,
// it will have better memory usage. If performance is more important, use sync.Pool instead.
// - NewPool(int, func()interface{}) *pool
// - *pool.Get() interface{}
// - *pool.Put(interface{})

// NewPool will create
func NewPool(Max int, New func() interface{}) (cp pool) {
	return pool{
		pool: make(chan interface{}, Max),
		New:  New,
	}
}

// pool is a struct with a channel and initialization function (New).
type pool struct {
	pool chan interface{}
	New  func() interface{}
}

// Get will pull an item (pointer) from the pool if exists,
// otherwise, it will create a new item.
func (p *pool) Get() (b interface{}) {
	select {
	case b = <-p.pool: // Reuse
	default:
		// Item not exists --> Create new
		b = p.New()
	}
	return
}

// Put will take an item (pointer) and put it back to the pool.
// If the pool is full, it will not put it the item back (discard).
func (p *pool) Put(b interface{}) {
	select {
	case p.pool <- b: // PUT BUFFER BACK
	default: // DISCARD BUF, POOL IS FULL
	}
}

