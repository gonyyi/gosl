// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/14/2022

package gosl

// NewLogger will create a Logger
// Per https://github.com/gonyyi/gosl/issues/13
func NewLogger(w Writer) Logger {
	l := Logger{}
	l = l.SetOutput(w)
	return l
}

// Logger uses value instead of pointer (8 bytes in 64bit system)
// as size of Logger is 24 bytes (21 bytes, but alloc is 24 bytes)
type Logger struct {
	w              Writer // 16
	enable         bool   // 1
	disableNewline bool   // if true, logger won't enforce newline.
}

// Enable will enable/disable logging
// Per https://github.com/gonyyi/gosl/issues/14
func (l Logger) Enable(t bool) Logger {
	// if writer is not set, no need to set enable as it will be always disabled,
	// also without writer set, it shouldn't be enabled.
	if l.w != nil {
		l.enable = t
	}
	return l
}

// Enabled will return if logger is currently enabled
func (l Logger) Enabled() bool {
	return l.enable
}

// SetNewline will add (enforce) newline at the end if doesn't have one.
func (l Logger) SetNewline(t bool) Logger {
	l.disableNewline = t == false
	return l
}

// SetOutput will update the output writer of Logger
func (l Logger) SetOutput(w Writer) Logger {
	if w != nil {
		l.w, l.enable = w, true
		return l
	}
	l.w, l.enable = nil, false
	return l
}

// Write takes bytes and returns number of bytes written and error
func (l Logger) Write(p []byte) (n int, err error) {
	if l.enable == true {
		return l.write(p, "")
	}
	return 0, nil
}

// WriteString takes string and returns number of bytes written and error
func (l Logger) WriteString(s string) (n int, err error) {
	if l.enable == true {
		return l.write(nil, s)
	}
	return 0, nil
}

// write writes bytes to buffer. This is separated from WriteString
// to speed up the process.
func (l Logger) write(p []byte, s string) (n int, err error) {
	if l.disableNewline {
		// Process s if p is nil
		if p == nil {
			buf := GetBuffer()
			buf.Buf = buf.Buf.WriteString(s)
			n, err = l.w.Write(buf.Buf)
			PutBuffer(buf)
			return n, err
		}
		// Process p
		return l.w.Write(p)
	}
	if len(p) > 0 && len(p)-1 == '\n' {
		return l.w.Write(p)
	}

	// APPEND NEWLINE IF NOT

	// At this point, need a buffer
	buf := GetBuffer()
	// Process p
	if p != nil { // either len(p) == 0 or len(p)-1 != '\n', so add '\n'
		buf.Buf = append(buf.Buf, p...)
	} else {
		buf.Buf = append(buf.Buf, s...)
	}

	if buf.Buf.LastByte() != '\n' {
		buf.Buf = buf.Buf.WriteBytes('\n')
	}

	n, err = l.w.Write(buf.Buf)
	PutBuffer(buf)
	return
}

// Output returns current output
func (l Logger) Output() Writer {
	return l.w
}

// Close will close writer if applicable
func (l Logger) Close() error {
	return Close(l.w)
}

// ********************************************************************************
// Interfaces
// ********************************************************************************

// Reader - to avoid importing "io"
type Reader interface {
	// Read reads a data to (fixed size) byte p and return how many bytes were read.
	// This will return EOF error when there's no more data to read.
	Read(p []byte) (n int, err error)
}

// Writer - to avoid importing "io"
type Writer interface {
	// Write takes bytes and returns number of bytes and error
	// Since not all writers has Close() method, method Close() isn't
	// required for gosl.Writer
	Write(p []byte) (n int, err error)
}

// LogWriter is an interface for the Logger
type LogWriter interface {
	Write(p []byte) (n int, err error)
	WriteString(s string) (n int, err error)
	Close() error
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
