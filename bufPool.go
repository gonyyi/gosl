// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package gosl

// ********************************************************************************
// Internal Buffer Pool - bufferPool
// ********************************************************************************

// bufferPool is a byte Buffer pool to be use internally for the logger, etc.
var bufferPool = Pool{
	New: func() interface{} {
		return &poolBuf{
			Buffer: make([]byte, 0, DefaultBufferSize),
		}
	},
}.Init(1024)

type poolBuf struct {
	Buffer Buf
}

// Free will return Buffer to the pool
// However, if the Buffer's capacity has been extended too large (2kb), drop the Buffer.
func (b *poolBuf) Free() {
	if cap(b.Buffer) > DefaultBufferSize {
		return
	}
	b.Buffer = b.Buffer[:0]
	bufferPool.Put(b)
}

func GetBuffer() *poolBuf {
	return bufferPool.Get().(*poolBuf)
}

func (b *poolBuf) Init(size int) {
	b.Buffer = make(Buf, 0, size)
}

// Write is returning (int, err) to qualify as a io.Writer
func (b *poolBuf) Write(p []byte) (n int, err error) {
	b.Buffer = b.Buffer.WriteBytes(p...)
	return len(p), nil
}

// WriteBytes will take a byte (or bytes) and write it to the Buf
func (b *poolBuf) WriteBytes(a ...byte) *poolBuf {
	b.Buffer = b.Buffer.WriteBytes(a...)
	return b
}

// WriteBool will take a bool value and add it to Buf as a string
func (b *poolBuf) WriteBool(t bool) *poolBuf {
	b.Buffer = b.Buffer.WriteBool(t)
	return b
}

// WriteInt will take integer and add it to Buf as a string
func (b *poolBuf) WriteInt(i int) *poolBuf {
	b.Buffer = b.Buffer.WriteInt(i)
	return b
}

// WriteFloat64 will take float64 and add it to Buf as a string
func (b *poolBuf) WriteFloat64(f64 float64) *poolBuf {
	b.Buffer = b.Buffer.WriteFloat64(f64)
	return b
}

// WriteString will take a string and write it to the Buf
func (b *poolBuf) WriteString(s string) *poolBuf {
	b.Buffer = b.Buffer.WriteString(s)
	return b
}

// Last returns late byte in the Buf
// If Buf is empty, this will return 0.
func (b *poolBuf) Last() byte {
	return b.Buffer.Last()
}

// Trim will cut n bytes from the end of Buf
// If larger number is given than the size, then just empty the Buf
func (b *poolBuf) Trim(n uint) {
	b.Buffer = b.Buffer.Trim(n)
}

// Cap returns current Buf capacity
func (b *poolBuf) Cap() int {
	return b.Buffer.Cap()
}

// Len returns current Buf size
func (b *poolBuf) Len() int {
	return b.Buffer.Len()
}

// Reset will resize current Buf to 0 in size
func (b *poolBuf) Reset() {
	b.Buffer = b.Buffer.Reset()
}

// String returns current Buf as a string.
func (b *poolBuf) String() string {
	return b.Buffer.String()
}

// Bytes returns Buf as a byte slice.
func (b *poolBuf) Bytes() []byte {
	return b.Buffer
}

// WriteTo writes Buf to a writer, and reset the Buf.
func (b *poolBuf) WriteTo(w Writer) (n int, err error) {
	return b.Buffer.WriteTo(w)
}
