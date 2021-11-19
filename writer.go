// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// ********************************************************************************
// Interfaces
// ********************************************************************************

// Writer - to avoid importing "io"
type Writer interface {
	// Write takes bytes and returns number of bytes and error
	// Since not all writers has Close() method, method Close() isn't
	// required for gosl.Writer
	Write(p []byte) (n int, err error)
}

// Closer is an interface for the writers that have Close method.
type Closer interface {
	Close() error
}

// ********************************************************************************
// Discard Writer
// ********************************************************************************

// discard for when nil is given as io.Writer
type discard struct{}

// (discard) Write - to satisfy Writer interface
func (discard) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// Discard is a discard writer object, so it can be used right away.
var Discard = &discard{}

// ********************************************************************************
// Function
// ********************************************************************************

// Close will take a writer or anything that has Close() method and close it if applicable.
// Originally this was just for Writer object, but extended to all by using an interface{}.
func Close(w interface{}) error {
	if c, ok := w.(Closer); ok {
		return c.Close()
	}
	return nil
}

// ********************************************************************************
// CUSTOM WRITERS: ALTER WRITER
// ********************************************************************************

// NewAlterWriter creates a writer that writes to two Writers.
// This will be used to create another writer with an AlterFn
// (See NewPrefixWriter() below)
func NewAlterWriter(dst Writer, f func([]byte) []byte) Writer {
	if f == nil {
		f = func(p []byte) []byte { return p }
	}
	if dst == nil {
		dst = Discard
	}
	return &alterWriter{
		w:     dst,
		alter: f,
	}
}

// alterWriter can modify anything that comes to the writer
type alterWriter struct {
	w     Writer
	alter func([]byte) []byte
}

// Write to meet io.Writer interface requirement
func (a alterWriter) Write(b []byte) (n int, err error) {
	if b = a.alter(b); b != nil {
		return a.w.Write(b)
	}
	return 0, nil
}

// Close will close the writer if applicable
func (a alterWriter) Close() error {
	return Close(a.w)
}

// ********************************************************************************
// CUSTOM WRITERS -- These are created using alter writer
// ********************************************************************************

// NewPrefixWriter will take a prefix and a writer
// and creates a writer that will append prefix
func NewPrefixWriter(prefix string, w Writer) Writer {
	pfx := []byte(prefix)
	return &alterWriter{
		w: w,
		alter: func(p []byte) []byte {
			return append(pfx, p...)
		},
	}
}

// NewMultiWriter will take writers and writes to those.
// When closed this will close only for the first writer.
func NewMultiWriter(w ...Writer) Writer {
	switch len(w) {
	case 0:
		return nil
	case 1:
		return w[0]
	}

	return &alterWriter{
		w: w[0], // first will be the main writer
		alter: func(p []byte) []byte {
			for _, v := range w[1:] {
				v.Write(p)
			}
			return p
		},
	}
}

