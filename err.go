// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/13/2021

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

// errWrap will be used to wrap an error
type errWrap struct {
	err  string
	prev error
}

// Unwrap for the error interface with wrap-able error
func (e *errWrap) Unwrap() error {
	if e.prev != nil {
		return e.prev
	}
	return nil
}

// Error to meet the error interface
func (e *errWrap) Error() string {
	if e.prev != nil {
		buf := GetBuffer()
		defer buf.Free()
		buf = buf.WriteString(e.err)
		// buf = buf.WriteBytes(':')
		// buf = buf.WriteString(e.prev.Error())
		// println("cur:", e.err, "prv:", e.prev.Error())
		return buf.String()
	}
	return e.err
}

// ErrorIs will check if error is same, or contains the given error
func ErrorIs(err, lookup error) bool {
	if err == lookup {
		return true
	}
	if ew, ok := err.(interface {
		Error() string
		Unwrap() error
	}); ok {
		if euw := ew.Unwrap(); euw != nil {
			if euw == lookup {
				return true
			}
			return ErrorIs(ew.Unwrap(), lookup)
		}
	}
	return err == lookup
}

// UnwrapError will unwrap error if available
func UnwrapError(e error) error {
	if ew, ok := e.(interface {
		Error() string
		Unwrap() error
	}); ok {
		return ew.Unwrap()
	}
	return nil
}

// WrapError will create a new error that has info
func WrapError(info string, e error) error {
	if e != nil {
		return &errWrap{
			err:  info + ":" + e.Error(),
			prev: e,
		}
	}
	return &errWrap{
		err: info,
	}
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
//         defer IfPanic("hello", func(m interface{}) { out = m.(string) })
//         panic("whatever")
//     }
func IfPanic(name string, f func(interface{})) {
	// This will only execute when recover() has something
	if r := recover(); r != nil {
		// NOTE: call to `recover()` cannot be inlined, so this function
		//       cannot be inlined even we split this.
		// if function function `f` is given, use it.
		// otherwise, write it to stdout using print.
		if f != nil {
			f(r)
		} else {
			// When no function is given, print it to screen
			buf := make(Buf, 0, 2<<10) // default Buffer to be 2k
			buf = buf.WriteString(name).WriteString(":Panic()")
			println(buf.String())
		}
	}
}
