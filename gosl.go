// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/26/2021

package gosl

// DoNothing will be used to prevent nil error
func DoNothing() {}

// Filesize
const (
	KB int64 = 1024
	MB       = KB * 1024
	GB       = MB * 1024
	TB       = GB * 1024
	PB       = TB * 1024
	EB       = PB * 1024
)

// ********************************************************************************
// Internal Buffer Pool - bufp
// ********************************************************************************

const (
	internalBufferSize     = 1024
	internalBufferPoolSize = 256
)

// bufp is a byte buffer pool to be use internally for the logger, etc.
var bufp = NewPool(internalBufferPoolSize, func() interface{} {
	return &bufpBuffer{
		Buffer{
			buf: make([]byte, 0, internalBufferSize),
		},
	}
})

type bufpBuffer struct {
	Buffer
}

func (b *bufpBuffer) ReturnBuffer() {
	if cap(b.Buffer.buf) > (64 << 10) {
		return
	}
	b.buf = b.buf[:0]
	bufp.Put(b)
}

func getBufpBuffer() *bufpBuffer {
	return bufp.Get().(*bufpBuffer)
}
