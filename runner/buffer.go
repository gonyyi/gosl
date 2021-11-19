// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package runner

import "github.com/gonyyi/gosl"

// ********************************************************************************
// Internal Buffer Pool - bp
// ********************************************************************************

const (
	internalBufferSize = 1 << 10 // init 1k, max 2k
)

// bp is a byte buffer pool to be use internally for the logger, etc.
var bp = gosl.Pool{
	New: func() interface{} {
		return &bpBuffer{
			gosl.Buffer{
				Buf: make([]byte, 0, internalBufferSize),
			},
		}
	},
}

type bpBuffer struct {
	gosl.Buffer
}

// Free will return buffer to the pool
// However, if the buffer's capacity has been extended too large (2kb), drop the buffer.
func (b *bpBuffer) Free() {
	if cap(b.Buffer.Buf) > (2 << 10) {
		return
	}
	b.Buf = b.Buf[:0]
	bp.Put(b)
}

func getBuffer() *bpBuffer {
	return bp.Get().(*bpBuffer)
}

