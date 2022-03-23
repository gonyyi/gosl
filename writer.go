// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 02/01/2022

package gosl

// ********************************************************************************
// LvWriter is a writer with level supports. Level can be any uint8 type, but this
// has LvTrace ... LvFatal are predefined for convenience.
//
// Eg. Using Predefined Levels:
//     w := NewLvWriter( os.Stdout, LvInfo ) // Set lvMin level as LvInfo
//     w.Debug().WriteString("blah bug blah.......") // will NOT be written
//     w.Info().WriteString("just fyi blah blah...") // will be written
//
// Eg. Using Custom Levels:
//     var (
//         INFO uint8 = 0    // Make sure levels are uint8 (0-255 level)
//         WARN uint8 = 3
//         ERR  uint8 = 7
//         // More...
//     )
//     w := NewLvWriter( os.Stdout, WARN ) // Writer will print for WARN or above
//     w.Lv(INFO).WriteString("blah info fyi...")  // will NOT be written
//     w.Lv(WARN).WriteString("HMM.............")  // will be written
//     w.Lv(ERR ).WriteString("blah ERROR!!!!!!")  // will be written
//     w = w.SetLevel(ERR)  // change the level to ERR
//     w.Lv(WARN).WriteString("HUH.............")  // now, this will NOT be written
// ********************************************************************************

// LvLevel is an alias of uint8 and is interchangeable.
type LvLevel = uint8

const (
	LvTrace LvLevel = iota + 1 // reserve 0 for none
	LvDebug
	LvInfo
	LvWarn
	LvError
	LvFatal
)

// NewLvWriter takes Writer and LvLevel (an alias for uint8) and returns LvWriter
func NewLvWriter(w Writer, lvl LvLevel) LvWriter {
	return LvWriter{}.SetOutput(w).SetLevel(lvl)
}

// LvWriter is a writer wrapper that supports levels such as
// LvTrace, LvDebug, LvInfo,, LvWarn, LvError, LvFatal or any uint8 (range 0-255).
type LvWriter struct {
	w       Writer
	lvMin   LvLevel
	lvCur   LvLevel // current level: this will be set for LvWriter.Lv()'s outputs
	enabled bool
}

// SetOutput will check if given Writer w is nil,
// If not nil, it will set the LvWriter with w, and enable it.
func (l LvWriter) SetOutput(w Writer) LvWriter {
	if w != nil {
		l.w = w
		l.enabled = true
		return l
	}
	l.w = nil
	l.enabled = false
	return l
}

// Output will return current output writer in LvWriter
func (l LvWriter) Output() Writer {
	return l.w
}

// Fd returns the integer Unix file descriptor if available
// This can be used to determine if a writer is capable for TTY. (like ANSI)
func (l LvWriter) Fd() uintptr {
	if l.w != nil {
		if v, ok := l.w.(interface{ Fd() uintptr }); ok {
			return v.Fd()
		}
	}
	return ^(uintptr(0))
}

// SetLevel will set log level of the LvWriter. Also LvWriter can have any uint8 values for logging.
// If fully customized log levels are being used, Lv(lvl) should be used instead of Info(), Warn()...
func (l LvWriter) SetLevel(lvl LvLevel) LvWriter {
	l.lvMin = lvl
	return l
}

// CopyLevel will set log level of the LvWriter. Also LvWriter can have any uint8 values for logging.
// IF lvCur is set from the source LvWriter (from), then, it will copy source's current value into minimum.
// This will allow: `log1.CopyLevel(log1.Info())` AND/OR `log1.CopyLevel(log2)`
// Otherwise, it will copy minimum value.
// If fully customized log levels are being used, Lv(lvl) should be used instead of Info(), Warn()...
func (l LvWriter) CopyLevel(from LvWriter) LvWriter {
	if from.lvCur != 0 {
		l.lvMin = from.lvCur
		return l
	}
	l.lvMin = from.lvMin
	return l
}

// GetLevel will return current lvMin level.
// Without explicitly set, this will be 0.
// (New, v0.7.12+)
func (l LvWriter) GetLevel() LvLevel {
	return l.lvMin
}

// Lv gets log level lvl, if it's above lvMin, it will return the LvWriter, and next func will print it.
// However, if given lvl is lower than lvMin, it will disable the write, and return it to next function,
// so it won't get printed. Without this setting, it will set to 0 (LvTrace) and prints all.
func (l LvWriter) Lv(lvl LvLevel) LvWriter {
	l.lvCur = lvl
	if l.lvMin <= lvl {
		return l
	}
	l.enabled = false
	return l
}

// Trace is a shortcut for Lv(LvTrace)
func (l LvWriter) Trace() LvWriter { return l.Lv(LvTrace) }

// Debug is a shortcut for Lv(LvDebug)
func (l LvWriter) Debug() LvWriter { return l.Lv(LvDebug) }

// Info is a shortcut for Lv(LvInfo)
func (l LvWriter) Info() LvWriter { return l.Lv(LvInfo) }

// Warn is a shortcut for Lv(LvWarn)
func (l LvWriter) Warn() LvWriter { return l.Lv(LvWarn) }

// Error is a shortcut for Lv(LvError)
func (l LvWriter) Error() LvWriter { return l.Lv(LvError) }

// Fatal is a shortcut for Lv(LvFatal)
func (l LvWriter) Fatal() LvWriter { return l.Lv(LvFatal) }

// Enabled will return true if LvWriter is enabled
func (l LvWriter) Enabled() bool { return l.enabled }

// Enable will set the writer's enabled value.
// Regardless of values given, if writer is not set,
// then it will be always disabled.
func (l LvWriter) Enable(enable bool) LvWriter {
	if l.w != nil {
		l.enabled = enable
		return l
	}
	l.enabled = false
	return l
}

// Close will close the writer w of LvWriter if compatible
func (l LvWriter) Close() error {
	if c, ok := l.w.(interface{ Close() error }); ok {
		return c.Close()
	}
	return nil
}

// Write will write byte slice p to the writer if available.
func (l LvWriter) Write(p []byte) (int, error) {
	if !l.enabled {
		return 0, nil
	}
	return l.w.Write(p)
}

// ********************************************************************************
// LvWriter Extended Methods: WriteString(), WriteAny()
// Extended methods have dependencies: sync, buf, bytes
// ********************************************************************************

// WriteString will take string and convert it to byte then writes.
// This may need some wrapper to solve the allocation issues
// DEPENDENCY: sync, buf, bytes
func (l LvWriter) WriteString(s string) (n int, err error) {
	if !l.enabled {
		return 0, nil
	}
	buf := GetBuffer()
	buf.Buf = buf.Buf.WriteString(s)
	if buf.Buf.LastByte() != '\n' {
		buf.Buf = buf.Buf.WriteBytes('\n')
	}
	n, err = l.w.Write(buf.Buf)
	PutBuffer(buf)
	return n, err
}

// WriteAny will take string and convert it to byte then writes.
// This may need some wrapper to solve the allocation issues.
// Note that currently WriteAny only supports few most used types.
// DEPENDENCY: sync, buf, bytes
func (l LvWriter) WriteAny(s ...interface{}) bool {
	if !l.enabled {
		return false
	}
	buf := GetBuffer()
	for i := 0; i < len(s); i++ {
		// If more type is needed, add it here.
		switch v := s[i].(type) {
		case string:
			buf.Buf = append(buf.Buf, v...)
		case int:
			buf.Buf = buf.Buf.WriteInt(v)
		case bool:
			buf.Buf = buf.Buf.WriteBool(v)
		case []byte:
			buf.Buf = append(buf.Buf, v...)
		case func([]byte) []byte: // This can be used to print current time as a function
			buf.Buf = v(buf.Buf)
		default:
			buf.Buf = append(buf.Buf, "UNSUPP"...)
		}
	}

	if buf.Buf.LastByte() != '\n' {
		buf.Buf = buf.Buf.WriteBytes('\n')
	}
	_, _ = l.w.Write(buf.Buf)
	PutBuffer(buf)
	return true
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
	// required for this Writer
	Write(p []byte) (n int, err error)
}

// StringWriter - to avoid importing "io"
type StringWriter interface {
	WriteString(s string) (n int, err error)
}

// Closer is an interface for the writers that have Close method.
type Closer interface {
	Close() error
}

// ********************************************************************************
// Discard Writer - satisfies Writer/StringWriter
// ********************************************************************************

// Discard - instead of using struct{}, just decided to use const (bool)
const Discard discardWriter = false

type discardWriter bool

// Write - to satisfy Writer interface
func (discardWriter) Write(p []byte) (int, error) { return len(p), nil }

// WriteString - for StringWriter interface
func (discardWriter) WriteString(s string) (int, error) { return len(s), nil }

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
