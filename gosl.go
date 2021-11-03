// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/3/2021

package gosl

// DoNothing will be used to prevent nil error
func DoNothing() {}

// Filesize
const (
	KB int64 = 1024
	MB       = KB * 1024
	GB       = MB * 1024
	TB       = GB * 1024
	PB       = TB * 1024
	EB       = PB * 1024
)

// ********************************************************************************
// Internal Buffer Pool - bufp
// ********************************************************************************

const (
	internalBufferSize     = 1024
	internalBufferPoolSize = 256
)

// bufp is a byte buffer pool to be use internally for the logger, etc.
var bufp = NewPool(internalBufferPoolSize, func() interface{} {
	return &bufpBuffer{
		Buffer{
			Buf: make([]byte, 0, internalBufferSize),
		},
	}
})

type bufpBuffer struct {
	Buffer
}

func (b *bufpBuffer) ReturnBuffer() {
	if cap(b.Buffer.Buf) > (64 << 10) {
		return
	}
	b.Buf = b.Buf[:0]
	bufp.Put(b)
}

func getBufpBuffer() *bufpBuffer {
	return bufp.Get().(*bufpBuffer)
}


// OnPanic will take name and a function func(error),
// if function f is given, it will use that function,
// if not given, print the message using println (stdout)
// Usage:
//     func hello() (out string) {
//         defer OnPanic("hello", func(m string) { out = m })
//         panic("whatever")
//     }
func OnPanic(name string, f func(m string)) {
	// This will only execute when recover() has something
	if r := recover(); r != nil {
		// Since `panic(interface{})` can be called with any value including error and string,
		// convert it to a string and will save to `m`
		var m string

		// based on the type of message from `recover()`, get string out of it.
		// if it was unexpected type, then set `m` with `<unknown>`.
		switch v := r.(type) {
		case error:
			m = v.Error()
		case string:
			m = v
		default:
			m = "<unknown>"
		}

		// if function function `f` is given, use it.
		// otherwise, write it to stdout using print.
		if f != nil {
			f(m)
		} else {
			// When no function is given, print it to screen
			buf := make(Buf, 0, 1024)
			buf = buf.WriteString(name).
				WriteString(":Panic -> ").WriteString(m)
			println(buf.String())
		}
	}
}

