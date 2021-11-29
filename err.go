// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/29/2021

package gosl

// err is string type for error to save errors on constant.
// However, when this is being used, need to return a pointer of err
// to save allocations
type err string

// Error is to meet error interface
func (e err) Error() string {
	return string(e)
}

// NewError takes string and creates an error
// If the string received is empty, it will return error with nil value.
func NewError(s string) error {
	if s == "" {
		return nil
	}
	return err(s)
}

// IfErr is a simple function that takes an error ID and error..
// If error is not nil, then it will print error message.
// This has zero allocation.
func IfErr(key string, e error) {
	// If given error is not nil, then get a Buffer from the internal Buffer pool,
	// then write an error as "key = value" format and then write it to
	// os.Stdout using println().
	if e != nil {
		buf := GetBuffer()
		buf.WriteString(key).WriteString(" -> (err) ").WriteString(e.Error())
		println(buf.String())
		buf.Free()
	}
}

// IfPanic will take name and a function func(error),
// if function f is given, it will use that function,
// if not given, print the message using println (stdout)
// Usage:
//     func hello() (out string) {
//         defer IfPanic("hello", func(m string) { out = m })
//         panic("whatever")
//     }
func IfPanic(name string, f func(error)) {
	// This will only execute when recover() has something
	if r := recover(); r != nil {
		// NOTE: call to `recover()` cannot be inlined, so this function
		//       cannot be inlined even we split this.
		// Since `panic(interface{})` can be called with any value including error and string,
		// convert it to a string and will save to `m`
		var m error

		// based on the type of message from `recover()`, get string out of it.
		// if it was unexpected type, then set `m` with `<unknown>`.
		switch v := r.(type) {
		case error:
			m = v
		case string:
			m = err(v)
		case interface{ String() string }:
			m = err(v.String())
		default:
			m = err("unsupported panic info")
		}

		// if function function `f` is given, use it.
		// otherwise, write it to stdout using print.
		if f != nil {
			f(m)
		} else {
			// When no function is given, print it to screen
			buf := make(Buf, 0, 2<<10) // default Buffer to be 2k
			buf = buf.WriteString(name).
				WriteString(":Panic(`").WriteString(m.Error()).WriteString("`)")
			println(buf.String())
		}
	}
}

