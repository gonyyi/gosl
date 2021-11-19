// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package gosl

// WARNING:
//     Code below will be slower than using standard libraries.
//     This is just a test to see if it can be built without
//     using any libraries (not even built-in), but only using
//     builtin functions of the language.

// NOTE:
//     gosl.Logger is a writer wrapper and to be used when developing
//     libraries. For a full featured logger, use
//     <https://github.com/gonyyi/alog> instead.

const (
	logStringQuote  = '"'
	logNewLine      = '\n'
	logKeyValueSign = ": "
	logKeyErrorSign = " -> "
)

// NewLogger will create a dlog.
func NewLogger(w Writer) Logger {
	return Logger{}.SetOutput(w)
}

// Logger uses value instead of pointer (8 bytes in 64bit system)
// as size of Logger is 24 bytes (21 bytes, but alloc is 24 bytes)
type Logger struct {
	w         Writer  // 16
	enable    bool    // 1
	prefix    [6]byte // 6
	prefixIdx uint8   // 1
	// prefix   func() []byte
}

// SetPrefix will a 6 byte prefix. If a string is longer, it will trim to 6.
func (l Logger) SetPrefix(prefix6 string) Logger {
	if prefix6 == "" {
		l.prefixIdx = 0
		return l
	}
	var tmp [6]byte
	for idx, c := range prefix6 {
		if idx > 5 {
			break
		}
		tmp[idx] = byte(c)
		l.prefixIdx = uint8(idx) + 1
	}
	l.prefix = tmp
	return l
}

// SetOutput will update the output writer of Logger
func (l Logger) SetOutput(w Writer) Logger {
	if w != nil {
		l.w = w
		l.enable = true
		return l
	}
	l.enable = false
	return l
}

// Enable only can be "true" when writer is not nil.
func (l Logger) Enable(enable bool) Logger {
	// if writer is nil, `enable` will be always false
	if l.w != nil {
		l.enable = enable
		return l
	}
	l.enable = false
	return l
}

// IfErr will take a key and err, and returns boolean OK.
// If there's an ERROR, this will return TRUE,
// if there's NO ERROR, this will return FALSE.
// this method log err ONLY WHEN it's not nil
func (l Logger) IfErr(key string, err error) (isError bool) {
	if l.enable && err != nil {
		l.ifErr(key, err)
		return true // has an error
	}
	return false // no error
}

// ifErr expects err won't be nil at this point.
func (l Logger) ifErr(key string, err error) {
	p := getBuffer()
	if l.prefixIdx != 0 {
		p.WriteByte(l.prefix[:l.prefixIdx]...)
	}
	p.WriteString(key).WriteString(logKeyErrorSign).WriteString("(err) ").WriteString(err.Error()).WriteByte(logNewLine)
	_, _ = l.w.Write(p.Buf)
	p.Free()
}

func (l Logger) write(p []byte) {
	_, _ = l.w.Write(p)
}

// String takes string and append newline
func (l Logger) String(s string) {
	if l.enable {
		l.string(s)
	}
}
func (l Logger) string(s string) {
	p := getBuffer()
	if l.prefixIdx != 0 {
		p.WriteByte(l.prefix[:l.prefixIdx]...)
	}
	p.WriteString(s).WriteByte(logNewLine)
	l.write(p.Buf)
	p.Free()
}

// KeyBool is a key-value log for string, bool
func (l Logger) KeyBool(key string, val bool) {
	if l.enable {
		l.keyBool(key, val)
	}
}
func (l Logger) keyBool(key string, val bool) {
	p := getBuffer()
	if l.prefixIdx != 0 {
		p.WriteByte(l.prefix[:l.prefixIdx]...)
	}
	p.WriteString(key)
	p.WriteString(logKeyValueSign)
	if val {
		p.WriteString("true")
	} else {
		p.WriteString("false")
	}
	p.WriteByte(logNewLine)
	l.write(p.Buf)
	p.Free()
	return
}

// KeyInt is a key-value log for string, int
func (l Logger) KeyInt(key string, val int) {
	if l.enable {
		l.keyInt(key, val)
	}
}
func (l Logger) keyInt(key string, val int) {
	p := getBuffer()
	if l.prefixIdx != 0 {
		p.WriteByte(l.prefix[:l.prefixIdx]...)
	}
	p.WriteString(key).WriteString(logKeyValueSign)
	p.Buf = AppendInt(p.Bytes(), val, false)
	p.WriteByte(logNewLine)
	l.write(p.Buf)
	p.Free()
}

// KeyFloat64 is a key-value log for string, float64
func (l Logger) KeyFloat64(key string, val float64) {
	if l.enable {
		l.keyFloat64(key, val)
	}
}
func (l Logger) keyFloat64(key string, val float64) {
	p := getBuffer()
	if l.prefixIdx != 0 {
		p.WriteByte(l.prefix[:l.prefixIdx]...)
	}
	p.WriteString(key).WriteString(logKeyValueSign)
	p.Buf = AppendFloat64(p.Bytes(), val, 2, false)
	p.WriteByte(logNewLine)
	l.write(p.Buf)
	p.Free()
}

// KeyString is a key-value log for string, string
func (l Logger) KeyString(key string, val string) {
	if l.enable {
		l.keyString(key, val)
	}
}
func (l Logger) keyString(key string, val string) {
	p := getBuffer()
	if l.prefixIdx != 0 {
		p.WriteByte(l.prefix[:l.prefixIdx]...)
	}
	p.WriteString(key).WriteString(logKeyValueSign).WriteByte(logStringQuote).WriteString(val).WriteByte(logStringQuote, logNewLine)
	l.write(p.Buf)
	p.Free()
}

// KeyError is a key-value log for an error,
func (l Logger) KeyError(key string, val error) {
	if l.enable {
		l.keyError(key, val)
	}
}
func (l Logger) keyError(key string, err error) {
	p := getBuffer()
	if l.prefixIdx != 0 {
		p.WriteByte(l.prefix[:l.prefixIdx]...)
	}
	p.WriteString(key).WriteString(logKeyErrorSign)
	if err != nil {
		p.WriteString("(err) ")
		p.WriteString(err.Error())
		p.WriteByte(logNewLine)
	} else {
		p.WriteString("<nil>")
		p.WriteByte(logNewLine)
	}

	l.write(p.Buf)
	p.Free()
	return
}

// Write takes bytes and returns number of bytes written and error
func (l Logger) Write(p []byte) (n int, err error) {
	if l.enable {
		if l.prefixIdx != 0 {
			return l.w.Write(append(l.prefix[:l.prefixIdx], p...))
		}
		return l.w.Write(p)
	}
	return 0, nil
}

// Enabled returns if current logger is enabled
func (l Logger) Enabled() bool {
	return l.enable
}

// Close will close writer if applicable
func (l Logger) Close() error {
	l.enable = false
	return Close(l.w)
}

