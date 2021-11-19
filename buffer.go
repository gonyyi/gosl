// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package gosl

// TODO: Do I this need Buffer here? Maybe just enable gosl.internalBuffer??

// NewBuffer will take a byte slice, use it in a Buf created
// returns a pointer of a Buffer created. This is not a thread-safe.
// For thread-safe, use with gosl.Pool.
func NewBuffer(p []byte) *Buffer {
	return &Buffer{
		Buf: p,
	}
}

// Buffer is a simplified version of a byte Buffer
type Buffer struct {
	Buf Buf
}

func (b *Buffer) Init() {
	b.Buf = make(Buf, 0, 1024)
}

// Write is returning (int, err) to qualify as a io.Writer
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Buf = b.Buf.WriteByte(p...)
	return len(p), nil
}

// WriteByte will take a byte (or bytes) and write it to the Buffer
func (b *Buffer) WriteByte(a ...byte) *Buffer {
	b.Buf = b.Buf.WriteByte(a...)
	return b
}

// WriteBool will take a bool value and add it to buffer as a string
func (b *Buffer) WriteBool(t bool) *Buffer {
	b.Buf = b.Buf.WriteBool(t)
	return b
}

// WriteInt will take integer and add it to buffer as a string
func (b *Buffer) WriteInt(i int) *Buffer {
	b.Buf = b.Buf.WriteInt(i)
	return b
}

// WriteFloat64 will take float64 and add it to buffer as a string
func (b *Buffer) WriteFloat64(f64 float64) *Buffer {
	b.Buf = b.Buf.WriteFloat64(f64)
	return b
}

// WriteString will take a string and write it to the Buffer
func (b *Buffer) WriteString(s string) *Buffer {
	b.Buf = b.Buf.WriteString(s)
	return b
}

// Last returns late byte in the Buffer
// If Buffer is empty, this will return 0.
func (b *Buffer) Last() byte {
	return b.Buf.Last()
}

// Trims will cut n bytes from the end of Buffer
// If larger number is given than the size, then just empty the Buffer
func (b *Buffer) Trim(n uint) {
	b.Buf = b.Buf.Trim(n)
}

// Cap returns current Buf capacity
func (b *Buffer) Cap() int {
	return b.Buf.Cap()
}

// Len returns current Buf size
func (b *Buffer) Len() int {
	return b.Buf.Len()
}

// Reset will resize current Buffer to 0 in size
func (b *Buffer) Reset() {
	b.Buf = b.Buf.Reset()
}

// String returns current Buffer as a string.
func (b *Buffer) String() string {
	return b.Buf.String()
}

// Bytes returns Buf as a byte slice.
func (b *Buffer) Bytes() []byte {
	return b.Buf
}

// WriteTo writes Buffer to a writer, and reset the Buffer.
func (b *Buffer) WriteTo(w Writer) (n int, err error) {
	return b.Buf.WriteTo(w)
}

