// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// FOR FULL FUNCTION LOGGER:
//       gosl.Logger is a writer wrapper and to be used when developing libraries.
//       For a full function logger, use https://github.com/gonyyi/alog

const (
	logStringQuote = '"'
	logNewLine     = '\n'
	logKeyValueSign = ": "
)

// Logger uses value instead of pointer (8 bytes in 64bit system)
// as size of Logger is 24 bytes (21 bytes, but alloc is 24 bytes)
type Logger struct {
	w      Writer // 16
	enable bool   // 1
}

// IfErr will take err and return err,
// this method log err ONLY WHEN it's not nil
func (l Logger) IfErr(key string, err error) {
	if l.enable && err != nil {
		l.ifErr(key, err)
	}
}

// ifErr expects err won't be nil at this point.
func (l Logger) ifErr(key string, err error) {
	p := getBufpBuffer()
	p.WriteString(key).WriteString(logKeyValueSign).WriteString("(err)").WriteString(err.Error()).WriteByte(logNewLine)
	p.ReturnBuffer()
}

// String takes string and append newline
func (l Logger) String(s string) {
	if l.enable {
		l.string(s)
	}
}
func (l Logger) string(s string) {
	p := getBufpBuffer()
	p.WriteString(s).WriteByte(logNewLine)
	_, _ = l.w.Write(p.Buf)
	p.ReturnBuffer()
}

// KeyBool is a key-value log for string, bool
// 10/20/2021, for the performance improvement, avoid using another function call
func (l Logger) KeyBool(key string, val bool) {
	if l.enable {
		l.keyBool(key, val)
	}
}
func (l Logger) keyBool(key string, val bool) {
	p := getBufpBuffer()
	p.WriteString(key)
	p.WriteString(logKeyValueSign)
	if val {
		p.WriteString("true")
	} else {
		p.WriteString("false")
	}
	p.WriteByte(logNewLine)
	_, _ = l.w.Write(p.Buf)
	p.ReturnBuffer()
	return
}

// KeyInt is a key-value log for string, int
func (l Logger) KeyInt(key string, val int) {
	if l.enable {
		l.keyInt(key, val)
	}
}
func (l Logger) keyInt(key string, val int) {
	p := getBufpBuffer()
	p.WriteString(key).WriteString(logKeyValueSign)
	p.Buf = AppendInt(p.Bytes(), val, false)
	p.WriteByte(logNewLine)
	_, _ = l.w.Write(p.Buf)
	p.ReturnBuffer()
}

// KeyFloat64 is a key-value log for string, float64
func (l Logger) KeyFloat64(key string, val float64) {
	if l.enable {
		l.keyFloat64(key, val)
	}
}
func (l Logger) keyFloat64(key string, val float64) {
	p := getBufpBuffer()
	p.WriteString(key).WriteString(logKeyValueSign)
	p.Buf = AppendFloat64(p.Bytes(), val, 2, false)
	p.WriteByte(logNewLine)
	_, _ = l.w.Write(p.Buf)
	p.ReturnBuffer()
}

// KeyString is a key-value log for string, string
func (l Logger) KeyString(key string, val string) {
	if l.enable {
		l.keyString(key, val)
	}
}
func (l Logger) keyString(key string, val string) {
	p := getBufpBuffer()
	p.WriteString(key).WriteString(logKeyValueSign).WriteByte(logStringQuote).WriteString(val).WriteByte(logStringQuote, logNewLine)
	_, _ = l.w.Write(p.Buf)
	p.ReturnBuffer()
}

// KeyError is a key-value log for an error,
func (l Logger) KeyError(key string, val error) {
	if l.enable {
		l.keyError(key, val)
	}
}
func (l Logger) keyError(key string, err error) {
	p := getBufpBuffer()
	p.WriteString(key).WriteString(logKeyValueSign)
	if err != nil {
		p.WriteString("(err)")
		p.WriteString(err.Error())
		p.WriteByte(logNewLine)
	} else {
		p.WriteString("<nil>")
		p.WriteByte(logNewLine)
	}

	_, _ = l.w.Write(p.Buf)
	p.ReturnBuffer()
	return
}

// Write takes bytes and returns number of bytes written and error
func (l Logger) Write(p []byte) (n int, err error) {
	if l.enable {
		return l.w.Write(p)
	}
	return 0, nil
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

// Enabled returns if current logger is enabled
func (l Logger) Enabled() bool {
	return l.enable
}

// Close will close writer if applicable
func (l Logger) Close() error {
	l.enable = false
	return Close(l.w)
}

// NewLogger will create a dlog.
func NewLogger(w Writer) Logger {
	return Logger{}.SetOutput(w)
}

