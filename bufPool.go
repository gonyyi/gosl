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

// **********************************************************************
// ADDITIONAL
// **********************************************************************

// Bytes returns []byte of current buffer
func (b *bufPoolItem) Bytes() []byte {
	return b.Buf.Bytes()
}

// Cap returns current capacity
func (b *bufPoolItem) Cap() int {
	return b.Buf.Cap()
}

// Len returns current length
func (b *bufPoolItem) Len() int {
	return b.Buf.Len()
}

// Println prints current buffer
func (b *bufPoolItem) Println() {
	b.Buf.Println()
}

// Reset resets current buffer
func (b *bufPoolItem) Reset() {
	b.Buf = b.Buf.Reset()
}

// Set resets and set current buffer with a given string
func (b *bufPoolItem) Set(s string) {
	b.Buf = b.Buf.Set(s)
}

// String returns current buffer in string format
func (b *bufPoolItem) String() string {
	return b.Buf.String()
}

// Write writes bytes into current buffer
// This meets Writer interface.
func (b *bufPoolItem) Write(p []byte) (n int, err error) {
	return b.Buf.Write(p)
}

// WriteBytes will write byte or bytes to current buffer
func (b *bufPoolItem) WriteBytes(bytes ...byte) {
	b.Buf = b.Buf.WriteBytes(bytes...)
}

// WriteBool will write boolean t to current buffer in string
// true => "true", false => "false"
func (b *bufPoolItem) WriteBool(t bool) {
	b.Buf = b.Buf.WriteBool(t)
}

// WriteFloat will write float f with decimal point dec to current buffer
func (b *bufPoolItem) WriteFloat(f float64, dec uint8) {
	b.Buf = b.Buf.WriteFloat(f, dec)
}

// WriteInt will write integer i to buffer as string format
func (b *bufPoolItem) WriteInt(i int) {
	b.Buf = b.Buf.WriteInt(i)
}

// WriteString will write a string s to current buffer
func (b *bufPoolItem) WriteString(s string) {
	b.Buf = b.Buf.WriteString(s)
}

// WriteStrings will write string slice to current buffer with a delimiter
func (b *bufPoolItem) WriteStrings(s []string, delim ...byte) {
	b.Buf = b.Buf.WriteStrings(s, delim...)
}

// WriteTo will write current buffer to writer w.
func (b *bufPoolItem) WriteTo(w Writer) (n int, err error) {
	return b.Buf.WriteTo(w)
}
