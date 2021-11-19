// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package gosl

// pool.go (gosl pool)
// This is to do what sync.Pool does, however, without importing any libraries at all (including standard library).
// Note that pooling using a channel is about 3 times slower than, sync.Pool. But, if this is correctly used,
// it will have better memory usage. If performance is more important, use sync.Pool instead.

// Pool is a struct with a channel and initialization function (New).
type Pool struct {
	init bool
	pool chan interface{}
	New  func() interface{}
}

// Init will set the pool size. If this wasn't set or invalid value was used, then default value will be used.
// (default value: 256)
func (p Pool) Init(size int) Pool {
	p.init = true
	p.pool = make(chan interface{}, size)
	return p
}

// Get will pull an item (pointer) from the pool if exists,
// otherwise, it will create a new item.
func (p *Pool) Get() interface{} {
	if p.init == false {
		*p = p.Init(256) // default to 256
	}
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

