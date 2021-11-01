// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/19/2021

package gosl

// NewBuffer will take a byte slice, use it in a buf created
// returns a pointer of a Buffer created. This is not a thread-safe.
// For thread-safe, use with gosl.Pool.
func NewBuffer(p []byte) *Buffer {
	return &Buffer{
		buf: p,
	}
}

// Buffer is a simplified version of a byte buffer
type Buffer struct {
	buf []byte
}

// Write is returning (int, err) to qualify as a io.Writer
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// WriteByte will take a byte (or bytes) and write it to the buffer
func (b *Buffer) WriteByte(a ...byte) *Buffer {
	b.buf = append(b.buf, a...)
	return b
}

// WriteString will take a string and write it to the buffer
func (b *Buffer) WriteString(s string) *Buffer {
	b.buf = append(b.buf, s...)
	return b
}

// Cap returns current buf capacity
func (b *Buffer) Cap() int {
	return cap(b.buf)
}

// Len returns current buf size
func (b *Buffer) Len() int {
	return len(b.buf)
}

// Reset will resize current buffer to 0 in size
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
}

// String returns current buffer as a string.
func (b *Buffer) String() string {
	return string(b.buf)
}

// Bytes returns buf as a byte slice.
func (b Buffer) Bytes() []byte {
	return b.buf
}

// Dump writes buffer to a writer, and reset the buffer.
func (b Buffer) Dump(w Writer) (n int, err error) {
	n, err = w.Write(b.buf)
	b.Reset()
	return n, err
}

