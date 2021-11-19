// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package gosl

// DoNothing will be used to prevent nil error
func DoNothing() {}

// Filesize
const (
	VERSION Ver = "Gosl v0.3.0-0"

	KB int64 = 1024
	MB       = KB * 1024
	GB       = MB * 1024
	TB       = GB * 1024
	PB       = TB * 1024
	EB       = PB * 1024

	IntType            = 32 << (^uint(0) >> 63) // 64 or 32
	internalBufferSize = 2 << 10                // init 1k, max 2k
)

// ********************************************************************************
// Internal Buffer Pool - bp
// ********************************************************************************

// bp is a byte buffer pool to be use internally for the logger, etc.
var bp = Pool{
	New: func() interface{} {
		return &bpBuffer{
			Buffer{
				Buf: make([]byte, 0, internalBufferSize),
			},
		}
	},
}

type bpBuffer struct {
	Buffer
}

// Free will return buffer to the pool
// However, if the buffer's capacity has been extended too large (2kb), drop the buffer.
func (b *bpBuffer) Free() {
	if cap(b.Buffer.Buf) > internalBufferSize {
		return
	}
	b.Buf = b.Buf[:0]
	bp.Put(b)
}

func getBuffer() *bpBuffer {
	return bp.Get().(*bpBuffer)
}
