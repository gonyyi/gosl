// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// NewBuffer will take a byte slice, use it in a Buf created
// returns a pointer of a Buffer created. This is not a thread-safe.
// For thread-safe, use with gosl.Pool.
func NewBuffer(p []byte) *Buffer {
	return &Buffer{
		Buf: p,
	}
}

// Buffer is a simplified version of a byte buffer
type Buffer struct {
	Buf []byte
}

// Write is returning (int, err) to qualify as a io.Writer
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Buf = append(b.Buf, p...)
	return len(p), nil
}

// WriteByte will take a byte (or bytes) and write it to the buffer
func (b *Buffer) WriteByte(a ...byte) *Buffer {
	b.Buf = append(b.Buf, a...)
	return b
}

// WriteString will take a string and write it to the buffer
func (b *Buffer) WriteString(s string) *Buffer {
	b.Buf = append(b.Buf, s...)
	return b
}

// LastByte returns late byte in the buffer
// If buffer is empty, this will return 0.
func (b *Buffer) LastByte() byte {
	if i := b.Len(); i > 0 {
		return b.Buf[i-1]
	}
	return 0
}

// Trim will cut n bytes from the end of buffer
// If larger number is given than the size, then just empty the buffer
func (b *Buffer) Trim(n uint) {
	if i := b.Len() - int(n); i > -1 {
		b.Buf = b.Buf[:i]
		return
	}
	b.Buf = b.Buf[:0]
}

// Cap returns current Buf capacity
func (b *Buffer) Cap() int {
	return cap(b.Buf)
}

// Len returns current Buf size
func (b *Buffer) Len() int {
	return len(b.Buf)
}

// Reset will resize current buffer to 0 in size
func (b *Buffer) Reset() {
	b.Buf = b.Buf[:0]
}

// String returns current buffer as a string.
func (b *Buffer) String() string {
	return string(b.Buf)
}

// Bytes returns Buf as a byte slice.
func (b Buffer) Bytes() []byte {
	return b.Buf
}

// Dump writes buffer to a writer, and reset the buffer.
func (b Buffer) Dump(w Writer) (n int, err error) {
	n, err = w.Write(b.Buf)
	b.Reset()
	return n, err
}

