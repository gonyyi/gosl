// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// Buf is similar to Buffer (gosl.Buffer), but without a pointer usage.
// Since Buf does not allocate when created, it will be good to use when created and not shared.
// If shared Buffer is needed. Use Buffer instead.
// Buf doesn't have a constructor to avoid allocation, and need to initialized with make like make(Buf, 512).
// Usage:  Buf := make(gosl.Buf, 0, 1024)
// Note:   Buf is slower than Buffer due to creation of Buf.
type Buf []byte

// WriteBytes will take a byte or more and write it to a Buffer
func (b Buf) WriteBytes(bytes ...byte) Buf {
	return append(b, bytes...)
}

// WriteBool will take a bool value and add it to Buffer as a string
func (b Buf) WriteBool(t bool) Buf {
	return AppendBool(b, t)
}

// WriteInt will take integer and add it to Buffer as a string
func (b Buf) WriteInt(i int) Buf {
	return AppendInt(b, i, false)
}

// WriteFloat64 will take float64 and add it to Buffer as a string
func (b Buf) WriteFloat64(f64 float64) Buf {
	return AppendFloat64(b, f64, 2, false)
}

// WriteString will take a string and write it to the Buffer
func (b Buf) WriteString(s string) Buf {
	return append(b, s...)
}

// Last will return last byte of Buffer.
// If it was not exist, it will return byte(0).
func (b Buf) Last() byte {
	if i := len(b); i > 0 {
		return b[i-1]
	}
	return 0
}

// Trim will take n (int) and remove last n bytes from the Buffer
func (b Buf) Trim(n uint) Buf {
	if i := len(b) - int(n); i > -1 {
		return b[:i]
	}
	return b[:0]
}

// Cap will return Buffer capacity
func (b Buf) Cap() int {
	return cap(b)
}

// Len will return Buffer size
func (b Buf) Len() int {
	return len(b)
}

// Reset will clear the Buffer
func (b Buf) Reset() Buf {
	return b[:0]
}

// String will return Buffer content in string
func (b Buf) String() string {
	return string(b)
}

// WriteTo will take a writer and dump Buffer contents into it.
// After that Buffer will be reset..
func (b Buf) WriteTo(w Writer) (n int, err error) {
	n, err = w.Write(b)
	b.Reset()
	return
}

// Write here is for io.Writer interface.
// And this one uses pointer receiver.
func (b *Buf) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(p), nil
}

// Println will print current buffer into stdout
func (b Buf) Println() {
	println(b.String())
}