// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/20/2022

package gosl

// ********************************************************************************
// LvWriter is a writer with level supports. Level can be any uint8 type, but this
// has LvTrace ... LvFatal are predefined for convenience.
// Eg. Using Predefined Levels:
//     w := NewLvWriter( os.Stdout, LvInfo ) // Set minimum level as LvInfo
//     w.Debug().WriteString("blah bug blah.......") // will NOT be written
//     w.Info().WriteString("just fyi blah blah...") // will be written
// Eg. Using Custom Levels:
//     var (
//         INFO uint8 = 0    // Make sure levels are uint8 (0-255 level)
//         WARN uint8 = 3
//         ERRR uint8 = 7
//         // More...
//     )
//     w := NewLvWriter( os.Stdout, WARN ) // Writer will print for WARN or above
//     w.Lv(INFO).WriteString("blah info fyi...")  // will NOT be written
//     w.Lv(WARN).WriteString("HMM.............")  // will be written
//     w.Lv(ERRR).WriteString("blah ERROR!!!!!!")  // will be written
//     w = w.SetLevel(ERRR)  // change the level to ERRR
//     w.Lv(WARN).WriteString("HUH.............")  // now, this will NOT be written
// ********************************************************************************

const (
	LvTrace uint8 = iota
	LvDebug
	LvInfo
	LvWarn
	LvError
	LvFatal
)

// NewLvWriter takes Writer and level in uint8 and returns LvWriter
func NewLvWriter(w Writer, lvl uint8) LvWriter {
	return LvWriter{}.SetOutput(w).SetLevel(lvl)
}

// LvWriter is a writer wrapper that supports levels such as Trace, Debug, Info, Warn...
type LvWriter struct {
	w       Writer
	enabled bool  // w+bool = 24
	logLvl  uint8 // w + bool+ uint8 = 24
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

// SetLevel will set log level of the LvWriter. Also LvWriter can have any uint8 values for logging.
// If fully customized log levels are being used, Lv(lvl) should be used instead of Info(), Warn()...
func (l LvWriter) SetLevel(Lv uint8) LvWriter {
	l.logLvl = Lv
	return l
}

// Lv gets log level lvl, if it's above logLvl, it will return the LvWriter, and next func will print it.
// However, if given lvl is lower than logLvl, it will disable the write, and return it to next function,
// so it won't get printed. Without this setting, it will set to 0 (LvTrace) and prints all.
func (l LvWriter) Lv(lvl uint8) LvWriter {
	if l.logLvl <= lvl {
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

// WriteString will take string and convert it to byte then writes.
// This may need some wrapper to solve the allocation issues
func (l LvWriter) WriteString(s string) (int, error) {
	if !l.enabled {
		return 0, nil
	}
	return l.w.Write([]byte(s))
}

// Write will write byte slice p to the writer if available.
func (l LvWriter) Write(p []byte) (int, error) {
	if !l.enabled {
		return 0, nil
	}
	return l.w.Write(p)
}

// Close will close the writer w of LvWriter if compatible
func (l LvWriter) Close() error {
	if c, ok := l.w.(interface{ Close() error }); ok {
		return c.Close()
	}
	return nil
}

// LineWriter is a wrapper around the Writer with some unique feathers:
// 1. LineWriter does not cause panic when not initialized
// 2. LineWriter is concurrent safe (using mutex)
// 3. LineWriter appends newline at the end if one not present.
//
// LineWriter can create duplicates, but duplicated one's output is set to differently,
// then duplicate should be done before LineWriter is initialized.
// eg. This will NOT work correctly -- lw, lw1, lw2 will have the same writer w2.
//     lw := gosl.LineWriter{}.Init()
//     lw1 := lw.SetOutput(w1)
//     lw2 := lw.SetOutput(w2)
// eg. This will work correctly -- lw1 and lw2 will have independent writer.
//     lw := gosl.LineWriter{}
//     lw1 := lw.SetOutput(w1)
//     lw2 := lw.SetOutput(w2)

// LineWriterBufferSize is a default buffer size when started
var LineWriterBufferSize = 1024

// NewLineWriter will create a new LineWriter and initialize it.
func NewLineWriter(output Writer) LineWriter {
	return LineWriter{lw: &lineWriter{outp: output}}.Init().Enable(true)
}

// LineWriter is a writer wrapper that will add newline at the end if not present.
type LineWriter struct {
	lw      *lineWriter // to speed up
	enabled bool
}

type lineWriter struct {
	outp Writer
	buf  []byte
	mu   chan struct{}
}

// Init initialize LineWriter's buffer and mutex
func (w LineWriter) Init() LineWriter {
	if w.lw == nil {
		w.lw = &lineWriter{
			buf: make([]byte, LineWriterBufferSize),
			mu:  make(chan struct{}, 1),
		}
	} else if w.lw.mu == nil {
		w.lw.buf = make([]byte, LineWriterBufferSize)
		w.lw.mu = make(chan struct{}, 1)
	}
	return w
}

// SetOutput will initialize the LineWriter if not and set the output
func (w LineWriter) SetOutput(iow Writer) LineWriter {
	// initialize if not
	w = w.Init()
	w.lw.outp = iow
	return w.Enable(true)
}

// Output returns current output
func (w LineWriter) Output() Writer {
	w.Init()
	return w.lw.outp
}

// Enable will check the LineWriter's output and enable/disable if can
func (w LineWriter) Enable(t bool) LineWriter {
	w.Init()
	if w.lw.outp != nil {
		w.enabled = t
		return w
	}
	w.enabled = false
	return w
}

// Enabled returned current status
func (w LineWriter) Enabled() bool {
	return w.enabled
}

// Close will close the writer if applicable.
func (w LineWriter) Close() error {
	if w.lw != nil && w.lw.outp != nil {
		if c, ok := w.lw.outp.(Closer); ok {
			return c.Close()
		}
	}
	return nil
}

// WriteString will write string to the LineWriter
func (w LineWriter) WriteString(s string) (n int, err error) {
	if w.enabled {
		return w.writeString(s)
	}
	return 0, nil
}

// Write will write bytes to the LineWriter
func (w LineWriter) Write(p []byte) (n int, err error) {
	if w.enabled {
		return w.write(p)
	}
	return 0, nil
}

func (w LineWriter) writeString(s string) (n int, err error) {
	// At this point, it has to be enabled. And therefore, lw isn't nil.
	w.lw.mu <- struct{}{} // LOCK
	w.lw.buf = append(w.lw.buf[:0], s...)
	var adjN = 0
	if ls := len(s); ls > 0 && s[ls-1] != '\n' {
		w.lw.buf = append(w.lw.buf, '\n')
		adjN = -1 // When writes, return the count for one without added newline considered.
	}
	n, err = w.lw.outp.Write(w.lw.buf)
	<-w.lw.mu // UNLOCK

	return n + adjN, err
}

func (w LineWriter) write(p []byte) (n int, err error) {
	// At this point, it has to be enabled. And therefore, lw isn't nil.
	w.lw.mu <- struct{}{} // LOCK: to make sure all writes are sequential..
	if lp := len(p); lp > 0 && p[lp-1] != '\n' {
		w.lw.buf = append(w.lw.buf[:0], p...)
		w.lw.buf = append(w.lw.buf, '\n')
		n, err = w.lw.outp.Write(w.lw.buf)
		n -= 1 // When writes, return the count for one without added newline considered.
	} else {
		n, err = w.lw.outp.Write(p)
	}
	<-w.lw.mu // UNLOCK
	return n, err
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

// StringWriter - to avoid importing "io"
type StringWriter interface {
	WriteString(s string) (n int, err error)
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
