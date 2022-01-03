// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// Logger uses value instead of pointer (8 bytes in 64bit system)
// as size of Logger is 24 bytes (21 bytes, but alloc is 24 bytes)
type Logger struct {
	w       Writer // 16
	enable  bool   // 1
	newline bool   // this forces newline to be added for each write
}

// SetOutput will update the output writer of Logger
// If newline is set to true, for each write, it will evaluate and append newline if missing
func (l Logger) SetOutput(w Writer, newline bool) Logger {
	if w != nil {
		l.w, l.enable, l.newline = w, true, newline
		return l
	}
	l.w, l.enable, l.newline = nil, false, newline
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
		buf.Free()
		return
	}

	// Process s if p is nil
	if p == nil {
		buf := GetBuffer()
		buf.Buf = buf.Buf.WriteString(s)
		n, err = l.w.Write(buf.Buf)
		buf.Free()
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
