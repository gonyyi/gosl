// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/05/2022

package gosl

// NewLogger will create a Logger
// Per https://github.com/gonyyi/gosl/issues/13
func NewLogger(w Writer) Logger {
	l := Logger{newline: true}
	l = l.SetOutput(w)
	return l
}

// Logger uses value instead of pointer (8 bytes in 64bit system)
// as size of Logger is 24 bytes (21 bytes, but alloc is 24 bytes)
type Logger struct {
	w       Writer // 16
	enable  bool   // 1
	newline bool   // this forces newline to be added for each write
}

// SetNewline will set newline value. If this is true, the Logger will
// ensure newline is being added after every `Logger.Write()` or `Logger.WriteString`
// By default, this is set to true.
func (l Logger) SetNewline(t bool) Logger {
	l.newline = t
	return l
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

// SetOutput will update the output writer of Logger
// If newline is set to true, for each write, it will evaluate and append newline if missing
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
	if l.newline {
		if len(p) > 0 && len(p)-1 == '\n' {
			return l.w.Write(p)
		}

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
