// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// bufPool is a byte Buf pool to be use internally for the logger, etc.
var (
	// bufPoolCreated = NewMuInt()
	bufPool = Pool{
		New: func() interface{} {
			// bufPoolCreated.Add(1)
			return &bufPoolItem{
				// id: bufPoolCreated.Get(),
				Buf: make([]byte, 0, GlobalBufferSize),
			}
		},
	}.Init(256)
)

type bufPoolItem struct {
	// id     int
	Buf Buf
}

// Free will return Buf to the pool
// However, if the Buf's capacity has been extended too large (2kb), drop the Buf.
func (b *bufPoolItem) Free() {
	if cap(b.Buf) > GlobalBufferSize {
		return
	}
	// b.Buf = b.Buf[:0]
	bufPool.Put(b)
}

// GetBuffer returns global buffer from the pool
func GetBuffer() *bufPoolItem {
	buf := bufPool.Get().(*bufPoolItem)
	buf.Buf = buf.Buf[:0]
	return buf
}
